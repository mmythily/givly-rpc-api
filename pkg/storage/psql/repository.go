package psql

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	// used by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
)

// Storage represents a repository
type Storage struct {
	DB *gorm.DB
}

// New creates a new repository with dependencies
func New(connection string) (*Storage, error) {
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

// Migrate runs the Automigration on the DB
func (s *Storage) Migrate() {
	s.DB.AutoMigrate(&Merchant{}, &Product{}, &Transaction{}, &Item{}, &Recipient{})
	s.DB.Model(&Product{}).AddForeignKey("transaction_uuid", "transaction(transaction_uuid)", "RESTRICT", "RESTRICT")
	s.DB.Model(&Transaction{}).AddForeignKey("merchant_uuid", "merchant(merchant_uuid)", "RESTRICT", "RESTRICT")
	s.DB.Model(&Transaction{}).AddForeignKey("recipient_crypto_id", "recipient(recipient_crypto_id)", "RESTRICT", "RESTRICT")
}

// Close closes the connection
func (s *Storage) Close() {
	s.DB.Close()
}

// AddMerchant adds a new merchant, creates uuid and returns it
func (s *Storage) AddMerchant(req merchantPb.CreateMerchantReq) (*merchantPb.Merchant, error) {
	newMerchant := Merchant{
		MerchantCryptoID: req.MerchantCryptoId,
		StoreEmail: req.StoreEmail,
		StoreName: req.StoreName,
	}
	if err := s.DB.Create(&newMerchant).Error; err != nil {
		return nil, err
	}
	timeStamp, _ := ptypes.TimestampProto(newMerchant.CreatedAt)
	return &merchantPb.Merchant{
		MerchantUuid: newMerchant.MerchantUUID,
		MerchantCryptoId: newMerchant.MerchantCryptoID,
		StoreEmail: newMerchant.StoreEmail,
		CreatedAt: timeStamp,
	}, nil
}

// AddItems adds a list of available items
func (s *Storage) AddItems(req transactionPb.ItemList) (*transactionPb.ItemList, error) {
	fmt.Println(req)
	for _, item := range req.SaleItems {
		dbItem := Item{
			ItemName: item.ItemName,
			ItemURL: item.ItemThumb,
		}
		if err := s.DB.Create(&dbItem).Error; err != nil {
			return nil, err
		}
	}
	return &req, nil
}

// GetItemList gets a list of available items
func (s *Storage) GetItemList(req transactionPb.ItemListReq) (*transactionPb.ItemList, error) {
	var items Items
	if err := s.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	var pbItems transactionPb.ItemList
	for _, item := range items {
		saleItem := transactionPb.SaleItem{
			ItemUuid: item.ItemUUID,
			ItemName: item.ItemName,
			ItemThumb: item.ItemURL,
		}
		pbItems.SaleItems = append(pbItems.SaleItems, &saleItem)
	}
	return &pbItems, nil
}

// GetBalanceByCryptoID gets balance of recipient
func (s *Storage) GetBalanceByCryptoID(req merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error) {
	return nil, nil
}

// AddTransaction adds a new transaction with uuid signals success
func (s *Storage) AddTransaction(req transactionPb.SubmitTxReq) (*transactionPb.Transaction, error) {
	return nil, nil
}

// GetTxByRecipientCryptoID gets transactions for a recipient by crypto id
func (s *Storage) GetTxByRecipientCryptoID(req transactionPb.TxByRecipientReq) (*transactionPb.TxRes, error) {
	return nil, nil
}

// GetTxByMerchantUUID gets transactions for a merchant by crypto id
func (s *Storage) GetTxByMerchantUUID(req transactionPb.TxByMerchantReq) (*transactionPb.TxRes, error) {
	return nil, nil
}

/*
	recipient := &Recipient{
		RecipientCryptoID: "123",
	}
	db.Create(recipient)
	product := &Product{
		ProductName: "Bico",
		Price: 55,
	}
	transaction := &Transaction{
		Products: []Product{*product},
		TotalPrice: 500,
		RecipientCryptoID: "123",
	}
	merchant := &Merchant{
		MerchantCryptoID: "123",
		StoreEmail: "r@r.com",
		StoreName: "Hasouna",
		Transactions: []Transaction{*transaction},
	}
	muuid, _ := uuid.Parse("951e83aa-9dba-403a-b09e-6c0981f1ce52")
	db.Create(transaction)
	db.Create(merchant)
*/
