package dtomusic

type MusicResponse struct {
	Title     string         `form:"title"`
	Year      int            `form:"year"`
	Thumbnail string         `form:"thumbnail"`
	Attache   string         `form:"attached"`
	ArtistID  int            `form:"artist_id"`
	Artist    ArtistResponse `json:"artist" `
}
type ArtistResponse struct {
	ID         int    `form:"id"`
	Name       string `form:"name"`
	Old        int    `form:"old"`
	Type       string `form:"type"`
	StartCarer string `form:"startcarer"`
}
