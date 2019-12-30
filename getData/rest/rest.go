package rest

import (
//	"fmt"
	"os/exec"
	utils "github.com/node-a-team/terra-validator_exporter/utils"
)

var (
        Addr string
	OperAddr string
)


type RESTData struct {

	BlockHeight	int64
	StakingPool	stakingPool

	Validatorsets	map[string][]string
	Validators	validator
	Delegations	delegationInfo
	Balances	[]Coin
	Rewards		[]Coin
	Commission	[]Coin

	Oracle		oracle
	Gov		govInfo
}

func newRESTData(blockHeight int64) *RESTData {

	rd := &RESTData {
		BlockHeight:	blockHeight,
		Validatorsets:	make(map[string][]string),
        }

	return rd
}

func GetData(blockHeight int64) (*RESTData, string) {

	accAddr := utils.GetAccAddrFromOperAddr(OperAddr)

	rd := newRESTData(blockHeight)
	rd.StakingPool = getStakingPool()

	rd.Validatorsets = getValidatorsets(blockHeight)
	rd.Validators = getValidators()
	rd.Delegations = getDelegations(accAddr)
	rd.Balances = getBalances(accAddr)
	rd.Rewards, rd.Commission = getRewardsAndCommisson()

	rd.Oracle = getOracleMiss()
	rd.Gov = getGovInfo()

	consHexAddr := utils.Bech32AddrToHexAddr(rd.Validatorsets[rd.Validators.ConsPubKey][0])
	return rd, consHexAddr
}

func runRESTCommand(str string) []uint8 {
        cmd := "curl -s -XGET " +Addr +str +" -H \"accept:application/json\""
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
//	fmt.Println(cmd)

        return out
}
