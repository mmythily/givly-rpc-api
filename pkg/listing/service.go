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
	// GetBalanceByCryptoId on Recipient db Model
	GetRecipientBalance(merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error)
	// GetItemList on the ItemList db model
	GetItemList(transactionPb.ItemListReq) (*transactionPb.ItemList, error)
	// GetTxByRecipientCryptoId
	GetRecipientTx(transactionPb.TxByRecipientReq) (*transactionPb.TxRes, error)
	// GetTxByMerchantUUID
	GetMerchantTx(transactionPb.TxByMerchantReq) (*transactionPb.TxRes, error)
}

// Repository provides access to storage layer
// Define the functions that you need the storage(db) layer to provide
// These functions are needed by this listing service
type Repository interface {
	// Add db functions requires to fullfil the service
	// GetRecipientBalance gets balance of recipient
	GetBalanceByCryptoID(merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error)
	// GetItemList gets a list of available items
	GetItemList(transactionPb.ItemListReq) (*transactionPb.ItemList, error)
	// GetTxByRecipientCryptoID
	GetTxByRecipientCryptoID(transactionPb.TxByRecipientReq) (*transactionPb.TxRes, error)
	// GetTxByMerchantUUID
	GetTxByMerchantUUID(transactionPb.TxByMerchantReq) (*transactionPb.TxRes, error)
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


func (s *service) GetRecipientBalance(req merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error){
	return nil, nil
}

func (s *service) GetItemList(transactionPb.ItemListReq) (*transactionPb.ItemList, error){
	return nil, nil
}

func (s *service)GetRecipientTx(transactionPb.TxByRecipientReq) (*transactionPb.TxRes, error){
	return nil, nil
}

func (s *service) GetMerchantTx(req transactionPb.TxByMerchantReq) (*transactionPb.TxRes, error){
	return nil, nil
}
