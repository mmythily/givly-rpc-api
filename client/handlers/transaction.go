package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	transactionPb "github.com/rumsrami/givly-rpc-api/rpc/transaction"
)

// TransactionHandler handles transaction client requests
type TransactionHandler struct {
	Client transactionPb.TransactionService
	Router *mux.Router
}

// NewTransactionHandler returns a transaction handler
func NewTransactionHandler(addr string, r *mux.Router) TransactionHandler {
	transactionClient := transactionPb.NewTransactionServiceProtobufClient(addr, &http.Client{})
	transactionHandler := TransactionHandler{
		Client: transactionClient,
		Router: r,
	}
	transactionHandler.route()
	return transactionHandler
}

func (t TransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Router.ServeHTTP(w,r)
}

// route Mounts the transaction handlers on Router
func (t TransactionHandler) route() {
	t.Router.HandleFunc("/postme", t.handleGetProductList()).Methods("POST")
}

func (t TransactionHandler) handleGetProductList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pbRequest := &transactionPb.GetProductListRequest{}
		pbResponse, err := t.Client.GetProductList(context.Background(), pbRequest)
		if err != nil {
			fmt.Println(err)
		}
		response, _ := json.Marshal(pbResponse)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
