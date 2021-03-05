package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

type DatabaseConnection struct {
	Server     string
	Database   string
	Collection string
	Port       int
	Username   string
	Password   string
	Timeout    int
}

func (configs *DatabaseConnection) CreateMongoDatabaseUri() string {
	var databaseUri string
	if configs.Username != "" && configs.Password != "" {
		databaseUri = fmt.Sprintf("mongodb://%s:%s@%s:%s", configs.Username, configs.Password, configs.Server, strconv.Itoa(configs.Port))
		log.Println("Created URI DB with auth")
	} else {
		databaseUri = fmt.Sprintf("mongodb://%s:%s", configs.Server, strconv.Itoa(configs.Port))
		log.Println("Created URI DB without auth")
	}
	return databaseUri
}

func (configs *DatabaseConnection) CreateMongoDatabaseSession() *mongo.Client {
	databaseUri := configs.CreateMongoDatabaseUri()
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Println(fmt.Printf("Failed to create client in database: %v ", err.Error()))
	}
	log.Println(fmt.Println("Connected to MongoDB!"))
	return client
}

func (configs *DatabaseConnection) GetCollection(session *mongo.Client) *mongo.Collection {
	log.Println(fmt.Printf("Getting mongo collection"))
	return session.Database(configs.Database).Collection(configs.Collection)
}

func (configs *DatabaseConnection) PingMongoDatabase() bool {
	databaseUri := configs.CreateMongoDatabaseUri()
	contextBg, _ := context.WithTimeout(context.Background(), configs.CreateTimeout())
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Println(fmt.Printf("Failed to create client in database: %v ", err.Error()))
	}
	err = client.Connect(contextBg)
	if err != nil {
		log.Println(fmt.Printf("Failed to connect in database, error: %v ", err.Error()))
	}
	log.Println(fmt.Println("Connected to MongoDB!"))
	defer client.Disconnect(contextBg)

	err = client.Ping(contextBg, nil)
	if err != nil {
		log.Println(fmt.Printf("Failed ping database, error: %v ", err.Error()))
		return false
	}
	return true
}
