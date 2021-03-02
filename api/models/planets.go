package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID                  primitive.ObjectID `bson:"_id" json:"id"`
	Name                string             `bson:"name" json:"name"`
	Climate             string             `bson:"climate" json:"climate"`
	Terrain             string             `bson:"terrain" json:"terrain"`
	AmountOfAppearances int                `bson:"amountOfAppearances" json:"amount_of_appearances"`
}