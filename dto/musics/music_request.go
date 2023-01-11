package dtomusic

type MusicRequest struct {
	Title     string `form:"title" validate:"required"`
	Year      string `form:"year" validate:"required"`
	Thumbnail string `form:"thumbnail"  validate:"required"`
	Attache   string `form:"attached"`
	ArtistID  int    `form:"artist_id"  validate:"required"`
}
