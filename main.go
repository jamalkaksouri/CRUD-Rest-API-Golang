package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-rest-api/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	RegisterProductRoutes(router)

	log.Printf("Starting server on port %s", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), handler))
}
