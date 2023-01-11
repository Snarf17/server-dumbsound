package main

import (
	"dumbsound/database"
	"dumbsound/pkg/mysql"
	"dumbsound/routes"
	
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()
	database.RunMigration()
	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})
	var port = os.Getenv("PORT")
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
