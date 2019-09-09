package exporter

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	t "github.com/node-a-team/terra-validator_exporter/types"

	block "github.com/node-a-team/terra-validator_exporter/function/block"
	commit "github.com/node-a-team/terra-validator_exporter/function/commit"
	consensus "github.com/node-a-team/terra-validator_exporter/function/consensus"
	keyutil "github.com/node-a-team/terra-validator_exporter/function/keyutil"
	staking "github.com/node-a-team/terra-validator_exporter/function/staking"
	utils "github.com/node-a-team/terra-validator_exporter/function/utils"
	validators "github.com/node-a-team/terra-validator_exporter/function/validators"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	denomList           = []string{"uluna", "ukrw", "usdr", "uusd"}
	gaugesNamespaceList = [...]string{"blockHeight", "currentBlockTime", "precommitRate", "proposerWalletAccountNumber", "validatorCount", "notBondedTokens", "bondedTokens", "totalBondedTokens", "bondingRate", "validatorCommitStatus", "proposerPriorityValue", "proposerPriority", "proposingStatus", "votingPower", "delegatorShares", "delegationRatio", "delegatorCount", "selfDelegationAmount", "commissionRate", "commissionMaxRate", "commissionMaxChangeRate", "minSelfDelegation", "jailed"}

	contentsColorInit string = "\033[0m"
)

func newGauge(nameSpace string, name string, help string) prometheus.Gauge {
	result := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "" + nameSpace,
			Name:      "" + name,
			Help:      "" + help,
		})

	return result
}

func Exporter() {

	var gauges []prometheus.Gauge = make([]prometheus.Gauge, len(gaugesNamespaceList))
	var gaugesDenom []prometheus.Gauge = make([]prometheus.Gauge, len(denomList)*3)

	for i := 0; i < len(gaugesNamespaceList); i++ {
		gauges[i] = newGauge("Terra", gaugesNamespaceList[i], "")
		prometheus.MustRegister(gauges[i])
	}

	count := 0
	for i := 0; i < len(denomList)*3; i += 3 {
		gaugesDenom[i] = newGauge("Terra_rewards", denomList[count], "")
		gaugesDenom[i+1] = newGauge("Terra_commission", denomList[count], "")
		gaugesDenom[i+2] = newGauge("Terra_balances", denomList[count], "")
		prometheus.MustRegister(gaugesDenom[i])
		prometheus.MustRegister(gaugesDenom[i+1])
		prometheus.MustRegister(gaugesDenom[i+2])

		count++
	}

	gaugesForLabel := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "Terra",
			Name:      "labels",
			Help:      "",
		},
		[]string{"chainId", "moniker", "validatorPubKey", "operatorAddress", "accountAddress", "consHexAddress"},
	)
	prometheus.MustRegister(gaugesForLabel)

	// for csv file export
	// Run once at first, then run once every 10000block
	fileExportChecker := 0
	baseBlockForFileExport := 10000

	for {

		blockStatus := block.BlockStatus()
		currentBlockHeight, _ := strconv.Atoi(blockStatus.Result.Block.Header.Height)
		consensusStatus := consensus.ConsensusStatus()
		commitStatus := commit.CommitStatus(currentBlockHeight)

		// validators
		validatorCountOrigin, validatorsetsStatus := validators.ValidatorsetsStatus()
		validatorsStatus := validators.ValidatorsStatus()

		// block
		chainId := blockStatus.Result.Block.Header.Chain_id
		blockTime := blockStatus.Result.Block.Header.Time.Format("060102 15:04:05")
		blockHeight := utils.StringToFloat64(blockStatus.Result.Block.Header.Height)
		currentBlockTime := block.CalcBlockTime(blockStatus)

		// commit
		precommitRate := commit.PrecommitRate(commitStatus) * 100
		validatorCount := float64(validatorCountOrigin)
		proposerConsHexAddress := blockStatus.Result.Block.Header.Proposer_address
		proposerMoniker := block.ProposerMoniker(blockStatus.Result.Block.Header.Proposer_address, validatorsetsStatus, validatorsStatus)
		proposerWalletAccountNumber := float64(0.0)

		// staking
		notBondedTokensOrigin, bondedTokensOrigin := staking.GetStakingPool()
		totalBondedTokensOrigin := notBondedTokensOrigin + bondedTokensOrigin

		notBondedTokens := notBondedTokensOrigin / math.Pow10(6)
		bondedTokens := bondedTokensOrigin / math.Pow10(6)
		totalBondedTokens := totalBondedTokensOrigin / math.Pow10(6)
		bondingRate := bondedTokensOrigin / totalBondedTokensOrigin

		// consensus
		latestRound := len(consensusStatus.Result.Round_state.Height_vote_set) - 1
		heightRoundStep := consensusStatus.Result.Round_state.Status
		prevotes := consensusStatus.Result.Round_state.Height_vote_set[latestRound].Prevotes_bit_array
		precommit := consensusStatus.Result.Round_state.Height_vote_set[latestRound].Precommits_bit_array

		// csv file export(validatorsAccountNumber)
		if fileExportChecker == 0 || int(blockHeight)%baseBlockForFileExport == 0 {
			validators.ValidatorsAccountNumber(blockHeight, validatorsetsStatus, validatorsStatus)
			fileExportChecker++
		}

		// sorting
		keys := []string{}

		for key := range validatorsetsStatus {
			keys = append(keys, key)
		}

		sort.Slice(keys, func(i, j int) bool {
			a, _ := strconv.Atoi(validatorsetsStatus[keys[i]][3])
			b, _ := strconv.Atoi(validatorsetsStatus[keys[j]][3])
			return a < b
		})

		// validator_pubkey: gaiad tendermint show-validator -> priv_validator_key.json
		for _, validatorPubKey := range keys {

			// validatorsStatus#1
			operatorAddress := validatorsStatus[validatorPubKey][1]

			// our validator
			if operatorAddress == t.OperatorAddr {

				// get proposerWalletAccountNumber
				for _, proposerValidatorPubKey := range keys {

					operatorAddress := validatorsStatus[proposerValidatorPubKey][1]
					consBech32Address := validatorsetsStatus[proposerValidatorPubKey][0]
					consHexAddress := keyutil.RunFromBech32(consBech32Address)

					if proposerConsHexAddress == consHexAddress {
						accountAddress := keyutil.OperAddrToOtherAddr(operatorAddress)[0]
						_, proposerWalletAccountNumber = validators.ValidatorAccount(accountAddress)
					}
				}

				// validatorsetsStatus
				consBech32Address := validatorsetsStatus[validatorPubKey][0]
				votingPower := utils.StringToFloat64(validatorsetsStatus[validatorPubKey][1])
				proposerPriorityValue := utils.StringToFloat64(validatorsetsStatus[validatorPubKey][2])
				proposerPriority := utils.StringToFloat64(validatorsetsStatus[validatorPubKey][3])

				// validatorsStatus#2
				moniker := validatorsStatus[validatorPubKey][0]

				jailed := utils.BoolStringToFloat64(validatorsStatus[validatorPubKey][2])
				//				tokens := utils.StringToFloat64(validatorsStatus[validatorPubKey][3])/math.Pow10(6)
				delegatorShares := utils.StringToFloat64(validatorsStatus[validatorPubKey][4]) / math.Pow10(6)
				commissionRate := utils.StringToFloat64(validatorsStatus[validatorPubKey][5])
				commissionMaxRate := utils.StringToFloat64(validatorsStatus[validatorPubKey][6])
				commissionMaxChangeRate := utils.StringToFloat64(validatorsStatus[validatorPubKey][7])
				//				commission_updateTime := validatorsStatus[validatorPubKey][8]
				minSelfDelegation := utils.StringToFloat64(validatorsStatus[validatorPubKey][9])
				//				unbonding_height := validatorsStatus[validatorPubKey][10]
				//				unbonding_time := validatorsStatus[validatorPubKey][11]
				//				identity := validatorsStatus[validatorPubKey][12]
				//				websote := validatorsStatus[validatorPubKey][13]
				//				details := validatorsStatus[validatorPubKey][14]

				// keyutil
				accountAddress := keyutil.OperAddrToOtherAddr(operatorAddress)[0]
				consHexAddress := keyutil.RunFromBech32(consBech32Address)

				// etc
				proposingStatus := float64(utils.GetPoposingCheck(blockStatus.Result.Block.Header.Proposer_address, consHexAddress))
				delegatorCount, selfDelegationAmountOrigin := validators.ValidatorDelegatorNumber(operatorAddress, accountAddress)
				delegationRatio := delegatorShares / bondedTokens
				selfDelegationAmount := selfDelegationAmountOrigin / math.Pow10(6)
				validatorCommitStatus := float64(commit.ValidatorPrecommitStatus(commitStatus, consHexAddress))
				rewards, commission := validators.ValidatorRewards(operatorAddress)
				balances, walletAccountNumber := validators.ValidatorAccount(accountAddress)

				// print
				if t.OutputPrint {
					fmt.Printf("\033[1m\033[7m\033[32m[ ############ Chain_id: %s ############ ]\n\n"+contentsColorInit, chainId)
					fmt.Printf("\033[1m> Height: \033[32m%0.0f\n"+contentsColorInit, blockHeight)

					fmt.Printf("  - Time: %s UTC\n", blockTime)
					fmt.Printf("  - Block_time: %0.2fs\n", currentBlockTime)

					fmt.Printf("  - Proposer: %s(%s)\n", proposerMoniker, proposerConsHexAddress)
					fmt.Printf("  - Precommit_rate: %f\n", precommitRate)

					fmt.Println("  - height/round/step:", heightRoundStep)
					fmt.Println("  - prevotes:", prevotes)
					fmt.Println("  - precommit: ", precommit)

					fmt.Println("\n  - notBondedTokens: ", notBondedTokens)
					fmt.Println("  - bondedTokens: ", bondedTokens)
					fmt.Println("  - totalTokens: ", totalBondedTokens)
					fmt.Println("  - bonding Rate: ", bondedTokens/totalBondedTokens)

					fmt.Printf("\n\n\033[1m> Moniker: \033[33m%s\n"+contentsColorInit, moniker)
					fmt.Println("  - validatorCount: ", validatorCount)
					fmt.Println("  - validator_pubKey: ", validatorPubKey)
					fmt.Println("  - operator_address: ", operatorAddress)
					fmt.Println("  - account_address: ", accountAddress)
					fmt.Println("  - cons_Bech32Address: ", consBech32Address)
					fmt.Println("  - cons_HexAddress: ", consHexAddress)

					fmt.Println("\n  - validatorCommitStatus: ", validatorCommitStatus)
					fmt.Println("  - proposer_priorityValue: ", proposerPriorityValue)
					fmt.Println("  - proposer_priority: ", proposerPriority)
					fmt.Println("  - proposingStatus: ", proposingStatus)
					fmt.Printf("  - voting_power: %f\n", votingPower)
					fmt.Println("  - jailed: ", jailed)
					//				fmt.Println("  - tokens: ", tokens)
					fmt.Println("  - delegatorShares: ", delegatorShares)
					fmt.Printf("  - delegationRatio: %0.4f\n", delegationRatio)
					fmt.Println("  - delegatorCount: ", delegatorCount)
					fmt.Println("  - selfDelegationAmount: ", selfDelegationAmount)
					fmt.Printf("  - commission_rate: %0.4f\n", commissionRate)
					fmt.Printf("  - commission_maxRate: %0.4f\n", commissionMaxRate)
					fmt.Printf("  - commission_maxChangeRate: %0.4f\n", commissionMaxChangeRate)
					fmt.Println("  - minSelfDelegation: ", minSelfDelegation)
					fmt.Println("  - walletAccountNumber: ", walletAccountNumber)
				}
				count := 0
				for i := 0; i < len(denomList)*3; i += 3 {
					gaugesDenom[i].Set(utils.GetAmount(rewards, denomList[count]))
					gaugesDenom[i+1].Set(utils.GetAmount(commission, denomList[count]))
					gaugesDenom[i+2].Set(utils.GetAmount(balances, denomList[count]))

					if t.OutputPrint {
						fmt.Println("\n  - rewards_"+denomList[count]+": ", utils.GetAmount(rewards, denomList[count]))
						fmt.Println("  - commission_"+denomList[count]+": ", utils.GetAmount(commission, denomList[count]))
						fmt.Println("  - balances_"+denomList[count]+": ", utils.GetAmount(balances, denomList[count]))
					}

					count++
				}

				// prometheus giages value
				gaugesValue := [...]float64{blockHeight, currentBlockTime, precommitRate, proposerWalletAccountNumber, validatorCount, notBondedTokens, bondedTokens, totalBondedTokens, bondingRate, validatorCommitStatus, proposerPriorityValue, proposerPriority, proposingStatus, votingPower, delegatorShares, delegationRatio, delegatorCount, selfDelegationAmount, commissionRate, commissionMaxRate, commissionMaxChangeRate, minSelfDelegation, jailed}

				for i := 0; i < len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
				}

				gaugesForLabel.WithLabelValues(chainId, moniker, validatorPubKey, operatorAddress, accountAddress, consHexAddress).Add(0)
			}

		}
		time.Sleep(2 * time.Second)
	}
}
