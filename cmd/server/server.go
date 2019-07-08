package main

import (
	"net/http"
	"github.com/rumsrami/givly-rpc-api/internal"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	merchantHandler := internal.NewMerchantDirectoryHandler()
	transactionHandler := internal.NewTransactionDirectoryHandler()

	r.Handle("/api/merchant", merchantHandler)
	r.Handle("/api/transaction", transactionHandler)

	http.ListenAndServe(":8080", r)
}