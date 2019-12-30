package utils

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/libs/bech32"
	"github.com/terra-project/core/types/util"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	Bech32Prefixes = []string{
                // account's address
                util.Bech32PrefixAccAddr,
                // account's public key
                util.Bech32PrefixAccPub,
                // validator's operator address
                util.Bech32PrefixValAddr,
                // validator's operator public key
                util.Bech32PrefixValPub,
                // consensus node address
                util.Bech32PrefixConsAddr ,
                // consensus node public key
                util.Bech32PrefixConsPub,
        }
)

// Hex Addr -> Ohter Addr
func runFromHex(hexaddr string) [6]string {

	// addr[0]: account's address
	// addr[1]: account's public key
	// addr[2]: validator's operator address
	// addr[3]: validator's operator public key
	// addr[4]: consensus node address
	// addr[5]: consensus node public key -> No tendermint show-validator
	var addr [6]string


	fmt.Println("Hexxxxxxxxxxxxxxxxxxxxxxxxxx: ", hexaddr)

	bz, _ := hex.DecodeString(hexaddr)

	for i, prefix := range Bech32Prefixes {
		bech32Addr, err := bech32.ConvertAndEncode(prefix, bz)

		if err != nil {
			panic(err)
		}

		addr[i] = bech32Addr
	}

	return addr
}

// Bech32 Addr -> Hex Addr
func Bech32AddrToHexAddr(bech32str string) string {
	_, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		fmt.Println("Not a valid bech32 string")
		return "function/keyutil) RunFrombech32() Err"
	}

	return fmt.Sprintf("%X", bz)
}


/*
// Operator Addr -> Ohter Addre
func OperAddrToAccAddr(operaddr string) [6]string {

	hexOperaddr := RunFromBech32(operaddr)
	addr := runFromHex(hexOperaddr)

	return addr
}
*/


func GetAccAddrFromOperAddr(operAddr string) string {

        // Get HexAddress
        hexAddr, err := sdk.ValAddressFromBech32(operAddr)
        if err != nil {
                // Error
        }

        accAddr, err := sdk.AccAddressFromHex(fmt.Sprint(hexAddr))
        if err != nil {
                // Error
        }

        return accAddr.String()
}

