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
	keyutil "github.com/node-a-team/terra-validator_exporter/function/keyutil"
	staking "github.com/node-a-team/terra-validator_exporter/function/staking"
	utils "github.com/node-a-team/terra-validator_exporter/function/utils"
	validators "github.com/node-a-team/terra-validator_exporter/function/validators"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	gaugesNum int = 2
	denomList = []string{"uluna", "ukrw", "usdr", "uusd"}
	gaugesNamespaceList = [...]string{"blockHeight", "currentBlockTime", "precommitRate", "validatorCount", "notBondedTokens", "bondedTokens", "totalBondedTokens", "bondingRate", "validatorCommitStatus", "proposerPriorityValue", "proposerPriority", "proposingStatus", "votingPower", "delegatorShares", "delegationRatio", "delegatorNumber", "commissionRate", "commissionMaxRate", "commissionMaxChangeRate", "minSelfDelegation"}
)


func newGauge(nameSpace string, name string, help string) prometheus.Gauge {
	result := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "" +nameSpace,
			Name:      "" +name,
			Help:      "" +help,
		})

	return result
}


func Exporter() {

	var gauges []prometheus.Gauge = make([]prometheus.Gauge, len(gaugesNamespaceList))
	var gaugesDenom []prometheus.Gauge = make([]prometheus.Gauge, len(denomList)*3)

	for i:=0; i<len(gaugesNamespaceList); i++ {
		gauges[i] = newGauge("Terra", gaugesNamespaceList[i], "")
		prometheus.MustRegister(gauges[i])
	}

	count := 0
        for i:=0; i<len(denomList)*3; i +=3 {
                gaugesDenom[i] = newGauge("Terra_rewards", denomList[count], "")
                gaugesDenom[i+1] = newGauge("Terra_commission", denomList[count], "")
                gaugesDenom[i+2] = newGauge("Terra_balances", denomList[count], "")
                prometheus.MustRegister(gaugesDenom[i])
                prometheus.MustRegister(gaugesDenom[i+1])
                prometheus.MustRegister(gaugesDenom[i+2])

		count++
        }




	for {

		blockStatus := block.BlockStatus()
		currentBlockHeight, _ := strconv.Atoi(blockStatus.Result.Block.Header.Height)

		commitStatus := commit.CommitStatus(currentBlockHeight)

		// validators
		validatorCountOrigin, validatorsetsStatus := validators.ValidatorsetsStatus()
		validatorsStatus := validators.ValidatorsStatus()

		// prints
		//              prints.PrintNew(blockStatus, commitStatus, consensusStatus, validatorsetsStatus, validatorsStatus, validatorNumber)








		// ##########################################################################################33

		fmt.Printf("\n\n\n> Chain_id: %s\n", blockStatus.Result.Block.Header.Chain_id)
		fmt.Printf("  - Height: %s\n", blockStatus.Result.Block.Header.Height)

/*		currentBlockTime := block.CalcBlockTime(blockStatus)
		fmt.Printf("  - Time: %s UTC\n", blockStatus.Result.Block.Header.Time.Format("060102 15:04:05"))
		fmt.Printf("  - Block_time: %0.2fs\n", currentBlockTime)

		fmt.Printf("  - Proposer: %s(%s)\n", block.ProposerMoniker(blockStatus.Result.Block.Header.Proposer_address, validatorsetsStatus, validatorsStatus), blockStatus.Result.Block.Header.Proposer_address)
		fmt.Printf("  - Precommit_rate: %f\n", commit.PrecommitRate(commitStatus)*100)
		fmt.Printf("  - Precommit_Last_commit.Len: %0.2f\n", float64(len(commitStatus.Result.Signed_header.Commit.Precommits)))
		fmt.Printf("  - Validaotorsets.Len: %0.2f\n", float64(validatorCount))

		latestRound := len(consensusStatus.Result.Round_state.Height_vote_set) - 1

		fmt.Println("  - height/round/step:", consensusStatus.Result.Round_state.Status)
		fmt.Println("  - prevotes:", consensusStatus.Result.Round_state.Height_vote_set[latestRound].Prevotes_bit_array)
		fmt.Println("  - precommit: ", consensusStatus.Result.Round_state.Height_vote_set[latestRound].Precommits_bit_array)

		notBondedTokens, bondedTokens := staking.GetStakingPool()
		totalBondedTokens := notBondedTokens + bondedTokens
		fmt.Println("\n  - notBondedTokens: ", notBondedTokens/math.Pow10(6))
		fmt.Println("  - bondedTokens: ", bondedTokens/math.Pow10(6))
		fmt.Println("  - totalTokens: ", totalBondedTokens/math.Pow10(6))
		fmt.Println("  - bonding Rate: ", bondedTokens/totalBondedTokens)
*/

		// block
		blockHeight := utils.StringToFloat64(blockStatus.Result.Block.Header.Height)
		currentBlockTime := block.CalcBlockTime(blockStatus)


		// commit
		precommitRate := commit.PrecommitRate(commitStatus)*100
		validatorCount := float64(validatorCountOrigin)


		// staking 
		notBondedTokensOrigin, bondedTokensOrigin := staking.GetStakingPool()
		totalBondedTokensOrigin := notBondedTokensOrigin + bondedTokensOrigin

		notBondedTokens := notBondedTokensOrigin/math.Pow10(6)
		bondedTokens := bondedTokensOrigin/math.Pow10(6)
		totalBondedTokens := totalBondedTokensOrigin/math.Pow10(6)
		bondingRate := bondedTokensOrigin/totalBondedTokensOrigin



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
		for _, validator_pubKey := range keys {

			// validatorsetsStatus
			cons_Bech32Address := validatorsetsStatus[validator_pubKey][0]

			votingPower := utils.StringToFloat64(validatorsetsStatus[validator_pubKey][1])
			proposerPriorityValue := utils.StringToFloat64(validatorsetsStatus[validator_pubKey][2])
			proposerPriority := utils.StringToFloat64(validatorsetsStatus[validator_pubKey][3])

			// validatorsStatus
			moniker := validatorsStatus[validator_pubKey][0]
			operator_address := validatorsStatus[validator_pubKey][1]

			jailed := validatorsStatus[validator_pubKey][2]
//			tokens := utils.StringToFloat64(validatorsStatus[validator_pubKey][3])/math.Pow10(6)
			delegatorShares := utils.StringToFloat64(validatorsStatus[validator_pubKey][4])/math.Pow10(6)
			commissionRate := utils.StringToFloat64(validatorsStatus[validator_pubKey][5])
			commissionMaxRate := utils.StringToFloat64(validatorsStatus[validator_pubKey][6])
			commissionMaxChangeRate := utils.StringToFloat64(validatorsStatus[validator_pubKey][7])
			//              commission_updateTime := validatorsStatus[validator_pubKey][8]
			minSelfDelegation := utils.StringToFloat64(validatorsStatus[validator_pubKey][9])
			//              unbonding_height := validatorsStatus[validator_pubKey][10]
			//              unbonding_time := validatorsStatus[validator_pubKey][11]
			//              identity := validatorsStatus[validator_pubKey][12]
			//              websote := validatorsStatus[validator_pubKey][13]
			//              details := validatorsStatus[validator_pubKey][14]


			delegationRatio := delegatorShares/bondedTokens


			// Additional ValidatorAddress
			cons_HexAddress := keyutil.RunFromBech32(cons_Bech32Address)
			account_address := keyutil.OperAddrToOtherAddr(operator_address)[0]

			// etc
			if operator_address == t.OperatorAddr {

				proposingStatus := float64(utils.GetPoposingCheck(blockStatus.Result.Block.Header.Proposer_address, cons_HexAddress))
				delegatorNumber := float64(validators.ValidatorDelegatorNumber(operator_address))

				fmt.Println("\n> Moniker: ", moniker)
				fmt.Println("  - validatorCount: ", validatorCount)
				fmt.Println("  - validator_pubKey: ", validator_pubKey)
				fmt.Println("  - operator_address: ", operator_address)
				fmt.Println("  - account_address: ", account_address)
				fmt.Println("  - cons_Bech32Address: ", cons_Bech32Address)
				fmt.Println("  - cons_HexAddress: ", cons_HexAddress)

				validatorCommitStatus := float64(commit.ValidatorPrecommitStatus(commitStatus, cons_HexAddress))
				fmt.Println("\n  - validatorCommitStatus: ", validatorCommitStatus)
				fmt.Println("  - proposer_priorityValue: ", proposerPriorityValue)
				fmt.Println("  - proposer_priority: ", proposerPriority)
				fmt.Println("  - proposingStatus: ", proposingStatus)
				fmt.Printf("  - voting_power: %f\n", votingPower)
				fmt.Println("  - jailed: ", jailed)
//				fmt.Println("  - tokens: ", tokens)
				fmt.Println("  - delegatorShares: ", delegatorShares)
				fmt.Printf("  - delegationRatio: %0.4f\n", delegationRatio)
				fmt.Println("  - delegatorNumber: ", delegatorNumber)
				fmt.Printf("  - commission_rate: %0.4f\n", commissionRate)
				fmt.Printf("  - commission_maxRate: %0.4f\n", commissionMaxRate)
				fmt.Printf("  - commission_maxChangeRate: %0.4f\n", commissionMaxChangeRate)
				fmt.Println("  - minSelfDelegation: ", minSelfDelegation)

				rewards, commission := validators.ValidatorRewards(operator_address)
				balances := validators.ValidatorBalances(account_address)

				count := 0
				for i:=0; i<len(denomList)*3; i +=3 {
					gaugesDenom[i].Set(utils.GetAmount(rewards, denomList[count]))
					gaugesDenom[i+1].Set(utils.GetAmount(commission, denomList[count]))
					gaugesDenom[i+2].Set(utils.GetAmount(balances, denomList[count]))


					fmt.Println("\n  - rewards_" +denomList[count] +": ", utils.GetAmount(rewards, denomList[count]))
					fmt.Println("  - commission_" +denomList[count] +": ", utils.GetAmount(commission, denomList[count]))
					fmt.Println("  - balances_" +denomList[count] +": ", utils.GetAmount(balances, denomList[count]))

					count++

				}
				// ##########################################################################################33

				gaugesValue := [...]float64{blockHeight, currentBlockTime, precommitRate, validatorCount, notBondedTokens, bondedTokens, totalBondedTokens, bondingRate, validatorCommitStatus, proposerPriorityValue, proposerPriority, proposingStatus, votingPower, delegatorShares, delegationRatio, delegatorNumber, commissionRate, commissionMaxRate, commissionMaxChangeRate, minSelfDelegation}

			        for i:=0; i<len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
			        }
			}

		}
		time.Sleep(2 * time.Second)
	}
}
