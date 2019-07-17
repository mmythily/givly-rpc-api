package psql

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	// used by gorm
  _ "github.com/jinzhu/gorm/dialects/postgres"
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
	uuid "github.com/google/uuid"
)

// Storage represents a repository
type Storage struct {
	DB *gorm.DB
}

// New creates a new repository with dependencies
func New(connection string) (*Storage, error){
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.LogMode(true)
	return &Storage{
		DB: db,
	}, nil
}

// Close closes the connection
func (s *Storage) Close() {
	s.DB.Close()
}

// AddMerchant adds a new merchant, creates uuid and returns it
func (s *Storage) AddMerchant(merchantPb.CreateMerchantReq) (uuid.UUID, error) {
	return uuid.New(), nil
}

// GetMerchantByUUID Queries the database for a merchant by his uuid
func (s *Storage) GetMerchantByUUID(uuid uuid.UUID) (*merchantPb.Merchant, error) {
	return nil, nil
}

// AddItems adds a list of available items
func (s *Storage) AddItems(transactionPb.ItemList) error {
	return nil
}

// GetItemList gets a list of available items
func (s *Storage) GetItemList(transactionPb.ItemListReq) (*transactionPb.ItemList, error) {
	return nil, nil
}

// GetBalanceByCryptoID gets balance of recipient
func (s *Storage) GetBalanceByCryptoID(merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error) {
	return nil, nil
}

// AddTransaction adds a new transaction with uuid signals success
func (s *Storage) AddTransaction(transactionPb.SubmitTxReq) (*transactionPb.Transaction, error) {
	return nil, nil
}

// GetTxByRecipientCryptoID gets transactions for a recipient by crypto id
func (s *Storage) GetTxByRecipientCryptoID(transactionPb.TxByRecipientReq) (*transactionPb.TxRes, error) {
	return nil, nil
}

// GetTxByMerchantUUID gets transactions for a merchant by crypto id
func (s *Storage) GetTxByMerchantUUID(transactionPb.TxByMerchantReq) (*transactionPb.TxRes, error) {
	return nil, nil
}

