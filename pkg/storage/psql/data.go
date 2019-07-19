package psql

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	uuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	// Used by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// getMD5Hash hashes a uuid to hex
func getMD5Hash() string {
	hash := md5.Sum([]byte(uuid.New().String()))
	return hex.EncodeToString(hash[:])
}

// Merchant represents a member merchant
type Merchant struct {
	MerchantUUID     string `gorm:"primary_key;unique_index"`
	MerchantCryptoID string `gorm:"not null"`
	StoreEmail       string `gorm:"unique_index;not null"`
	StoreName        string `gorm:"not null"`
	CreatedAt        time.Time
	Transactions     []Transaction `gorm:"foreignkey:MerchantUUID"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (m *Merchant) BeforeCreate(scope *gorm.Scope) error {
	uuid := getMD5Hash()
	return scope.SetColumn("MerchantUUID", uuid)
}

// Product represents one product in a transaction
type Product struct {
	ProductUUID     string  `gorm:"primary_key;"`
	ProductName     string  `gorm:"not null"`
	Price           float32 `gorm:"not null"`
	TransactionUUID string  `gorm:"not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	uuid := getMD5Hash()
	return scope.SetColumn("ProductUUID", uuid)
}

// Transaction represents a sale transaction
type Transaction struct {
	TransactionUUID   string    `gorm:"primary_key;"`
	Products          []Product `gorm:"foreignkey:TransactionUUID"`
	TotalPrice        float32   `gorm:"not null"`
	MerchantUUID      string    `gorm:"not null"`
	RecipientCryptoID string    `gorm:"not null"`
	CreatedAt         time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Transaction) BeforeCreate(scope *gorm.Scope) error {
	uuid := getMD5Hash()
	return scope.SetColumn("TransactionUUID", uuid)
}

// Item represents an item in available Items for purchase
type Item struct {
	ItemUUID string `gorm:"primary_key;"`
	ItemName string `gorm:"not null"`
	ItemURL  string
}

//Items represent mutliples of on item uset to retrieve items
type Items []Item

// BeforeCreate will set a UUID rather than numeric ID.
func (i *Item) BeforeCreate(scope *gorm.Scope) error {
	uuid := getMD5Hash()
	return scope.SetColumn("ItemUUID", uuid)
}

// Recipient represents a beneficiary
type Recipient struct {
	RecipientCryptoID string        `gorm:"primary_key;"`
	Transactions      []Transaction `gorm:"foreignkey:MerchantUUID"`
}
