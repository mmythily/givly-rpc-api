package listing



// // Merchant represents a member merchant
// // This is the data returned to the frontend client
// type Merchant struct {
// 	MerchantUUID string `json:"merchantUuid"`
// 	StoreEmail   string    `json:"storeEmail"`
// 	StoreName    string    `json:"storeName"`
// }

// // Product represents one product in a transaction
// // This is the data recived from frontend client
// type Product struct {
// 	ProductUUID string `json:"productUuid"`
// 	ProductName string    `json:"productName"`
// 	Price       float32   `json:"price"`
// }

// // Transaction represents a sale transaction
// // This is the data recived from frontend client
// type Transaction struct {
// 	TransactionUUID   string `json:"transactionUuid"`
// 	Products          []Product `json:"products"`
// 	TotalPrice        float32   `json:"totalPrice"`
// 	MerchantUUID      string `json:"merchantUuid"`
// 	RecipientCryptoID string    `json:"recipientCryptoId"`
// }

// // Item represents an item in available Items for purchase
// type Item struct {
// 	ItemUUID string `json:"itemUuid"`
// 	ItemName string    `json:"itemName"`
// 	ItemURL  string    `json:"itemUrl"`
// }

// // ItemList represents batch addition of Items
// type ItemList struct {
// 	ItemList []Item `json:"itemList"`
// }

// // RecipientReq representa a beneficiary
// type RecipientReq struct {
// 	RecipientCryptoID string `schema:"recipientCryptoId"`
// }

// // MerchantReq representa a merchant
// type MerchantReq struct {
// 	MerchantUUID string `schema:"merchantUuid"`
// }
