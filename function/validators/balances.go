package validators

import (
        t "github.com/node-a-team/terra-validator_exporter/types"

	"os/exec"
        "encoding/json"
)


var (
)

func ValidatorBalances(accountAddr string) (balances []t.Coin) {

	var validatorBalances []t.Coin

        cmd := "curl -s -XGET " +t.RestServer +"/bank/balances/"+accountAddr  +" -H \"accept:application/json\""
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
        json.Unmarshal(out, &validatorBalances)

	balances = validatorBalances

	return balances
}

