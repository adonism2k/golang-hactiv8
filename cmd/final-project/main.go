package main

import (
	"log"

	"github.com/adonism2k/golang-hactiv8/api/handlers"
	"github.com/adonism2k/golang-hactiv8/api/routes"
	"github.com/adonism2k/golang-hactiv8/internal/database"
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

const webPort = ":8000"

func main() {
	db := database.Connect()
	handl := handlers.New(db)
	app := routes.Api(handl)

	log.Fatal(app.Listen(webPort))
}
