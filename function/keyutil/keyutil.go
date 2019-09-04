package keyutil

import (
        "encoding/hex"
        "fmt"

        t "github.com/node-a-team/terra-validator_exporter/types"
        "github.com/tendermint/tendermint/libs/bech32"
)

// Print info from HEX
func RunFromHex(hexaddr string) [6]string {

	var bech32Prefixes = []string{

	        // account's address
	        t.Bech32MainPrefix,
	        // account's public key
	        t.Bech32MainPrefix + t.PrefixPublic,
	        // validator's operator address
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixOperator,
	        // validator's operator public key
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixOperator + t.PrefixPublic,
	        // consensus node address
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixConsensus,
	        // consensus node public key
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixConsensus + t.PrefixPublic,
	}


	// keys[0]: account's address
	// keys[1]: account's public key
	// keys[2]: validator's operator address
	// keys[3]: validator's operator public key
	// keys[4]: consensus node address
	// keys[5]: consensus node public key -> No tendermint show-validator
	var keys [6]string

        bz, _ := hex.DecodeString(hexaddr)

        for i, prefix := range bech32Prefixes {
                bech32Addr, err := bech32.ConvertAndEncode(prefix, bz)

                if err != nil {
                        panic(err)
                }

		keys[i] = bech32Addr
        }

	return keys
}


// Print info from bech32.
func RunFromBech32(bech32str string) string{
        _, bz, err := bech32.DecodeAndConvert(bech32str)
        if err != nil {
                fmt.Println("Not a valid bech32 string")
                return "function/keyutil) RunFrombech32() Err"
        }

        return fmt.Sprintf("%X", bz)
}


// Operator address -> Ohter address
func OperAddrToOtherAddr(operaddr string) [6]string {

        hexOperaddr := RunFromBech32(operaddr)
        keys := RunFromHex(hexOperaddr)

        return keys
}






/*
//////////////////////////////////////////////////////////////////////////////////
// Test

func KeyUtilTest(address string) {
        if len(address) < 2 {
                fmt.Println("Must specify an input string")
        }

        runFromBech32Test(address)
        runFromHexTest(address)
}


// Print info from bech32.
func runFromBech32Test(bech32str string) {
        hrp, bz, err := bech32.DecodeAndConvert(bech32str)
        if err != nil {
                fmt.Println("Not a valid bech32 string")
                return
        }
        fmt.Println("Bech32 parse:")
        fmt.Printf("Human readible part: %v\nBytes (hex): %X\n",
                hrp,
                bz,
        )
}

// Print info from HEX
func runFromHexTest(hexaddr string) {
        bz, err := hex.DecodeString(hexaddr)
        if err != nil {
                fmt.Println("Not a valid hex string")
                return
        }
        fmt.Println("Hex parse:")
        fmt.Println("Bech32 formats:")
        for _, prefix := range bech32Prefixes {
                bech32Addr, err := bech32.ConvertAndEncode(prefix, bz)
                if err != nil {
                        panic(err)
                }
                fmt.Println("  - " + bech32Addr)
        }
}

//////////////////////////////////////////////////////////////////////////////////
*/
