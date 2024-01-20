package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"crudapp/config"
)

var db *sql.DB

func main() {
	// Load configuration
	appConfig := config.LoadConfig()

	// Open a connection to the database
	var err error
	db, err = sql.Open("postgres", appConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the database table is created
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		author VARCHAR(255)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router and routes
	router := gin.Default()
	InitializeRoutes(router)

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server is running on :8080")
	log.Fatal(server.ListenAndServe())
}

