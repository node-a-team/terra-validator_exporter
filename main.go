package main

import (
	"log"
	"net/http"

	config "github.com/node-a-team/terra-validator_exporter/function/config"
	exporter "github.com/node-a-team/terra-validator_exporter/function/exporter"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ()

func main() {

	var port string
	port = "8888"

	http.Handle("/metrics", promhttp.Handler())

	config.Init()

	go exporter.Exporter()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
