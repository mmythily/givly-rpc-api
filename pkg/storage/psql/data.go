package psql

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/google/uuid"
	// Used by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Merchant represents a member merchant
type Merchant struct {
	MerchantUUID     uuid.UUID `gorm:"type:uuid;primary_key;unique_index"`
	MerchantCryptoID string    `gorm:"not null"`
	StoreEmail       string    `gorm:"unique_index;not null"`
	StoreName        string    `gorm:"not null"`
	CreatedAt        time.Time
	Transactions     []Transaction `gorm:"foreignkey:MerchantUUID"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (m *Merchant) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("MerchantUUID", uuid)
}

// Product represents one product in a transaction
type Product struct {
	ProductUUID uuid.UUID `gorm:"type:uuid;primary_key;"`
	ProductName string    `gorm:"not null;unique_index"`
	Price       float32   `gorm:"not null"`
	TransactionUUID uuid.UUID `gorm:"not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("ProductUUID", uuid)
}

// Transaction represents a sale transaction
type Transaction struct {
	TransactionUUID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Products          []Product	`gorm:"foreignkey:TransactionUUID"`
	TotalPrice        float32   `gorm:"not null"`
	MerchantUUID      uuid.UUID `gorm:"not null"`
	RecipientCryptoID string    `gorm:"not null"`
	CreatedAt        time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Transaction) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("TransactionUUID", uuid)
}

// Item represents an item in available Items for purchase
type Item struct {
	ItemUUID uuid.UUID `gorm:"type:uuid;primary_key;"`
	ItemName string    `gorm:"not null"`
	ItemURL  string
}

// BeforeCreate will set a UUID rather than numeric ID.
func (i *Item) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("ItemUUID", uuid)
}

// Recipient represents a beneficiary
type Recipient struct {
	RecipientCryptoID string `gorm:"primary_key;"`
	Transactions      []Transaction `gorm:"foreignkey:MerchantUUID"`
}
