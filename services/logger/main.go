package main

import (
	"controller/internal/model"
	"controller/internal/server"
	"controller/internal/util"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	if os.Getenv("DEPLOYMENT") != "PROD" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)

	// Create db conn
	err := model.InitDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create the router
	router := mux.NewRouter()

	// Use logging middleware
	router.Use(util.LoggingMiddleware)

	// Register the routes
	server.RegisterRoutes(router)
	port := os.Getenv("APP_PORT")
	fmt.Println(fmt.Sprintf("Started server at: :%s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

}
