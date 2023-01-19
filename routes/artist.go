package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func ArtistRoute(r *mux.Router) {
	ArtistRepository := repositories.RepositoryArtist(mysql.DB)
	h := handlers.HandleArtist(ArtistRepository)

	r.HandleFunc("/artis", h.ShowArtists).Methods("GET")
	r.HandleFunc("/artis", middleware.Auth(h.AddArtist)).Methods("POST")
	r.HandleFunc("/artis/{id}", middleware.Auth(h.GetArtist)).Methods("GET")

}
