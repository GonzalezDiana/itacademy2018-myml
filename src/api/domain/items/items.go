package items

type Item struct {
    Id 				int64  		`json:"id"`
    SiteId			string 		`json:"site_id"`
    Title 			string 		`json:"title"`
    CategoryId 		string 		`json:"category_id"`
    Price 			float32 	`json:"price"`
    CurrencyId 		string 		`json:"currency_id"`
    Quantity		int32 		`json:"quantity"`
    Condition 		string 		`json:"condition"`
    Pictures		[]string 	`json:"pictures"`
}
