package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
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
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", WriteHTML)

	r.Run(":8000")
}

func WriteHTML(c *gin.Context) {
	WriteJSON()
	file, err := ioutil.ReadFile("status.json")
	if err != nil {
		panic(err)
	}

	var data Status
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	var wind_status, water_status string

	switch {
	case data.Status.Water < 5:
		water_status = "Aman" + strconv.Itoa(data.Status.Water)
	case data.Status.Water > 6 && data.Status.Water < 8:
		water_status = "Siaga" + strconv.Itoa(data.Status.Water)
	case data.Status.Water > 8:
		water_status = "Bahaya" + strconv.Itoa(data.Status.Water)
	}

	switch {
	case data.Status.Wind < 6:
		wind_status = "Aman" + strconv.Itoa(data.Status.Wind)
	case data.Status.Wind > 7 && data.Status.Wind < 15:
		wind_status = "Siaga" + strconv.Itoa(data.Status.Wind)
	case data.Status.Wind > 15:
		wind_status = "Bahaya" + strconv.Itoa(data.Status.Wind)
	}

	c.HTML(200, "index.tmpl", gin.H{
		"wind":  wind_status,
		"water": water_status,
	})
}

func WriteJSON() {
	rand.Seed(time.Now().UTC().UnixNano())
	wind := rand.Intn(100)
	water := rand.Intn(100)

	log.Println("Wind: ", wind)
	log.Println("Water: ", water)

	data := Status{
		Data{
			Wind:  wind,
			Water: water,
		},
	}

	jsonString, _ := json.MarshalIndent(data, "", "    ")
	ioutil.WriteFile("status.json", jsonString, os.ModePerm)
}
