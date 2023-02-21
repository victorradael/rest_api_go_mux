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
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetOnePersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreateNewPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeleteOnePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdateOnePersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
