package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	AuthRepository := repositories.RepositoryLogReg(mysql.DB)
	h := handlers.HandlerLogReg(AuthRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}
