package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/victorradael/rest_api_go_mux/controllers"
	"github.com/victorradael/rest_api_go_mux/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetOnePersonalitie).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreateNewPersonalitie).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteOnePersonalitie).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdateOnePersonalitie).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
