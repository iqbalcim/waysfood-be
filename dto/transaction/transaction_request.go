package transaction

type CreateTransactionRequest struct {
	BuyerID   int    `json:"buyer_id"`
	Status    string `json:"status"  gorm:"type:varchar(25)"`
	ProductID int    `json:"product_id"`
	Qty       int    `json:"qty" `
}

type CreateTransactionbyMany struct {
	BuyerID   int    `json:"buyer_id"`
	Status    string `json:"status"  gorm:"type:varchar(25)"`
	ProductID []int  `json:"product_id" gorm:"type:int"`
	Qty       int    `json:"qty" `
}

type UpdateTransactionRequest struct {
	Status string `json:"status"  gorm:"type:varchar(25)"`
	Qty    int    `json:"qty" `
}
