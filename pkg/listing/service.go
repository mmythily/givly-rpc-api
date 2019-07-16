package listing

import (
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
)

// Service provides listing operations
// Merchant, Transaction, Product
// Errors here are carried over from the database layer
// Server Handler calls Service layer functions
// Passing in a pbType data Request
// Service returns a pbType data Response
// Service calles storage(db) layer with Data type entity
type Service interface {
	UpdateMerchant(merchantPb.UpdateMerchantRequest) (*merchantPb.Merchant, error)
	VerifyRecipientAccount(merchantPb.VerifyRTAccountRequest) (*merchantPb.RTAccountBalanceVerified, error)
	GetProductList(transactionPb.GetProductListRequest) (*transactionPb.ProductList, error)
	GetTransactionsByRecipient(transactionPb.GetTransactionsByRTRequest) (*transactionPb.TransactionsByRTResponse, error)
}

// Repository provides access to storage layer
// Define the functions that you need the storage(db) layer to provide
// These functions are needed by this listing service
type Repository interface {
	// Add db functions requires to fullfil the service
}

// service represents a Listing service used by Server Handlers
// Used to interact with entities
type service struct {
	repo Repository
}

// NewService creates a Listing service with dependencies
func NewService(r Repository) Service {
	return &service{r}
}


func (s *service)UpdateMerchant(merchantPb.UpdateMerchantRequest) (*merchantPb.Merchant, error){
	return nil, nil
}

func (s *service)VerifyRecipientAccount(merchantPb.VerifyRTAccountRequest) (*merchantPb.RTAccountBalanceVerified, error){
	return nil, nil
}

func (s *service)GetProductList(transactionPb.GetProductListRequest) (*transactionPb.ProductList, error){
	return nil, nil
}

func (s *service)GetTransactionsByRecipient(transactionPb.GetTransactionsByRTRequest) (*transactionPb.TransactionsByRTResponse, error){
	return nil, nil
}
