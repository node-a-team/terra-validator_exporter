package rest

import (
	"encoding/json"
	utils "github.com/node-a-team/terra-validator_exporter/utils"
)

type oracle struct {
	Miss	float64
}

type oracleMiss struct {
	Height string	`json:"height"`
	Result string	`json:"result"`
}

func getOracleMiss() oracle {

	var o oracle
        var om oracleMiss

        res := runRESTCommand("/oracle/voters/" +OperAddr +"/miss")
        json.Unmarshal(res, &om)

	o.Miss = utils.StringToFloat64(om.Result)

        return o
}
