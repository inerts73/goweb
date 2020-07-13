// 31/51

package main

import (
	"context"
	"log"
	"errors"

	"github.com/inerts73/myGo/goweb/mongotesting/dbiface"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User a dummy user
type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

//actual collection *mongo.Collection
//implements CollectionAPI
//InsertOne will work on actual collection
//fake collection mockCollection
//implements CollectionAPI
//InsertOne fake implementation, when invoked it will work on fake collection

func insertData(collection dbiface.CollectionAPI, user User) (*mongo.InsertOneResult, error) {
	if user.FirstName != "kunal" {
		return nil, errors.New("my bad")	
	}
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))	
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res, err := insertData(col, User{"kunal", "shimpi"})
	log.Println(res, err)
}