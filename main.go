package main

import (
	"fmt"
	"net/http"
	"os"
	"go.uber.org/zap"

	sdk "github.com/cosmos/cosmos-sdk/types"
	core "github.com/terra-project/core/types"

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


	cfg.ConfigPath = os.Args[1]

	port := cfg.Init()
	rpc.OpenSocket(log)

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start(log)

	err := http.ListenAndServe(":" +port, nil)
	// log
        if err != nil {
                // handle error
                log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: " +port +")"),)
        }

}
