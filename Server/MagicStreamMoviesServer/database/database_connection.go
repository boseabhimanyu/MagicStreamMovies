package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() *mongo.Client {
	//load file
	err := godotenv.Load(".env")
	//checking .env file
	if err != nil {
		log.Println("Warning: unable to find the env file")
	}
	//checking env file
	MongoDb := os.Getenv("MONGODB_URI")
	//checking for empty string
	if MongoDb == "" {
		log.Fatal("MONGODB_URI not set")
	}
	fmt.Println("MONGODB URI: ", MongoDb)
	//connecting to the mongodb
	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil
	}

	return client
}

func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {

	//load .env file
	err := godotenv.Load(".env")
	//error checking
	if err != nil {
		log.Println("Warning: unable to find the env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE_NAME", databaseName)
	// database uri and database
	collection := client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}

	return collection
}
