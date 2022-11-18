package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Cart, error)
	UpdateTransaction(transaction models.Cart, ID int) (models.Cart, error)
	DeleteTransaction(transaction models.Cart, ID int) (models.Cart, error)
	CreateTransaction2(transaction models.Cart) (models.Cart, error)
	FindProductById(ProductID []int) ([]models.Product, error)
	GetTransaction2(ID int) (models.Cart, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Cart, error) {
	var transaction []models.Cart
	err := r.db.Preload("Buyer").Preload("Product").Preload("Product.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Cart, ID int) (models.Cart, error) {
	err := r.db.Model(&transaction).Where("id=?", ID).Updates(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Cart, ID int) (models.Cart, error) {
	err := r.db.Preload("Buyer").Preload("Product").Delete(&transaction).Error

	return transaction, err
}

func (r *repository) FindProductById(ProductID []int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products, ProductID).Error

	return products, err
}

func (r *repository) CreateTransaction2(transaction models.Cart) (models.Cart, error) {
	err := r.db.Preload("Buyer").Preload("Product").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction2(ID int) (models.Cart, error) {
	var transaction models.Cart
	err := r.db.Preload("Buyer").Preload("Product").First(&transaction, ID).Error

	return transaction, err
}
