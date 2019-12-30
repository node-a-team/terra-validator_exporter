package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"

//	"github.com/terra-project/core/types/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	core "github.com/terra-project/core/types"
//	"github.com/tendermint/tendermint/libs/bech32"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	cfg "github.com/node-a-team/terra-validator_exporter/config"
	"github.com/node-a-team/terra-validator_exporter/exporter"
	rpc "github.com/node-a-team/terra-validator_exporter/getData/rpc"
)

var ()

func main() {

	log,_ := zap.NewDevelopment()
	defer log.Sync()


	config := sdk.GetConfig()
	config.SetCoinType(core.CoinType)
	config.SetFullFundraiserPath(core.FullFundraiserPath)
	config.SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(core.Bech32PrefixValAddr, core.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(core.Bech32PrefixConsAddr, core.Bech32PrefixConsPub)
	config.Seal()











	cfg.Init()
	rpc.OpenSocket(log)

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start()

//	log.Fatal(http.ListenAndServe(":8080", nil))

	err := http.ListenAndServe(":26661", nil)
	if err != nil {
                // handle error
		log.Fatal("http.ListenAndServe",
		    zap.String("Success", "false"),
		    zap.String("err", fmt.Sprintf("%s", err)),
		)
        }
}
