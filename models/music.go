package models

type Music struct {
	ID        int            `json:"id" gorm:"primary_key:auto_increment"`
	Title     string         `json:"title" gorm:"type : varchar (255)"`
	Year      string         `json:"year" gorm:"type : int"`
	Thumbnail string         `json:"thumbnail" gorm:"type : varchar (255)"`
	Attache   string         `json:"attache" gorm:"type:varchar(255)"`
	Artist_id int            `json:"artist_id" gorm:"int"`
	Artist    ArtistResponse `json:"artist" `
}

type MusicResponse struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Year      string         `json:"year"`
	Thumbnail string         `json:"thumbnail"`
	Attache   string         `json:"attache"`
	Artist_id int            `json:"artist_id"`
	Artist    ArtistResponse `json:"artist" `
}

func (MusicResponse) TableName() string {
	return "Musics"
}
