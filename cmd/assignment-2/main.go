package main

import (
	"assignment-2/api/handlers"
	"assignment-2/api/routes"
	"assignment-2/internal/database"
	"assignment-2/internal/models"
	"log"
	"net/http"
	"time"
)

const webPort = ":8000"

func main() {
	log.Println("Starting the application...")

	// connect to postgres
	conn := database.Connect()

	// Start the server
	StartServer(conn)
}

func StartServer(conn database.Config) {
	h := handlers.New(conn, models.New(conn.DB))

	s := &http.Server{
		Addr:         webPort,
		Handler:      routes.Api(h),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server is listening on 127.0.0.1%s ", webPort)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
