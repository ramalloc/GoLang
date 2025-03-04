package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/ramalloc/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://ramalloc:ramalloc@cluster0.ocvjngn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName string = "stream"
const collectionName string = "watchlist"

var collection *mongo.Collection

// Connecting with mongodb in init function which is special function in go runs firstly and only once.
func init() {
	// client option to get clients (mongo, sql etc.) present in mod
	clientOption := options.Client().ApplyURI(connectionString)

	// connect ot mongodb
	// Context is responsible for handling the connection from other machine, When we do not know which context to use then we use
	// TODO()
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully....")

	collection = client.Database(dbName).Collection(collectionName)
	// Now the above collection got or pointing to the Database and collection from mongoDB

	fmt.Println("Collection instance created...")
}

// MONGO DB Helpers

// --. Insert 1 movie in db
func insertOneMovie(movie models.Netflix) interface{} {
	inserted, insertErr := collection.InsertOne(context.Background(), movie)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println("Inserted one movie with id :- ", inserted.InsertedID)
	return inserted.InsertedID
}

// --> Update a movie

func updateMovie(movieId string) {
	// Converting heax id into ObjectID
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	// fmt.Printf("Filter: %+v\n", filter)
	update := bson.M{"$set": bson.M{"watched": true}}
	// fmt.Printf("Update: %+v\n", update)
	movieWatched, watchingErr := collection.UpdateOne(context.Background(), filter, update)
	if watchingErr != nil {
		log.Fatal(watchingErr)
	}
	fmt.Println("Movie Watched :- ", movieWatched.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletedMovie, deletionErr := collection.DeleteOne(context.Background(), filter)
	if deletionErr != nil {
		log.Fatal(deletionErr)
	}
	fmt.Println("Movie Deleted with count :- ", deletedMovie.DeletedCount)
}

func deleteAllMovie() {
	deletedAllMovie, deletionErr := collection.DeleteMany(context.Background(), bson.D{{}})
	if deletionErr != nil {
		log.Fatal(deletionErr)
	}
	fmt.Println("Movies Deleted with count :- ", deletedAllMovie.DeletedCount)
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		decodeErr := cursor.Decode(&movie)
		if decodeErr != nil {
			log.Fatal(decodeErr)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Actual Controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("All-Control-Alloc-Methods", "POST")

	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	id := insertOneMovie(movie)

	// Perform type assertion on the inserted ID
	objectID, ok := id.(primitive.ObjectID)
	if !ok {
		log.Println("Error converting inserted ID to primitive.ObjectID")
		http.Error(w, "Failed to retrieve movie ID", http.StatusInternalServerError)
		return
	}

	// Assign the ObjectID to the movie ID
	movie.Id = objectID
	response := map[string]interface{}{
		"message": "Movie Added Successfully",
		"movie":   movie,
	}
	json.NewEncoder(w).Encode(response)
}

func MarkMovieWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("All-Control-Alloc-Methods", "PUT")

	params := mux.Vars(r)
	updateMovie(params["id"])
	response := map[string]interface{}{
		"message":  "Movie Updated Successfully",
		"movie id": params["id"],
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("All-Control-Alloc-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	response := map[string]interface{}{
		"message":  "Movie Deleted Successfully",
		"movie id": params["id"],
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("All-Control-Alloc-Methods", "DELETE")

	deleteAllMovie()
	response := map[string]interface{}{
		"message": "All Movies Deleted Successfully",
	}
	json.NewEncoder(w).Encode(response)
}
