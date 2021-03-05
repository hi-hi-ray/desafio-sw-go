package unit

import (
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"github.com/hi-hi-ray/desafio-sw-go/api/services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestPlanetModelValidatorSuccess(t *testing.T) {
	var MockPlanet = models.Planet{
		ID:                  primitive.ObjectID{},
		Name:                "Teste",
		Climate:             "Teste",
		Terrain:             "Teste",
		AmountOfAppearances: 6,
	}
	err, isValid := services.PlanetModelValidator(&MockPlanet)
	assert.Equal(t, true, isValid)
	assert.Nil(t, err, nil)
}

func TestPlanetModelValidatorNameFail(t *testing.T) {
	var MockPlanet = models.Planet{
		ID:                  primitive.ObjectID{},
		Name:                "",
		Climate:             "Teste",
		Terrain:             "Teste",
		AmountOfAppearances: 6,
	}
	err, isValid := services.PlanetModelValidator(&MockPlanet)
	assert.Equal(t, false, isValid)
	assert.Error(t, err, errors.PlanetMandatoryFields)
}

func TestPlanetModelValidatorClimateFail(t *testing.T) {
	var MockPlanet = models.Planet{
		ID:                  primitive.ObjectID{},
		Name:                "Teste",
		Climate:             "",
		Terrain:             "Teste",
		AmountOfAppearances: 6,
	}
	err, isValid := services.PlanetModelValidator(&MockPlanet)
	assert.Equal(t, false, isValid)
	assert.Error(t, err, errors.PlanetMandatoryFields)
}

func TestPlanetModelValidatorTerrainFail(t *testing.T) {
	var MockPlanet = models.Planet{
		ID:                  primitive.ObjectID{},
		Name:                "Teste",
		Climate:             "Teste",
		Terrain:             "",
		AmountOfAppearances: 6,
	}
	err, isValid := services.PlanetModelValidator(&MockPlanet)
	assert.Equal(t, false, isValid)
	assert.Error(t, err, errors.PlanetMandatoryFields)
}

func TestPlanetModelValidatorAllEmptyFail(t *testing.T) {
	var MockPlanet = models.Planet{
		ID:                  primitive.ObjectID{},
		Name:                "",
		Climate:             "",
		Terrain:             "",
		AmountOfAppearances: 6,
	}
	err, isValid := services.PlanetModelValidator(&MockPlanet)
	assert.Equal(t, false, isValid)
	assert.Error(t, err, errors.PlanetMandatoryFields)
}