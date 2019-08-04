package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_URI string = "mongodb+srv://mongouser:mongouser123@mongo-test-instance-musz7.mongodb.net/test?retryWrites=true&w=majority"
const DB_NAME = "test"

var DB *mongo.Database

func InitDB() { // Set client options

	clientOptions := options.Client().ApplyURI(DB_URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(DB_NAME)
}
