package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ShowUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUsers(user models.User) (models.User, error)
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

func (r *repository) UpdateUsers(user models.User) (models.User, error) {
	// var user models.User
	err := r.db.Save(&user).Error

	return user, err
}
