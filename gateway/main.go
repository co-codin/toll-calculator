package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/co-codin/toll-calculator/aggregator/client"
	"github.com/sirupsen/logrus"
)

func main() {
	listenAddr := flag.String("listenAddr", ":6000", "the listen address of the HTTP server")
	aggregatorServiceAddr := flag.String("aggServiceAddr", "http://localhost:3000", "the listen address of the aggregator service")
	flag.Parse()

	var (
		client     = client.NewHTTPClient(*aggregatorServiceAddr) // endpoint of the aggregator service
		invHandler = newInvoiceHandler(client)
	)

	http.HandleFunc("/invoice", makeAPIFunc(invHandler.handleGetInvoice))
	logrus.Infof("gateway HTTP server running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func handleGetInvoice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("working"))
}
