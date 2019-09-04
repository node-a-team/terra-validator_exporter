package setting

import (

        t "github.com/node-a-team/terra-validator_exporter/types"

        "fmt"
        "time"
)


func Init() {

        rpcServer, restServer, network, operatorAddr := "localhost:26657", "localhost:1317", "terra", ""


        fmt.Println("##### Symple Tendermint Explorer #####")

        fmt.Println("\nYou need an RPC-Server and a REST-Server")
        fmt.Printf("RPC-Server(default-> localhost:26657): ")
        fmt.Scanf("%s", &rpcServer)
        t.RpcServer = rpcServer

        fmt.Printf("Rest-Server(default-> localhost:1317 ): ")
        fmt.Scanf("%s", &restServer)
        t.RestServer = restServer

        fmt.Printf("\nNetwork(default-> terra): ")
        fmt.Scanf("%s", &network)
        t.Bech32MainPrefix = network

	fmt.Printf("\nValidator OperatorAddr(ex: gaiacli keys show KEY --bech=val): ")
        fmt.Scanf("%s", &operatorAddr)
        t.OperatorAddr = operatorAddr


        fmt.Println("\nYour RPC-Server: ", t.RpcServer)
        fmt.Println("Your Rest-Server: ", t.RestServer)
        fmt.Println("Your Bech32MainPrefix: ", t.Bech32MainPrefix)
        fmt.Println("Your OperatorAddr: ", t.OperatorAddr)

        time.Sleep(2*time.Second)
}

