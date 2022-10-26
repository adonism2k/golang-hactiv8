package main

import (
	"fmt"
	"log"

	"github.com/adonism2k/golang-hactiv8/api/handlers"
	"github.com/adonism2k/golang-hactiv8/api/routes"
	"github.com/adonism2k/golang-hactiv8/internal/database"
	"github.com/adonism2k/golang-hactiv8/internal/initializers"
)

// @title          Final Project API
// @version        1.0
// @description    Ini adalah dokumentasi API Contract untuk tugas Final Project Hactiv8.
// @termsOfService http://swagger.io/terms/

// @contact.name  Abdian Rizky
// @contact.url   https://linktr.ee/adonism2k
// @contact.email dev.abdianrizky@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8000
// @BasePath /

// @securitydefinitions.apikey
// @name Authorization
// @in header
// @description Bearer Token

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}

	db := database.Connect(config)
	handl := handlers.New(db, config)
	app := routes.Api(handl, config)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.APPUrl, config.ServerPort)))
}
