package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rumsrami/givly-rpc-api/client/handlers"
)

func main() {
	r := mux.NewRouter()
	serviceRouter := mux.NewRouter()
	merchantRouter := serviceRouter.PathPrefix("/merchant").Subrouter()
	transactionRouter := serviceRouter.PathPrefix("/transaction").Subrouter()
	
	r.PathPrefix("/merchant/").Handler(handlers.NewMerchantHandler("http://localhost:8080", merchantRouter))
	r.PathPrefix("/transaction/").Handler(handlers.NewTransactionHandler("http://localhost:8080", transactionRouter))
	http.ListenAndServe(":8000", r)
}
