package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const driver = "mongodb+srv://roshannahak:112233Raja@cluster0.sq55i.mongodb.net/?retryWrites=true&w=majority"

func ConnectDB() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(driver))

	if err != nil {
		log.Fatal("database connection failed : ", err)
	}

	println("database successfully connected...")

	return client;
}

// //get db instance
// var DB *mongo.Client = ConnectDB()

func GetCollection(collectionName string) *mongo.Collection {
	collection := ConnectDB().Database("Attendance").Collection(collectionName)

	return collection;
}

