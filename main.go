package main

import (
	"log"
	"net/http"

        setting "github.com/node-a-team/terra-validator_exporter/function/setting"
        exporter "github.com/node-a-team/terra-validator_exporter/function/exporter"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (


)



func main() {

        var port string
        port = "8888"

        http.Handle("/metrics", promhttp.Handler())

	setting.Init()

	go exporter.Exporter()

	 log.Fatal(http.ListenAndServe(":"+port, nil))
 }
