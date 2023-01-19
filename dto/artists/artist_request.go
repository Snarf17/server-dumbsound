package artistdto

type ArtistRequest struct {
	Name       string `form:"name"  validate:"required"`
	Old        int    `form:"old"  validate:"required"`
	Type       string `form:"type" validate:"required"`
	StartCarer string `form:"startcarer" validate:"required"`
}
