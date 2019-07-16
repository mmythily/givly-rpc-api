package psql

import (
	"github.com/jinzhu/gorm"
	// Blank import for the use of gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Merchant is gorm psql model
type Merchant struct {
	gorm.Model
	StoreEmail   string `gorm:"not null;unique_index"`
	StoreName    string `gorm:"not null"`
	Transactions []Transaction
}

// ProductList is a list of available products
type ProductList struct {
	gorm.Model
	Productname string `gorm:"not null;unique_index"`
}

// Product is gorm psql model
type Product struct {
	gorm.Model
	Productname string `gorm:"not null"`
	Productunit string
	UnitPrice   float32
}

// Recipient represents a service beneficiary
type Recipient struct {
	gorm.Model
	Transactions []Transaction
}

// Transaction represents a sale transaction
type Transaction struct {
	gorm.Model
	TotalPrice float32
	Products   []Product
}
