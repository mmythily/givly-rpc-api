package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rumsrami/givly-rpc-api/internal"
	merchantPb "github.com/rumsrami/givly-rpc-api/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/rpc/transaction"
)

func main() {
	r := mux.NewRouter()

	merchantHandler := internal.NewMerchantDirectoryHandler()
	transactionHandler := internal.NewTransactionDirectoryHandler()

	r.PathPrefix(merchantPb.MerchantServicePathPrefix).Handler(merchantHandler)
	r.PathPrefix(transactionPb.TransactionServicePathPrefix).Handler(transactionHandler)

	http.ListenAndServe(":8080", r)
}
