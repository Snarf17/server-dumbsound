package artistdto

type ArtistResponse struct {
	ID         int    `form:"id"`
	Name       string `form:"name"`
	Old        int    `form:"old"`
	Type       string `form:"type"`
	StartCarer string `form:"startcarer"`
}
