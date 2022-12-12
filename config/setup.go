package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var driver = DotEnvVar("DRIVER_PATH")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(driver))

	if err != nil {
		log.Fatal("database connection failed : ", err)
	}

	println("database successfully connected...")

	return client
}

// //get db instance
// var DB *mongo.Client = ConnectDB()

func GetCollection(collectionName string) *mongo.Collection {
	collection := ConnectDB().Database("Attendance").Collection(collectionName)

	return collection
}

func DotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
