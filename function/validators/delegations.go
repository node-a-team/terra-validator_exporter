package validators

import (
        t "github.com/node-a-team/terra-validator_exporter/types"

        "os/exec"
        "encoding/json"
)


var (
)

func ValidatorDelegatorNumber(operatorAddr string) int64 {

	var validatorDelegationStatus []t.ValidatorDelegationStatus
	var result int64

        cmd := "curl -s -XGET " +t.RestServer +"/staking/validators/"+operatorAddr +"/delegations"  +" -H \"accept:application/json\""
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
        json.Unmarshal(out, &validatorDelegationStatus)

	result = int64(len(validatorDelegationStatus))


	return result
}


