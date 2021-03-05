package services

import (
	"github.com/hi-hi-ray/desafio-sw-go/api/config"
	"github.com/hi-hi-ray/desafio-sw-go/api/database"
	"github.com/hi-hi-ray/desafio-sw-go/api/database/repository"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"log"
)

func PlanetModelValidator(planet *models.Planet) (error, bool) {
	log.Println("[PLANET MODEL VALIDATOR] Checking body of the requisition.")
	if (planet.Name == "") || (planet.Terrain == "") || (planet.Climate == "") {
		err := errors.Create(errors.PlanetMandatoryFields)
		return err, false
	}
	return nil, true
}

func PlanetExistsValidator(planet *models.Planet) bool {
	log.Println("[PLANET EXISTS VALIDATOR] Checking if the already planet exists.")
	configs := config.ConfigVariables{}
	config := configs.GetConfigEnvoriments()

	dbConnection := database.DatabaseConnection{
		Server:     config.Database.Server,
		Database:   config.Database.Database,
		Collection: config.Database.Collection,
		Port:       config.Database.Port,
		Username:   config.Database.Username,
		Password:   config.Database.Password,
		Timeout:    config.Database.Timeout,
	}

	dbSession := dbConnection.CreateMongoDatabaseSession()
	mongoApp := repository.MongoApp{
		Client:          dbSession,
		Collection:      dbConnection.GetCollection(dbSession),
		DatabaseConfigs: &dbConnection,
	}
	_, _, planetExits := mongoApp.FindByName(planet.Name)
	return planetExits
}
