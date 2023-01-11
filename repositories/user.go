package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ShowUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var users models.User
	err := r.db.First(&users, ID).Error

	return users, err
}




