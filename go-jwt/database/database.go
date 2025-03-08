package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading .env !")
	}

	MongoDbUri := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MongoDbUri)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Database Connection Failed with error : ", err)
	}

	fmt.Println("Connected To Database...")

	return client
}

var Client *mongo.Client = DBInstance()


// Returning Particular Collection, To Get Particular collection of a DB

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DbName := os.Getenv("DB_NAME")
	var collection *mongo.Collection = client.Database(DbName).Collection(collectionName)
	return collection
}
