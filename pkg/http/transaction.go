package http

import (
	"context"
	"github.com/rumsrami/givly-rpc-api/pkg/adding"
	"github.com/rumsrami/givly-rpc-api/pkg/listing"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
)

// TODO Define error types to send back to the RPC client
// Map the errors to the DB errors that are captured here

// TransactionDirectory implements twirp TransactionService interface
type TransactionDirectory struct {
	A adding.Service
	L listing.Service
}

// NewTransactionDirectory returns a new configured TransactionDirectory
func NewTransactionDirectory(a adding.Service, l listing.Service) *TransactionDirectory {
	return &TransactionDirectory{
		A: a,
		L: l,
	}
}

// ===================================================
// TransactionDirectory implementation of TransactionService
// ===================================================

// GetItemList gets list of all available products in db
func (t *TransactionDirectory) GetItemList(ctx context.Context, req *transactionPb.ItemListReq) (*transactionPb.ItemList, error) {
	/*
		For demo purpose, hits the db and gets back the list of essential
		products, it is up to the merchant to price them as he wishes
	*/
	// itemList := &transactionPb.ItemList{
	// 	SaleItems: []*transactionPb.SaleItem{
	// 		&transactionPb.SaleItem{
	// 			ItemName: "Banana",
	// 		},
	// 	},
	// }
	items, err := t.L.GetItemList(*req)
	if err != nil {
		return nil, err
	}
	return items, nil
}
// CreateItems adds to the list of available products in the item list
func (t *TransactionDirectory) CreateItems(ctx context.Context, req *transactionPb.ItemList) (*transactionPb.ItemList, error) {
	/*
		For demo purpose, hits the db directly and gets back the list of essential
		products, it is up to the merchant to price them as he wishes
	*/
	// itemList := &transactionPb.ItemList{
	// 	SaleItems: []*transactionPb.SaleItem{
	// 		&transactionPb.SaleItem{
	// 			ItemName: "Banana",
	// 		},
	// 	},
	// }
	items, err := t.A.CreateItems(*req)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// SubmitTx commits the transaction to the database
func (t *TransactionDirectory) SubmitTx(ctx context.Context, req *transactionPb.SubmitTxReq) (*transactionPb.Transaction, error) {
	/*
		Saves the transaction only if it has the reviewed flag
		Time limit of 1 minute after the reviewd flag is checked to true
		Returns an error if time increases above limit
	*/
	return nil, nil
}

// GetRecipientTx gets transactions by a specific recipient
func (t *TransactionDirectory) GetRecipientTx(ctx context.Context, req *transactionPb.TxByRecipientReq) (*transactionPb.TxRes, error) {
	/*
		Gets all transactions from DB related to one recipient.
		formats the response for client
	*/
	return nil, nil
}

// GetMerchantTx gets transactions by a specific merchant
func (t *TransactionDirectory) GetMerchantTx(ctx context.Context, req *transactionPb.TxByMerchantReq) (*transactionPb.TxRes, error) {
	/*
		Gets all transactions from DB related to one merchant.
		formats the response for client
	*/
	return nil, nil
}
