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

// Items represent new item list
type Items struct {
	ItemList []struct {
		ItemUUID string `json:"itemUuid"` 
		ItemName string `json:"itemName"`
		ItemURL  string `json:"itemUrl"`
	} `json:"itemList"`
}
