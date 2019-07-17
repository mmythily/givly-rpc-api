package adding

import (
	uuid "github.com/google/uuid"
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
)

// Service provides Adding operations
// Merchant, Transaction, Item
// Errors here are carried over from the database layer
// Server Handler calls Service layer functions
// Passing in a pbType data Request
// Service returns a pbType data Response
// Service calles storage(db) layer with Data type entity
type Service interface {
	// Add and Get actions on Merchant db model
	CreateMerchant(merchantPb.CreateMerchantReq) (*merchantPb.Merchant, error)
	// Add and Get actions on ItemList db model
	CreateItems(transactionPb.ItemList) (*transactionPb.ItemList, error)
	// Submit will have
	// 1. Get Recipient Balance by Crypto Id
	// 2. Send error if balance < transaction total value
	// 3. AddTransaction on crypto endpoint if balance ok
	// 4. AddTransaction on Transaction db model
	// 5. Send Success if AddTransaction crypto endpoint works
	// Add and Get actions on Transaction db model
	SubmitTx(transactionPb.SubmitTxReq) (*transactionPb.Transaction, error)
}

// Repository provides access to storage layer
// Define the functions that you need the storage(db) layer to provide
// These functions are needed by this adding service
type Repository interface {
	// Add db functions requires to fullfil the service
	// AddMerchant adds a new merchant, creates uuid and returns it
	AddMerchant(merchantPb.CreateMerchantReq) (uuid.UUID, error)
	// GetMerchantByUUID Queries the database for a merchant by his uuid
	GetMerchantByUUID(uuid uuid.UUID) (*merchantPb.Merchant, error)
	// AddItems adds a list of available items
	AddItems(transactionPb.ItemList) error
	// GetItemList gets a list of available items
	GetItemList(transactionPb.ItemListReq) (*transactionPb.ItemList, error)
	// GetBalanceByCryptoID gets balance of recipient
	GetBalanceByCryptoID(merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error)
	// AddTransaction adds a new transaction with uuid signals success
	AddTransaction(transactionPb.SubmitTxReq) (*transactionPb.Transaction, error)
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

func (s *service) CreateMerchant(req merchantPb.CreateMerchantReq) (*merchantPb.Merchant, error) {
	uid, err := s.repo.AddMerchant(req)
	if err != nil {
		return nil, err
	}
	createdMerchant, err := s.repo.GetMerchantByUUID(uid)
	if err != nil {
		return nil, err
	}
	return createdMerchant, nil
}

func (s *service) CreateItems(req transactionPb.ItemList) (*transactionPb.ItemList, error) {
	return nil, nil
}

func (s *service) SubmitTx(req transactionPb.SubmitTxReq) (*transactionPb.Transaction, error) {
	return nil, nil
}
