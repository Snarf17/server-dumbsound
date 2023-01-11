package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	ShowTransactions() ([]models.Transaction, error)
	AddTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID int) (models.Transaction, error)
	// DeleteTransaksi(transaksi models.Transaction) (models.Transaction, error)
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

func (r *repository) AddTransaction(transaction models.Transaction) (models.Transaction, error) {

	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var transaction models.Transaction
		r.db.First(&transaction, transaction.ID)
		r.db.Save(&transaction)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return transaction, err
}

// func (r *repository) DeleteTransaksi(transaksi models.Transaction) (models.Transaction, error) {
// 	err := r.db.Delete(&transaksi).Error

//		return transaksi, err
//	}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
