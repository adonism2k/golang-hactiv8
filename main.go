package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Status Data `json:"status"`
}

type Data struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

var count int

func main() {

	WriteJSON()

	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", WriteHTML)

	r.Run(":8000")
}

func WriteHTML(c *gin.Context) {
	data := gin.H{
		"wind_status":  "Bahaya",
		"water_status": "Bahaya",
	}

	c.HTML(200, "index.html", data)
}

func WriteJSON() {
	rand.Seed(time.Now().UTC().UnixNano())
	wind := rand.Intn(100)
	water := rand.Intn(100)

	data := Status{
		Data{
			Wind:  wind,
			Water: water,
		},
	}

	jsonString, _ := json.MarshalIndent(data, "", "    ")
	ioutil.WriteFile("status.json", jsonString, os.ModePerm)
}
