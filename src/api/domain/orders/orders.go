package orders

type Order struct {
    Id 			int64  	 `json:"id"`
    DateCreated string 	 `json:"date_created"`
    Items 		[]string `json:"order_items"`
    Address		string	 `json:"address"`
    TotalAmount	float32  `json:"total_amount"`
    Status		string 	 `json:"status"`
    CurrencyId	string 	 `json:"currency_id"`
    Tags 		[]string `json:"tags"`
    Buyer 		string 	 `json:"buyer"`
    Seller 		string 	 `json:"seller"`
    Payments 	[]string `json:"payments"`
}