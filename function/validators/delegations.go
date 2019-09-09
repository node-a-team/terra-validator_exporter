package validators

import (
	"encoding/json"
	"os/exec"
	utils "terra-validator_exporter/function/utils"
	t "terra-validator_exporter/types"
)

var ()

func ValidatorDelegatorNumber(operatorAddr string, accountAddr string) (delegatorCount float64, selfDelegationAmount float64) {

	var validatorDelegationStatus []t.ValidatorDelegationStatus

	delegatorCount, selfDelegationAmount = 0.0, 0.0

	cmd := "curl -s -XGET " + t.RestServer + "/staking/validators/" + operatorAddr + "/delegations" + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &validatorDelegationStatus)

	delegatorCount = float64(len(validatorDelegationStatus))

	for _, value := range validatorDelegationStatus {
		if value.Delegator_Address == accountAddr {
			selfDelegationAmount = utils.StringToFloat64(value.Shares)
		}
	}

	return delegatorCount, selfDelegationAmount
}
