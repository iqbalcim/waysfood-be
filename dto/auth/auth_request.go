package authdto

type RegisterRequest struct {
	Name     string `json:"name" form:"name" `
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password" `
	Phone    string `json:"phone" form:"phone" `
	Gender   string `json:"gender" form:"gender" `
	Location string `json:"location" form:"location"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
