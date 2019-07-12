package http

import (
	"context"
	transactionPb "github.com/rumsrami/givly-rpc-api/rpc/transaction"
)

// TransactionDirectory implements twirp TransactionService interface
type TransactionDirectory struct {}

// NewTransactionDirectory returns a new configured TransactionDirectory
func NewTransactionDirectory() (*TransactionDirectory) {
	return &TransactionDirectory{}
}

// ===================================================
// TransactionDirectory implementation of TransactionService
// ===================================================

// GetProductList gets list of all available products in db
func (t *TransactionDirectory) GetProductList(ctx context.Context, req *transactionPb.GetProductListRequest) (*transactionPb.ProductList, error) {
	/*
	For demo purpose, hits the db directly and gets back the list of essential
	products, it is up to the merchant to price them as he wishes
	*/
	product1 := &transactionPb.Product{
		Productname: "Banana",
	}
	product2 := &transactionPb.Product{
		Productname: "Apple",
	}
	transactionList := &transactionPb.ProductList{
		ProductList: []*transactionPb.Product{product1, product2},
	}
	return transactionList, nil
}

// ReviewTransaction validates an ongoing transaction 
func (t *TransactionDirectory) ReviewTransaction(ctx context.Context, req *transactionPb.ReviewTransactionRequest) (*transactionPb.Transaction, error) {
	/*
	Hits the crypto endpoint to get recipient balance and calculated the sum price
	of all items in the transaction, if balance permits, adds a flag to the transaction and
	returns it, otherwise returns error
	*/
	return nil, nil
}

// SubmitTransaction commits the transaction to the database
func (t *TransactionDirectory) SubmitTransaction(ctx context.Context, req *transactionPb.SubmitTransactionRequest) (*transactionPb.Transaction, error) {
	/*
	Saves the transaction only if it has the reviewed flag
	Time limit of 1 minute after the reviewd flag is checked to true
	Returns an error if time increases above limit
	*/
	return nil, nil
}