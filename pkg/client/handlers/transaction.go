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
	t.Router.HandleFunc("/getProductList", t.handleGetProductList()).Methods("POST")
	t.Router.HandleFunc("/review", t.handleReviewTransaction()).Methods("POST")
	t.Router.HandleFunc("/submit", t.handleSubmitTransaction()).Methods("POST")
	t.Router.HandleFunc("/getTransactionsByRT", t.handleGetTransactionsByRT()).Methods("POST")
}

// handleGetProductList retrieves a list of eligible products
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

// handleReviewTransaction checks the transaction value vs recipient balance
func (t TransactionHandler) handleReviewTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// handleSubmitTransaction submits transation to blockchain endpoint
func (t TransactionHandler) handleSubmitTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// handleGetTransactionsByRT gets transactions by a specific recipient
func (t TransactionHandler) handleGetTransactionsByRT() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
