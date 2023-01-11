package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type ArtistRepository interface {
	ShowArtists() ([]models.Artist, error)
	GetArtist(ID int) (models.Artist, error)
	AddArtist(artist models.Artist) (models.Artist, error)
}

func (r *repository) AddArtist(artist models.Artist) (models.Artist, error) {
	err := r.db.Create(&artist).Error

	return artist, err
}

func RepositoryArtist(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := r.db.Find(&artists).Error
	return artists, err
}

func (r *repository) GetArtist(ID int) (models.Artist, error) {
	var artist models.Artist
	err := r.db.First(&artist, ID).Error

	return artist, err
}
