package repository

import (
	"context"
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/database"
	"github.com/hi-hi-ray/desafio-sw-go/api/errors"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoApp struct {
	Client          *mongo.Client
	Collection      *mongo.Collection
	DatabaseConfigs *database.DatabaseConnection
}

func (db MongoApp) Insert(planet *models.Planet) (*mongo.InsertOneResult, error) {
	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	log.Println("[REPOSITORY INSERT] Start the process to insert in database.")
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY INSERT] Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	log.Println("[REPOSITORY INSERT] Connection made with success.")

	planetResult, errInsert := db.Collection.InsertOne(contextBg, planet)
	if errInsert != nil {
		log.Println(fmt.Printf("[REPOSITORY INSERT] Failed to insert, error: %v ", errInsert.Error()))
		return nil, errInsert
	}
	log.Println(fmt.Printf("[REPOSITORY INSERT] Insert success: %v ", planetResult))

	return planetResult, nil
}

func (db MongoApp) InsertMany(planets []interface{}) (*mongo.InsertManyResult, error) {
	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	log.Println("[REPOSITORY INSERTMANY] Start the process to insert many in database.")
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY INSERTMANY] Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	log.Println("[REPOSITORY INSERTMANY] Connection made with success.")

	planetResult, errInsert := db.Collection.InsertMany(contextBg, planets)
	if errInsert != nil {
		log.Println(fmt.Printf("[REPOSITORY INSERTMANY] Failed to insert, error: %v ", errInsert.Error()))
		return nil, errInsert
	}
	log.Println(fmt.Printf("[REPOSITORY INSERTMANY] Insert many success: %v ", planetResult))

	return planetResult, nil
}

func (db MongoApp) FindAll() ([]*models.Planet, error) {
	var planets []*models.Planet

	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	log.Println("[REPOSITORY FIND ALL] Start the process to find all data in database.")
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY FIND ALL] Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	log.Println("[REPOSITORY FIND ALL] Connection made with success.")

	cursor, err := db.Collection.Find(contextBg, bson.D{{}})
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY FIND ALL] Failed to find all data, error: %v ", err.Error()))
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var elem models.Planet
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println(fmt.Printf("[REPOSITORY FIND ALL] Failed to decode data, error: %v ", err.Error()))
			return nil, err
		}
		planets = append(planets, &elem)
	}
	if err := cursor.Err(); err != nil {
		log.Println(fmt.Printf("[REPOSITORY FIND ALL] Failed in cursor data, error: %v ", err.Error()))
		return nil, err
	}
	log.Println(fmt.Printf("[REPOSITORY FIND ALL] Founded all with success: %v ", planets))

	return planets, nil
}

func (db MongoApp) FindById(id string) (models.Planet, error) {
	idMongo, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": idMongo}
	var planet models.Planet
	log.Println("[REPOSITORY FIND BY ID] Start the process to find by id in database.")

	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY FIND BY ID] Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	log.Println("[REPOSITORY FIND BY ID] Connection made with success.")

	err = db.Collection.FindOne(contextBg, filter).Decode(&planet)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY FIND BY ID] Failed to find planet, error: %v ", err.Error()))
		return planet, err
	}
	log.Println(fmt.Printf("[REPOSITORY FIND BY ID] Found planet with success: %v ", planet))

	return planet, nil
}

func (db MongoApp) Update(id string, planet *models.Planet) (*mongo.UpdateResult, error) {
	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	err := db.Client.Connect(contextBg)
	log.Println("[REPOSITORY UPDATE] Start the process to update planet.")
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	log.Println("[REPOSITORY UPDATE] Connection made with success.")

	idMongo, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", idMongo}}
	var update bson.D

	update = bson.D{
		{"$set", bson.D{
			{"_id", idMongo},
			{"name", planet.Name},
			{"climate", planet.Climate},
			{"terrain", planet.Terrain},
			{"amountOfAppearances", planet.AmountOfAppearances},
		}},
	}

	planetResult, err := db.Collection.UpdateOne(contextBg, filter, update)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY UPDATE] Failed to update planet, error: %v ", err.Error()))
		return nil, err
	}
	log.Println(fmt.Printf("[REPOSITORY UPDATE] Planet updated with success: %v ", planetResult))

	return planetResult, nil
}

func (db MongoApp) DeleteById(id string) error {
	idMongo, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": idMongo}
	log.Println("[REPOSITORY DELETE BY ID] Start the process to delete by id in database.")

	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY DELETE BY ID] Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	_, err = db.FindById(id)
	if err.Error() != "mongo: no documents in result" {
		planetResult, err := db.Collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Println(fmt.Printf("[REPOSITORY DELETE BY ID] Failed to delete planet, error: %v ", err.Error()))
			return err
		}
		log.Println(fmt.Printf("[REPOSITORY DELETE BY ID] Planet deleted with success: %v ", planetResult))

		return nil
	}
	log.Println("[REPOSITORY DELETE BY ID] Planet was not found.")
	err = errors.Create(errors.PlanetDoesNotExist)
	return err
}

func (db MongoApp) FindByName(name string) (models.Planet, error, bool) {
	filter := bson.M{"name": name}
	var planet models.Planet
	log.Println("[REPOSITORY FINDBYNAME] Start the process to find planet by name in database.")

	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY FINDBYNAME] Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)
	log.Println("[REPOSITORY FINDBYNAME] Connection made with success.")

	err = db.Collection.FindOne(context.TODO(), filter).Decode(&planet)
	if err != nil {
		log.Println(fmt.Printf("[REPOSITORY FINDBYNAME] Planet not founded with success, error: %v ", planet))
		return planet, err, false
	}
	log.Println(fmt.Printf("[REPOSITORY FINDBYNAME] Planet founded with success: %v ", planet))

	return planet, nil, true
}
