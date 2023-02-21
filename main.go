package main

import (
	"fmt"

	"github.com/victorradael/rest_api_go_mux/database"
	"github.com/victorradael/rest_api_go_mux/models"
	"github.com/victorradael/rest_api_go_mux/routes"
)

func main() {
	models.Personalities = []models.Personalitie{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}
	database.ConnectWithDatabase()
	fmt.Println("Start Server")
	routes.HandleRequest()

}
