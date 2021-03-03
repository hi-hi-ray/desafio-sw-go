package repository

import (
	"context"
	"fmt"
	"github.com/hi-hi-ray/desafio-sw-go/api/database"
	"github.com/hi-hi-ray/desafio-sw-go/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const defaultTimeout = 30 * time.Second

type MongoApp struct {
	Client     *mongo.Client
	Collection *mongo.Collection
	DatabaseConfigs *database.DatabaseConnection
}

func (db MongoApp) Insert(planet *models.Planet) (*mongo.InsertOneResult, error) {
	contextBg, _ := context.WithTimeout(context.Background(), db.DatabaseConfigs.CreateTimeout())

	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	planetResult, errInsert := db.Collection.InsertOne(contextBg, planet)
	if errInsert != nil {
		log.Println(fmt.Printf("Failed to insert, error: %v ", errInsert.Error()))
		return nil, errInsert
	}
	log.Println(fmt.Printf("Insert success: %v ", planetResult))

	return planetResult, nil
}

func (db MongoApp) InsertMany(planets []interface{}) (*mongo.InsertManyResult, error) {
	contextBg, _ := context.WithTimeout(context.Background(), defaultTimeout)
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	planetResult, errInsert := db.Collection.InsertMany(contextBg, planets)
	if errInsert != nil {
		log.Println(fmt.Printf("Failed to insert, error: %v ", errInsert.Error()))
		return nil, errInsert
	}
	log.Println(fmt.Printf("Insert many success: %v ", planetResult))

	return planetResult, nil
}

func (db MongoApp) FindAll() ([]*models.Planet, error) {
	var planets []*models.Planet

	contextBg, _ := context.WithTimeout(context.Background(), defaultTimeout)
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	cursor, err := db.Collection.Find(contextBg, bson.D{{}})
	if err != nil {
		log.Println(fmt.Printf("Failed to find all data, error: %v ", err.Error()))
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var elem models.Planet
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println(fmt.Printf("Failed to decode data, error: %v ", err.Error()))
			return nil, err
		}
		planets = append(planets, &elem)
	}
	if err := cursor.Err(); err != nil {
		log.Println(fmt.Printf("Failed in cursor data, error: %v ", err.Error()))
		return nil, err
	}
	log.Println(fmt.Printf("Founded all with success: %v ", planets))

	return planets, nil
}

func (db MongoApp) FindById(id string) (models.Planet, error) {
	idMongo, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": idMongo}
	var planet models.Planet

	contextBg, _ := context.WithTimeout(context.Background(), defaultTimeout)
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	err = db.Collection.FindOne(contextBg, filter).Decode(&planet)
	if err != nil {
		log.Println(fmt.Printf("Failed to find planet, error: %v ", err.Error()))
		return planet, err
	}
	log.Println(fmt.Printf("Found planet with success: %v ", planet))

	return planet, nil
}

func (db MongoApp) Update(id string, planet bson.D) (*mongo.UpdateResult, error) {
	idMongo, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", idMongo}}

	contextBg, _ := context.WithTimeout(context.Background(), defaultTimeout)
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	planetResult, err := db.Collection.UpdateOne(contextBg, filter, planet)
	if err != nil {
		log.Println(fmt.Printf("Failed to update planet, error: %v ", err.Error()))
		return nil, err
	}
	log.Println(fmt.Printf("Planet updated with success: %v ", planetResult))

	return planetResult, nil
}

func (db MongoApp) DeleteById(id string) (*mongo.DeleteResult, error) {
	idMongo, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": idMongo}

	contextBg, _ := context.WithTimeout(context.Background(), defaultTimeout)
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	planetResult, err := db.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(fmt.Printf("Failed to delete planet, error: %v ", err.Error()))
		return nil, err
	}
	log.Println(fmt.Printf("Planet deleted with success: %v ", planetResult))

	return planetResult, nil
}

func  (db MongoApp) FindByName(name string) (models.Planet, error) {
	filter := bson.M{"name": name}
	var planet models.Planet

	contextBg, _ := context.WithTimeout(context.Background(), defaultTimeout)
	err := db.Client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	defer db.Client.Disconnect(contextBg)

	err = db.Collection.FindOne(context.TODO(), filter).Decode(&planet)
	if err != nil {
		log.Println(fmt.Printf("Planet not founded with success, error: %v ", planet))
		return planet, err
	}
	log.Println(fmt.Printf("Planet founded with success: %v ", planet))

	return planet, nil
}
