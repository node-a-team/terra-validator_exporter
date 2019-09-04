package prints

import (
	t "github.com/node-a-team/terra-validator_exporter/types"

	block "github.com/node-a-team/terra-validator_exporter/function/block"
	commit "github.com/node-a-team/terra-validator_exporter/function/commit"
	keyutil "github.com/node-a-team/terra-validator_exporter/function/keyutil"
	validators "github.com/node-a-team/terra-validator_exporter/function/validators"
	staking "github.com/node-a-team/terra-validator_exporter/function/staking"
	utils "github.com/node-a-team/terra-validator_exporter/function/utils"

	"fmt"
	"sort"
	"strconv"
	"math"
)
var (
	emptySpace string = "  "

	contentsColor string
        contentsColorInit string = "\033[0m"
)

func PrintNew(blockStatus t.BlockStatus, commitStatus t.CommitStatus, consensusStatus t.ConsensusStatus, validatorsetsStatus map[string][]string, validatorsStatus map[string][]string, validatorCount int) {


	fmt.Printf("\n\n\n> Chain_id: %s\n", blockStatus.Result.Block.Header.Chain_id)
	fmt.Printf("  - Height: %s\n", blockStatus.Result.Block.Header.Height)

	currentBlockTime := block.CalcBlockTime(blockStatus)
	fmt.Printf("  - Time: %s UTC\n", blockStatus.Result.Block.Header.Time.Format("060102 15:04:05"))
	fmt.Printf("  - Block_time: %0.2fs\n", currentBlockTime)

	fmt.Printf("  - Proposer: %s(%s)\n", block.ProposerMoniker(blockStatus.Result.Block.Header.Proposer_address, validatorsetsStatus, validatorsStatus), blockStatus.Result.Block.Header.Proposer_address)
	fmt.Printf("  - Precommit_rate: %f\n", commit.PrecommitRate(commitStatus)*100)
	fmt.Printf("  - Precommit_Last_commit.Len: %0.2f\n", float64(len(commitStatus.Result.Signed_header.Commit.Precommits)))
	fmt.Printf("  - Validaotorsets.Len: %0.2f\n", float64(validatorNumber))


	latestRound := len(consensusStatus.Result.Round_state.Height_vote_set)-1

	fmt.Println("  - height/round/step:", consensusStatus.Result.Round_state.Status)
        fmt.Println("  - prevotes:", consensusStatus.Result.Round_state.Height_vote_set[latestRound].Prevotes_bit_array)
        fmt.Println("  - precommit: ", consensusStatus.Result.Round_state.Height_vote_set[latestRound].Precommits_bit_array)

	notBondedTokens, bondedTokens := staking.GetStakingPool()
	totalBondedTokens := notBondedTokens + bondedTokens
	fmt.Println("\n  - notBondedTokens: ", notBondedTokens/math.Pow10(6))
	fmt.Println("  - bondedTokens: ", bondedTokens/math.Pow10(6))
	fmt.Println("  - totalTokens: ", totalBondedTokens/math.Pow10(6))
	fmt.Println("  - bonding Rate: ", bondedTokens/totalBondedTokens)





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
                cons_HexAddress := keyutil.RunFromBech32(cons_Bech32Address)
                voting_power := validatorsetsStatus[validator_pubKey][1]
                proposer_priorityValue := validatorsetsStatus[validator_pubKey][2]
                proposer_priority := validatorsetsStatus[validator_pubKey][3]


                // validatorsStatus
		moniker := validatorsStatus[validator_pubKey][0]
                operator_address := validatorsStatus[validator_pubKey][1]

                account_address :=  keyutil.OperAddrToOtherAddr(operator_address)[0]
//              accountDenom, accountAmount := account.Account(account_address)


                jailed := validatorsStatus[validator_pubKey][2]
                tokens := validatorsStatus[validator_pubKey][3]
                delegatorShares := validatorsStatus[validator_pubKey][4]
                commission_rate := utils.StringToFloat64(validatorsStatus[validator_pubKey][5])
              commission_maxRate := utils.StringToFloat64(validatorsStatus[validator_pubKey][6])
              commission_maxChangeRate := utils.StringToFloat64(validatorsStatus[validator_pubKey][7])
//              commission_updateTime := validatorsStatus[validator_pubKey][8]
                minSelfDelegation := validatorsStatus[validator_pubKey][9]
//              unbonding_height := validatorsStatus[validator_pubKey][10]
//              unbonding_time := validatorsStatus[validator_pubKey][11]
//              identity := validatorsStatus[validator_pubKey][12]
//              websote := validatorsStatus[validator_pubKey][13]
//              details := validatorsStatus[validator_pubKey][14]

//              setContentsColor(proposer_priority)


		// etc
                if operator_address == t.OperatorAddr {

			proposingStatus := poposingCheck(blockStatus.Result.Block.Header.Proposer_address, cons_HexAddress)
			delegatorNumber := validators.ValidatorDelegatorNumber(operator_address)

                        fmt.Println("\n> Moniker: ", moniker)
			fmt.Println("  - validatorCount: ", validatorCount)
                        fmt.Println("  - validator_pubKey: ", validator_pubKey)
                        fmt.Println("  - operator_address: ", operator_address)
                        fmt.Println("  - account_address: ", account_address)
                        fmt.Println("  - cons_Bech32Address: ", cons_Bech32Address)
                        fmt.Println("  - cons_HexAddress: ", cons_HexAddress)

			validatorCommitStatus := commit.ValidatorPrecommitStatus(commitStatus, cons_HexAddress)
			fmt.Println("\n  - validatorCommitStatus: ", validatorCommitStatus)
			fmt.Println("  - proposer_priorityValue: ", proposer_priorityValue)
			fmt.Println("  - proposer_priority: ", proposer_priority)
			fmt.Println("  - proposingStatus: ", proposingStatus)
			fmt.Println("  - voting_power: ", voting_power)
			fmt.Println("  - jailed: ", jailed)
			fmt.Println("  - tokens: ", tokens)
			fmt.Println("  - delegatorShares: ", utils.StringToFloat64(delegatorShares)/math.Pow10(6))
			fmt.Printf("  - delegationRatio: %0.4f\n", (utils.StringToFloat64(delegatorShares)/math.Pow10(6))/(bondedTokens/math.Pow10(6)))
			fmt.Println("  - delegatorNumber: ", delegatorNumber)
			fmt.Printf("  - commission_rate: %0.4f\n", commission_rate)
			fmt.Printf("  - commission_maxRate: %0.4f\n", commission_maxRate)
			fmt.Printf("  - commission_maxChangeRate: %0.4f\n", commission_maxChangeRate)
			fmt.Println("  - minSelfDelegation: ", minSelfDelegation)

			rewards, commission := validators.ValidatorRewards(operator_address)
			fmt.Println("\n  - rewards: ", utils.GetAmount(rewards, "uatom")/math.Pow10(6))
			fmt.Println("  - commission: ", utils.GetAmount(commission, "uatom")/math.Pow10(6))

			balances := validators.ValidatorBalances(account_address)
			fmt.Println("  - balances: ", utils.GetAmount(balances, "uatom")/math.Pow10(6))
			fmt.Println("", )
		}
/*
                fmt.Printf(contentsColor)
                fmt.Printf(emptySpace +"%s\t\t", proposer_priority)
                fmt.Printf(emptySpace +" \033[32m%s\033[0m\t", split(moniker, 10))
                fmt.Printf(contentsColor)
                fmt.Printf(emptySpace +" %s\t", split(voting_power, 10))

                if len(accountAmount) != 0 {
                        fmt.Printf(emptySpace +" %14.6f %s\t", float64(accountAmount[0])/100000, accountDenom[0])

                } else {
                        fmt.Printf(emptySpace +"\t[ null ]\t")
                }

                fmt.Printf(emptySpace +" %0.2f\t\t", commission_rate)
                fmt.Printf(emptySpace +" %s\t", jailed)
//              fmt.Printf(emptySpace +"%s\t ", split(operator_address, 20))
//              fmt.Printf(emptySpace +"%s\n", split(validator_pubKey, 30))
                fmt.Printf(contentsColorInit +"\n")
*/

        }





}

func poposingCheck(proposerAddress string, validatorConsHexAddress string) int{
	var result int = 0

	if proposerAddress == validatorConsHexAddress {
		result = 1
	}

	return result
}

func split(str string, length int) string {

	var result string

	if len(str)+1 >= length {
		result = str[:length-1] +".."
	} else {
		result = str
		for i:=0; i < length - (len(str)+1); i++ {
			result = result +" "
		}
	}

	return result
}

