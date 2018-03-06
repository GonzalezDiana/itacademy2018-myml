package payments

type Payment struct {
    Id                int64   `json:"id"`
    OrderId           int     `json:"order_id"`
    PaymentMethodId   string  `json:"payment_method_id"`
    CurrencyId        string  `json:"currency_id"`
    Status            string  `json:"status"`
    TransactionAmount float32 `json:"transaction_amount"`
}
