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
}

// NewTransactionHandler returns a transaction subrouter
func NewTransactionHandler(addr string) TransactionHandler {
	transactionClient := transactionPb.NewTransactionServiceProtobufClient(addr, &http.Client{})

	return TransactionHandler{
		Client: transactionClient,
	}
}

// Route Mounts the transaction handlers on Router
func (t TransactionHandler) Route(r *mux.Router) {
	r.HandleFunc("/postme", t.handleGetProductList()).Methods("POST")
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
