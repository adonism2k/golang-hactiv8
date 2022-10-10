package main

// @title          Assignment 2 API
// @version        1.0
// @description    Ini adalah dokumentasi API Contract untuk tugas assignment 2 Hactiv8.
// @termsOfService http://swagger.io/terms/

// @contact.name  Abdian Rizky
// @contact.url   https://linktr.ee/adonism2k
// @contact.email dev.abdianrizky@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8000
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

import (
	"assignment-2/api/handlers"
	"assignment-2/api/routes"

	// docs "assignment-2/docs"
	"assignment-2/internal/database"
	"assignment-2/internal/model"
	"log"
	"net/http"
	"time"
	// "github.com/gin-gonic/gin"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

const webPort = ":8000"

func main() {
	log.Println("Starting the application...")

	conn := database.Connect()

	StartServer(conn)
	// r := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// r.Run(":8000")
}

func StartServer(conn database.Config) {
	h := handlers.Config{
		DB:     conn,
		Models: model.New(conn.DB),
	}

	s := &http.Server{
		Addr:         webPort,
		Handler:      routes.Api(h),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server is listening on http://localhost%s/ ", webPort)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
