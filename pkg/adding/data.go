package adding

// Merchant represents a member merchant
// This is the data recived from frontend client
type Merchant struct {
	StoreEmail string `schema:"storeEmail"`
	StoreName  string `schema:"storeName"`
	Wallet     string `schema:"wallet"`
}

// Product represents one product in a transaction
// This is the data recived from frontend client
type Product struct {
	ProductName string  `schema:"productName"`
	ProductUnit string  `schema:"productUnit"`
	UnitPrice   float32 `schema:"unitPrice"`
}

// Transaction represents a sale transaction
// This is the data recived from frontend client
type Transaction struct {
	Products    []Product `schema:"products"`
	TotalPrice  float32   `schema:"totalPrice"`
	MerchantID  string    `schema:"merchantId"`
	RecipientID string    `schema:"recipientId"`
}

// ProductList represents a list of available products
type ProductList struct {
	ProductName string `schema:"productName"`
}

// BatchProductList represents batch addition of products
type BatchProductList struct {
	Products []ProductList `schema:"products"`
}
