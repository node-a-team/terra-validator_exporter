package rest

import (
//        "fmt"
      "encoding/json"
)

type balances struct {
	Height	string	`json:"height"`
	Result	[]Coin
}

type Coin struct {
	Denom   string
        Amount  string
}

func getBalances(accAddr string) []Coin {

	var b balances

	res := runRESTCommand("/bank/balances/" +accAddr)
	json.Unmarshal(res, &b)

	return b.Result
}
