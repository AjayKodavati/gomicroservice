package main

import (
	"PriceFetcher/service"
	"PriceFetcher/api"
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := service.NewPriceFetcherService()

	server := api.NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}