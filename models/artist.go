package models

type Artist struct {
	ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	Name       string `json:"name" gorm:"type : varchar (255)"`
	Old        int    `json:"old" gorm:"type : int"`
	Type       string `json:"type" gorm:"type : varchar (255)"`
	StartCarer string `json:"startcarer" gorm:"type:varchar(255)"`
}

type ArtistResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Old        string `json:"old"`
	Type       string `json:"type"`
	StartCarer string `json:"startcarer"`
}

func (ArtistResponse) TableName() string {
	return "artists"
}
