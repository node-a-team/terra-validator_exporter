package rest

import (
      "encoding/json"
      utils "github.com/node-a-team/terra-validator_exporter/utils"
)

type stakingPool struct {
	Height	string	`json:"height"`
	Result	struct {
		Not_bonded_tokens	string	`json:"not_bonded_tokens"`
		Bonded_tokens		string	`json:"bonded_tokens"`
		Total_supply		float64
	}
}

type totalSupply struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

func getStakingPool() stakingPool {

	var sp stakingPool

	res := runRESTCommand("/staking/pool")
	json.Unmarshal(res, &sp)

	sp.Result.Total_supply = getTotalSupply("luna")

	return sp
}

func getTotalSupply(denom string) float64 {

        var ts totalSupply

        res := runRESTCommand("/supply/total/u" +denom)
        json.Unmarshal(res, &ts)

        return utils.StringToFloat64(ts.Result)
}
