package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
)

// txRPCProcessor represents the rpc function caller
type txRPCProcessor func(*http.Request, transactionPb.TransactionService) (proto.Message, error)

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
	t.Router.HandleFunc("/getItemList", t.respond(getItemList, handleError)).Methods("POST")
	t.Router.HandleFunc("/addItems", t.respond(createItems, handleError)).Methods("POST")
	t.Router.HandleFunc("/submitTx", t.respond(submitTx, handleError)).Methods("POST")
	t.Router.HandleFunc("/getTxByRecipient", t.respond(getRecipientTx, handleError)).Methods("POST")
	t.Router.HandleFunc("/getTxByMerchant", t.respond(getMerchantTx, handleError)).Methods("POST")
}

// respond wraps the response with headers and logging
func (t *TransactionHandler) respond(process txRPCProcessor, format errFormater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pbResponse, err := process(r, t.Client)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			// Change status to match twirp error TODO
			w.WriteHeader(http.StatusNotFound)
			w.Write(format(err))
		}
		// Marshall the response
		response, _ := json.Marshal(pbResponse)
		// Send back to Client
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

/*
	txRPCProcessors definitions
*/

// getItemList retrieves a list of eligible products
func getItemList(r *http.Request, pb transactionPb.TransactionService) (proto.Message, error) {
	pbRequest := &transactionPb.ItemListReq{}
	pbResponse, err := pb.GetItemList(context.Background(), pbRequest)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

// createItems submits transation to blockchain endpoint
func createItems(r *http.Request, pb transactionPb.TransactionService) (proto.Message, error) {
	// Parse incoming JSON
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}
	// Marshall JSON request
	var items Items
	err := json.Unmarshal([]byte(r.FormValue("itemList")), &items)
	if err != nil {
		return nil, err
	}
	// Create protobuf request
	pbRequest := &transactionPb.ItemList{
		SaleItems: []*transactionPb.SaleItem{},
	}
	for _, item := range items.ItemList {
		pbRequest.SaleItems = append(pbRequest.SaleItems, &transactionPb.SaleItem{
			ItemName:  item.ItemName,
			ItemThumb: item.ItemURL,
		})
	}
	// Call RPC function and get protobuf response
	pbResponse, err := pb.CreateItems(context.Background(), pbRequest)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

// submitTx submits transation to blockchain endpoint
func submitTx(r *http.Request, pb transactionPb.TransactionService) (proto.Message, error) {
	// Parse incoming JSON
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}
	// Marshall JSON request
	var newTransaction Transaction
	err := json.Unmarshal([]byte(r.FormValue("transaction")), &newTransaction)
	if err != nil {
		return nil, err
	}
	var txProducts []*transactionPb.Product
	for _, product := range newTransaction.Products {
		txProducts = append(txProducts, &transactionPb.Product{
			ProductName: product.ProductName,
			Price: product.Price,
		})
	}
	// Create protobuf request
	pbRequest := transactionPb.Transaction{
		TotalPrice: newTransaction.TotalPrice,
		MerchantUuid: newTransaction.MerchantUUID,
		RecipientCryptoId: newTransaction.RecipientCryptoID,
		Products: txProducts,
	}
	// Call RPC function and get protobuf response
	pbResponse, err := pb.SubmitTx(context.Background(), &transactionPb.SubmitTxReq{
		Transaction: &pbRequest,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(pbResponse)
	return pbResponse, nil
}

// getRecipientTx gets transactions by a specific recipient
func getRecipientTx(r *http.Request, pb transactionPb.TransactionService) (proto.Message, error) {
	// Create protobuf request
	// Call RPC function and get protobuf response
	// Marshall the response
	// Send back to Client
	return nil, nil
}

// getMerchantTx gets transactions by a specific merchant
func getMerchantTx(r *http.Request, pb transactionPb.TransactionService) (proto.Message, error) {
	// Create protobuf request
	// Call RPC function and get protobuf response
	// Marshall the response
	// Send back to Client
	return nil, nil
}
