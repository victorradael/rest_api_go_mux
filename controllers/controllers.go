package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorradael/rest_api_go_mux/database"
	"github.com/victorradael/rest_api_go_mux/models"
)

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalitie
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetOnePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalitie
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func DeleteOnePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalitie
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

// Esse put funciona perfeitamente como um patch também
func UpdateOnePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalitie
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func CreateNewPersonalitie(w http.ResponseWriter, r *http.Request) {
	var p models.Personalitie
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}
