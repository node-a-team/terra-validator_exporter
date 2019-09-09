package utils

import (
	"math"
	t "terra-validator_exporter/types"
)

func GetAmount(rewards []t.Coin, denom string) float64 {

	var r float64

	for _, value := range rewards {
		if value.Denom == denom {
			r = StringToFloat64(value.Amount) / math.Pow10(6)
		}
	}

	return r

}

func GetPoposingCheck(proposerAddress string, validatorConsHexAddress string) int {
	var result int = 0

	if proposerAddress == validatorConsHexAddress {
		result = 1
	}

	return result
}
