package handlers

/*
JS code to produce this object
var itemList = {
  itemList: [
    {
      itemName: "Banana",
      itemUrl: "https://www.amazons3.com"
    }
  ]
}
console.log((JSON.stringify(itemList)))
and then name the outcome -> "itemList"
send it as:
"itemList":{"itemList":[{"itemName":"Rami","itemUrl":"www"}]}
*/

// Items represent new itemList addition request
type Items struct {
	ItemList []struct {
		ItemUUID string `json:"itemUuid"`
		ItemName string `json:"itemName"`
		ItemURL  string `json:"itemUrl"`
	} `json:"itemList"`
}

/*
var transaction = {
  totalPrice: 1,    /// Cannot be 0
  merchantUuid: "123",
  recipientCryptoId: "123",
  products: [
    {
      productName: "Banana",
      price: 1 	/// cannot be 0
    },
    {
      productName: "apple",
      price: 1	/// cannot be 0
    }
  ]
}

console.log((JSON.stringify(transaction)))
*/

// Transaction represents a new client Transaction
type Transaction struct {
	TotalPrice        float32 `json:"totalPrice"`
	MerchantUUID      string  `json:"merchantUuid"`
	RecipientCryptoID string  `json:"recipientCryptoId"`
	Products          []struct {
		ProductName string  `json:"productName"`
		Price       float32 `json:"price"`
	} `json:"products"`
}
