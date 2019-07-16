package adding

import (
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
)

// Service provides Adding operations
// Merchant, Transaction, Product
// Errors here are carried over from the database layer
// Server Handler calls Service layer functions
// Passing in a pbType data Request
// Service returns a pbType data Response
// Service calles storage(db) layer with Data type entity
type Service interface {
	CreateMerchant(merchantPb.CreateMerchantRequest) (*merchantPb.Merchant, error)
	AddProducts(transactionPb.ProductList) (*transactionPb.ProductList, error)
	SubmitTransaction(transactionPb.Transaction) (*transactionPb.Transaction, error)
	ReviewTransaction(transactionPb.Transaction) (*transactionPb.Transaction, error)
}

// Repository provides access to storage layer
// Define the functions that you need the storage(db) layer to provide
// These functions are needed by this adding service
type Repository interface {
	// Add db functions requires to fullfil the service
}

// service represents an Adding service used by Server Handlers
// Used to interact with entities
type service struct {
	repo Repository
}

// NewService creates an adding service with dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateMerchant(merchantPb.CreateMerchantRequest) (*merchantPb.Merchant, error){
	return nil, nil
}

func (s *service) AddProducts(transactionPb.ProductList) (*transactionPb.ProductList, error) {
	return nil, nil
}

func (s *service) SubmitTransaction(transactionPb.Transaction) (*transactionPb.Transaction, error) {
	return nil, nil
}

func (s *service) ReviewTransaction(transactionPb.Transaction) (*transactionPb.Transaction, error) {
	return nil, nil
}





