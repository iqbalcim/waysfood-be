package payment

type CreatePaymentRequest struct {
	BuyerID int `json:"buyer_id"`
	OrderID int `json:"transaction_id"`
}
