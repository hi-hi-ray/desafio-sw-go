package services

import (
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/config"
	"github.com/hi-hi-ray/desafio-sw-go/api/database"
	"github.com/hi-hi-ray/desafio-sw-go/api/database/repository"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strings"
)

func HealthCheckDatabaseService() string {
	log.Println("Passing by HealthCheckService")
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

	status := dbConnection.PingMongoDatabase()

	if status {
		return "MongoDB it's up."
	}
	return "MongoDB isn't up."
}

func CreatePlanetService(planet *models.Planet) error {
	log.Println(fmt.Printf("[CREATE PLANET SERVICE] Creating Planet %s", planet.Name))

	configs := config.ConfigVariables{}
	config := configs.GetConfigEnvoriments()
	consumer := SwapiConsumer{
		Urlbase:  config.Swapi.Urlbase,
		Endpoint: config.Swapi.Endpoint,
	}

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

	var err error

	planet.ID = primitive.NewObjectID()

	if planet.AmountOfAppearances == 0 {
		amount, err := consumer.GetAmountOfApparitions(planet.Name)
		if amount != 0 {
			planet.AmountOfAppearances = 0
		}
		if err != nil {
			if strings.Contains(err.Error(), errors.SwapiNotFoundError) {
				planet.AmountOfAppearances = 0
			} else {
				return err
			}
		}
		log.Println(fmt.Printf("[CREATE PLANET SERVICE] Updating Appearances of the Planet  %s", planet.Name))
		planet.AmountOfAppearances = amount
	}
	_, err = mongoApp.Insert(planet)
	return err
}

func GetPlanetsService() ([]*models.Planet, error) {
	log.Println("[GET PLANETS SERVICE] Getting all Planets from database")

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

	planets, err := mongoApp.FindAll()
	return planets, err
}

func GetPlanetService(id string) (models.Planet, error) {
	log.Println("[GET PLANETS SERVICE] Getting all Planets from database")

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

	var planet models.Planet
	var err error

	planet, err = mongoApp.FindById(id)
	return planet, err
}

func UpdatePlanetService(id string, planet models.Planet) error {
	log.Println(fmt.Printf("[UPDATE PLANET SERVICE] Creating Planet %s", planet.Name))

	configs := config.ConfigVariables{}
	config := configs.GetConfigEnvoriments()

	consumer := SwapiConsumer{
		Urlbase:  config.Swapi.Urlbase,
		Endpoint: config.Swapi.Endpoint,
	}

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

	if planet.AmountOfAppearances == 0 {
		amount, err := consumer.GetAmountOfApparitions(planet.Name)
		if amount != 0 {
			planet.AmountOfAppearances = 0
		}
		if err != nil {
			if strings.Contains(err.Error(), errors.SwapiNotFoundError) {
				planet.AmountOfAppearances = 0
			} else {
				return err
			}
		}
		log.Println(fmt.Printf("[UPDATE PLANET SERVICE] Updating Appearances of the Planet  %s", planet.Name))
		planet.AmountOfAppearances = amount
	}

	var err error
	_, err = mongoApp.Update(id, &planet)
	return err
}

func DeletePlanetService(id string) error {
	log.Println("[DELETE PLANETS SERVICE] Delete Planet from database")

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

	var err error

	err = mongoApp.DeleteById(id)
	return err
}

func GetPlanetByNameService(name string) (models.Planet, error) {
	log.Println("[GET PLANET BY NAME SERVICE] Getting all Planets from database")

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

	planet, err, _ := mongoApp.FindByName(name)
	return planet, err
}
