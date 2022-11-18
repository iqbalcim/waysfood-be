package models

type Transaction struct {
	ID        int                `json:"id" gorm:"primary_key:auto_increment"`
	ProductID int                `json:"product_id" gorm:"type: int"`
	Products  ProductResponse    `json:"product" gorm:"foreignKey:product_id;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	UsersID   int                `json:"user_id"`
	Users     UsersOrderResponse `json:"user" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Qty       int                `json:"qty" form:"qty"`
	Price     int                `json:"price" form:"price"`
}

type Cart struct {
	ID        int                `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID   int                `json:"buyer_id" `
	Buyer     UsersOrderResponse `json:"userOrder"`
	Status    string             `json:"status"  gorm:"type:varchar(25)"`
	Product   []Product          `json:"product"  gorm:"many2many:product_cart"`
	ProductID []int              `json:"product_id" form:"product_id" gorm:"type:int"`
	Qty       int                `json:"qty" `
}

type CartResponse struct {
	ID        int                `json:"id" `
	BuyerID   int                `json:"buyer_id" `
	Buyer     UsersOrderResponse `json:"userOrder" gorm:"foreignKey:BuyerID"`
	Status    string             `json:"status"  gorm:"type:varchar(25)"`
	Product   []Product          `json:"product" gorm:"many2many:product_cart"`
	ProductID []int              `json:"product_id" form:"product_id" gorm:"type:int"`
	Qty       int                `json:"qty" `
}

type TransactionResponse struct {
	ID       int                `json:"id"`
	User     UsersOrderResponse `json:"user"`
	Products ProductResponse    `json:"product`
	Qty      int                `json:"qty"`
	Price    int                `json:"price"`
}

func (CartResponse) TableName() string {
	return "cart"
}
func (TransactionResponse) TableName() string {
	return "cart"
}
