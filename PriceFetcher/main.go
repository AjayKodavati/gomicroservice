package main

import (
	"PriceFetcher/api"
	// "PriceFetcher/client"
	"PriceFetcher/service"
	// "context"
	"flag"
	"fmt"
)

func main() {
	// testing http client
	// client := client.NewgRpcClient("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "NSE")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)

	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()
	
	svc := service.NewPriceFetcherService()

	server := api.NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}