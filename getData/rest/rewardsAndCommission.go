package rest

import (
      "encoding/json"
)

type rewardsAndCommisson struct {

	Height	string		`json:"height"`
	Result	struct {
                Operator_Address        string  `"json:"operator_address"`
                Self_bond_rewards         []Coin  `"json:"self_bond_rewards"`
                Val_commission      []Coin  `"json:"val_commission"`
	}

}

func getRewardsAndCommisson() ([]Coin, []Coin) {

	var rc rewardsAndCommisson

	res := runRESTCommand("/distribution/validators/" +OperAddr)
	json.Unmarshal(res, &rc)

	return rc.Result.Self_bond_rewards, rc.Result.Val_commission
}
