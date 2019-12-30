package rest

import (
      "encoding/json"
      utils "github.com/node-a-team/terra-validator_exporter/utils"
)

type delegations struct {
	Height	string	`json:"height"`
	Result	[]delegation
}

type delegation struct {
	Delegator_address	string	`json:"delegator_address"`
	Validator_address	string	`json:"validator_address"`
	Shares			string	`json:"shares"`
	Balance			string	`json:"balance"`
}

type delegationInfo struct {
	DelegationCount	float64
	SelfDelegation	float64
}

func getDelegations(accAddr string) delegationInfo {

	var d delegations
	var dInfo delegationInfo

	res := runRESTCommand("/staking/validators/" +OperAddr +"/delegations")
	json.Unmarshal(res, &d)


	dInfo.DelegationCount = float64(len(d.Result))

	for _, value := range d.Result {
		if accAddr == value.Delegator_address {
			dInfo.SelfDelegation = utils.StringToFloat64(value.Shares)
		}
	}

	return dInfo
}
