package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	ShowTransactions() ([]models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowTransactions() ([]models.Transaction, error) {
	var Transaction []models.Transaction
	err := r.db.Preload("User").Find(&Transaction).Error

	return Transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {

	err := r.db.Preload("User").Create(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	// var Transaction
	// r.db.Preload("User").First(&transaction)
	err := r.db.Model(&transaction).Updates(transaction).Error
	return transaction, err
	// if status != transaction.Status && status == "success" {
	// 	var transaction models.Transaction
	// 	r.db.First(&transaction, transaction.ID)
	// 	r.db.Save(&transaction)
	// }

	// transaction.Status = status

	// err := r.db.Save(&transaction).Error

	// return transaction, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
