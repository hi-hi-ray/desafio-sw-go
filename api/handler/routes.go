package handler

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/hi-hi-ray/desafio-sw-go/api/config"
	"log"
	"strconv"
)

func HandleRequests() {
	configs := config.ConfigVariables{}
	config := configs.GetConfigEnvoriments()

	router := martini.Classic()

	router.Get("/healthcheck", HealthCheck)
	router.Post("/planets", CreatePlanet)
	router.Get("/planets", GetPlanets)
	router.Get("/planets/:id", GetPlanet)
	router.Put("/planets/:id", UpdatePlanet)
	router.Delete("/planets/:id", DeletePlanet)
	router.Get("/planets/search/:name", GetPlanetByName)

	router.RunOnAddr(fmt.Sprintf(":%s", strconv.Itoa(config.Servers.Port)))
	log.Println("Server it's up")
}
