package utils

import (
	"fmt"
	"go.uber.org/zap"

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

// Bech32 Addr -> Hex Addr
func Bech32AddrToHexAddr(bech32str string, log *zap.Logger) string {
	_, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "Bech32Addr To HexAddr"),)
        }

	return fmt.Sprintf("%X", bz)
}

func GetAccAddrFromOperAddr(operAddr string, log *zap.Logger) string {

        // Get HexAddress
        hexAddr, err := sdk.ValAddressFromBech32(operAddr)
	// log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "OperAddr To HexAddr"),)
        }

        accAddr, err := sdk.AccAddressFromHex(fmt.Sprint(hexAddr))
	// log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "HexAddr To AccAddr"),)
        }

        return accAddr.String()
}

