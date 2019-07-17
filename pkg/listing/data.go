package listing

import (
	uuid "github.com/google/uuid"
)

// Merchant represents a member merchant
// This is the data returned to the frontend client
type Merchant struct {
	MerchantUUID uuid.UUID `json:"merchantUuid"`
	StoreEmail   string    `json:"storeEmail"`
	StoreName    string    `json:"storeName"`
}

// Product represents one product in a transaction
// This is the data recived from frontend client
type Product struct {
	ProductUUID uuid.UUID `json:"productUuid"`
	ProductName string    `json:"productName"`
	Price       float32   `json:"price"`
}

// Transaction represents a sale transaction
// This is the data recived from frontend client
type Transaction struct {
	TransactionUUID   uuid.UUID `json:"transactionUuid"`
	Products          []Product `json:"products"`
	TotalPrice        float32   `json:"totalPrice"`
	MerchantUUID      uuid.UUID `json:"merchantUuid"`
	RecipientCryptoID string    `json:"recipientCryptoId"`
}

// Item represents an item in available Items for purchase
type Item struct {
	ItemUUID uuid.UUID `json:"itemUuid"`
	ItemName string    `json:"itemName"`
	ItemURL  string    `json:"itemUrl"`
}

// ItemList represents batch addition of Items
type ItemList struct {
	ItemList []Item `json:"itemList"`
}

// RecipientReq representa a beneficiary
type RecipientReq struct {
	RecipientCryptoID string `schema:"recipientCryptoId"`
}

// MerchantReq representa a merchant
type MerchantReq struct {
	MerchantUUID uuid.UUID `schema:"merchantUuid"`
}
