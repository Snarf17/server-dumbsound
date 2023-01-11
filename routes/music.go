package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func MusicRoute(r *mux.Router) {
	musicRepository := repositories.RepositoryMusic(mysql.DB)
	h := handlers.HandleMusic(musicRepository)

	r.HandleFunc("/musics", h.ShowMusics).Methods("GET")
	r.HandleFunc("/music/{id}", h.GetMusic).Methods("GET")
	r.HandleFunc("/music", middleware.Auth(middleware.UploadFile(middleware.UploadMusic(h.CreateMusic)))).Methods("POST")
}
