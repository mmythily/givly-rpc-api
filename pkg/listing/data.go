package listing

// Merchant represents a member merchant
// This is the data recived from frontend client
type Merchant struct {
	MerchantID   string        `json:"merchantId"`
	StoreEmail   string        `json:"storeEmail"`
	StoreName    string        `json:"storeName"`
	Transactions []Transaction `json:"transactions"`
}

// Product represents one product in a transaction
// This is the data recived from frontend client
type Product struct {
	ProductName string  `json:"productName"`
	ProductUnit string  `json:"productUnit"`
	UnitPrice   float32 `json:"unitPrice"`
}

// Transaction represents a sale transaction
// This is the data recived from frontend client
type Transaction struct {
	TransactionID string `json:"transactionId"`
	Products    []Product `json:"products"`
	TotalPrice  float32   `json:"totalPrice"`
	MerchantID  string    `json:"merchantId"`
	RecipientID string    `json:"recipientId"`
}

// ProductList represents a list of available products
type ProductList struct {
	ProductName string `json:"productName"`
}

// BatchProductList represents batch addition of products
type BatchProductList struct {
	Products []ProductList `json:"products"`
}

// Recipient representa a beneficiary
type Recipient struct {
	RecipientID string `json:"recipientId"`
}
