package rpc

import (
//        "log"
	"fmt"
	"go.uber.org/zap"

        tmclient "github.com/tendermint/tendermint/rpc/client"
//	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)


type RPCData struct {
        Commit	commitInfo
}

var (
        Addr string

        Client *tmclient.HTTP
)

func newRPCData() *RPCData {

        rd := &RPCData {
		//
        }

        return rd
}

func GetData(blockHeight int64, consHexAddr string) *RPCData {

	rd := newRPCData()

	var commitHeight int64 = blockHeight -1
	commitData, _ := Client.Commit(&commitHeight)

	rd.Commit = getCommit(commitData, consHexAddr)


        return rd
}

func OpenSocket(log *zap.Logger) {

        Client = tmclient.NewHTTP("tcp://"+Addr, "/websocket")

        err := Client.Start()
        if err != nil {
                // handle error
		log.Fatal("OpenSocket",
		    zap.String("Success", "false"),
		    zap.String("err", fmt.Sprintf("%s", err)),
		)
        }
        defer Client.Stop()

	log.Info("RPC Server Connect",
                    zap.Bool("Success", true),
		    zap.String("err", "nil"),
        )
}

func BlockHeight() (res int64) {

	info, _ := Client.ABCIInfo()
	res = info.Response.LastBlockHeight

	return res
}


