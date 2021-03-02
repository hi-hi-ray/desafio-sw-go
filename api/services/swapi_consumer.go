package services

import (
	"encoding/json"
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"io/ioutil"
	"log"
	"net/http"
)

type SwapiConsumer struct {
	Urlbase  string
	Endpoint string
}

func (consumer SwapiConsumer) GetPlanets(planet string) (models.PlanetSwapi, error) {
	var swPlanet models.PlanetSwapi
	url := fmt.Sprintf("%s%s?search=%s", consumer.Urlbase, consumer.Endpoint, planet)
	log.Println(fmt.Printf("[SWAPI CONSUMER] Getting infos from swapi using the url: %s", url))

	response, err := http.Get(url)
	if err != nil {
		log.Println(fmt.Sprint(errors.SwapiConsumerError, err.Error()))
		return swPlanet, err
	}

	if response.StatusCode == http.StatusOK {
		log.Println(fmt.Printf("[SWAPI CONSUMER] Status request response: %v", response.StatusCode))

		bodyBytes, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			log.Println(fmt.Sprint(errors.ReadBodyError, err.Error()))
			return swPlanet, err
		}
		log.Println(fmt.Printf("[SWAPI CONSUMER] Body request response : %v", string(bodyBytes)))

		err = json.Unmarshal(bodyBytes, &swPlanet)
		if err != nil {
			log.Println(fmt.Sprint(errors.MarshallBodyError, err.Error()))
			return swPlanet, err
		}
		log.Println(fmt.Printf("[SWAPI CONSUMER] Success getting planet: %s", swPlanet.Results[0].Name))
		return swPlanet, nil
	}
	if response.StatusCode == http.StatusNotFound {
		return swPlanet, errors.Create(errors.SwapiNotFoundError)
	}
	return swPlanet, errors.Create(errors.SwapiInternalError)
}

func (consumer SwapiConsumer) GetAmountOfApparitions(planet string) (int, error) {
	swPlanet := models.PlanetSwapi{}
	var err error
	swPlanet, err = consumer.GetPlanets(planet)
	if err != nil{
		return 0, err
	}

	quantityOfMovies := len(swPlanet.Results[0].FilmURLs)

	if quantityOfMovies != 0 {
		return quantityOfMovies, nil
	}
	return quantityOfMovies, nil
}
