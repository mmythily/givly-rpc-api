package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
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

// ServeHTTP implements Handler
func (t TransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Router.ServeHTTP(w, r)
}

// route Mounts the transaction handlers on Router
func (t TransactionHandler) route() {
	t.Router.HandleFunc("/getItemList", t.handleGetItemList()).Methods("POST")
	t.Router.HandleFunc("/createItemList", t.handleCreateItems()).Methods("POST")
	t.Router.HandleFunc("/submitTx", t.handleSubmitTx()).Methods("POST")
	t.Router.HandleFunc("/getTxByRecipient", t.handleGetRecipientTx()).Methods("POST")
	t.Router.HandleFunc("/getTxByMerchant", t.handleGetMerchantTx()).Methods("POST")
}

// handleGetProductList retrieves a list of eligible products
func (t TransactionHandler) handleGetItemList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pbRequest := &transactionPb.ItemListReq{}
		pbResponse, err := t.Client.GetItemList(context.Background(), pbRequest)
		if err != nil {
			fmt.Println(err)
		}
		response, _ := json.Marshal(pbResponse)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// handleCreateItems submits transation to blockchain endpoint
func (t TransactionHandler) handleCreateItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// handleSubmitTx submits transation to blockchain endpoint
func (t TransactionHandler) handleSubmitTx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// handleGetRecipientTx gets transactions by a specific recipient
func (t TransactionHandler) handleGetRecipientTx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// handleGetMerchantTx gets transactions by a specific merchant
func (t TransactionHandler) handleGetMerchantTx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
