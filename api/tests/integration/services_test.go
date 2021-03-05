package integration

import (
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"github.com/hi-hi-ray/desafio-sw-go/api/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapiIntegrationGetAmountOfAppearencesSuccess(t *testing.T) {
	consumer := services.SwapiConsumer{
		Urlbase:  "https://swapi.dev/",
		Endpoint: "api/planets/",
	}

	result, err := consumer.GetAmountOfApparitions("Tatooine")
	assert.Nil(t, err)
	assert.Equal(t, result, 5)
}

func TestSwapiIntegrationGetAmountOfAppearencesNotFound(t *testing.T) {
	consumer := services.SwapiConsumer{
		Urlbase:  "https://swapi.dev/",
		Endpoint: "api/planets/",
	}

	result, err := consumer.GetAmountOfApparitions("Teste")
	assert.Error(t, err, errors.SwapiNotFoundError)
	assert.Equal(t, result, 0)
}

func TestSwapiIntegrationGetPlanetSuccess(t *testing.T) {
	consumer := services.SwapiConsumer{
		Urlbase:  "https://swapi.dev/",
		Endpoint: "api/planets/",
	}
	expectedResponse := models.PlanetSwapi{Count: 1, Next: "", Previous: "", Results: []models.Results{Result}}

	result, err := consumer.GetPlanets("Tatooine")
	assert.Nil(t, err)
	assert.Equal(t, result, expectedResponse)
}

func TestSwapiIntegrationGetPlanetNotFound(t *testing.T) {
	consumer := services.SwapiConsumer{
		Urlbase:  "https://swapi.dev/",
		Endpoint: "api/planets/",
	}
	expectedResponse := models.PlanetSwapi{Count: 0, Next: "", Previous: "", Results: []models.Results{}}

	result, err := consumer.GetPlanets("Teste")
	assert.Error(t, err, errors.SwapiNotFoundError)
	assert.Equal(t, result, expectedResponse)
}
