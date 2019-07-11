package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rumsrami/givly-rpc-api/client"
)

func main() {
	r := mux.NewRouter()
	client := client.New("http://localhost:8080")
	merchantRouter := r.PathPrefix("/merchant").Subrouter()
	transactionRouter := r.PathPrefix("/transaction").Subrouter()
	client.MerchantHandler.Route(merchantRouter)
	client.TransactionHandler.Route(transactionRouter)
	http.ListenAndServe(":8000", r)
}
