package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart models.Transaction) (models.Transaction, error)
	GetCartByID(ID int) (models.Transaction, error)
	GetChartByUserID(userID int) ([]models.Transaction, error)
	GetChartByUser(userID int, productID int) (models.Transaction, error)
	UpdateCartQty(Cart models.Transaction, userID int, productID int) (models.Transaction, error)
	DeleteCartByID(Cart models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddToCart(cart models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) GetCartByID(ID int) (models.Transaction, error) {
	var cart models.Transaction
	err := r.db.First(&cart, ID).Preload("User").Preload("Products").Error
	return cart, err
}

func (r *repository) GetChartByUser(userID int, productID int) (models.Transaction, error) {
	var cart models.Transaction
	err := r.db.Preload("Users").Preload("Products").Where("users_id = ? and product_id=?", userID, productID).First(&cart).Error
	return cart, err
}

func (r *repository) GetChartByUserID(userID int) ([]models.Transaction, error) {
	var cart []models.Transaction
	err := r.db.Preload("Users").Preload("Products").Where("users_id=?", userID).Find(&cart).Error
	return cart, err
}

func (r *repository) UpdateCartQty(Cart models.Transaction, userID int, productID int) (models.Transaction, error) {
	err := r.db.Model(&Cart).Where("users_id=? and product_id=?", userID, productID).Updates(&Cart).Error
	return Cart, err
}

func (r *repository) DeleteCartByQty(Cart models.Transaction, userID int, productID int) (models.Transaction, error) {
	err := r.db.Model(&Cart).Where("users_id=? and product_id=?", userID, productID).Updates(&Cart).Error
	return Cart, err
}

func (r *repository) DeleteCartByID(Cart models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&Cart, ID).Error
	return Cart, err
}
