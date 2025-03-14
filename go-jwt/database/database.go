// DBInstance returns a MongoDB client instance connected to the MongoDB URI specified in the .env file.
// It uses the context.WithTimeout function to set a timeout of 100 seconds for the connection attempt.
// If the connection is successful, it prints "Connected To Database..." to the console.
// func DBInstance() *mongo.Client {
	// ...
// }

// OpenCollection returns a MongoDB collection for the specified collection name.
// It uses the DB_NAME environment variable to determine the database name.
// func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// ...
// }
package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("Database Connection Failed with error : ", err)
	}

	fmt.Println("Connected To Database...")

	return client
}

// MongoDB Client Instance
var Client *mongo.Client = DBInstance()


// Returning Particular Collection, To Get Particular collection of a DB
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DbName := os.Getenv("DB_NAME")
	var collection *mongo.Collection = client.Database(DbName).Collection(collectionName)
	return collection
}
