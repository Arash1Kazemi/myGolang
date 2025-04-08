package controller

import (
	"DB/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionStrinf = "mongodb+srv://arashkazemi7l5o:<db_password>@cluster0.nm1bsu6.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collectionName = "movies"

var collection *mongo.Collection

// connect with mongo\

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionStrinf)

	//connect to mangodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo db connection success")

	//connect to collection
	collection = client.Database(dbName).Collection(collectionName)

	//collection instance
	fmt.Println("collection instance created is ready")

}

//MongoDB Helper - file

//insert one movie to mongodb

func insertOneMovie(movie model.Netflix) {

	//insert the movie in the collection
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	//print the inserted ID of the movie
	fmt.Println("Inserted a single movie in MongoDB with ID:", insertResult.InsertedID)
	return
}

// Update one movie in mongodb

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // convert string to objectID
	//create a filter to find the movie by id
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched ": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

// Dleate one movie in mongodb

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // convert string to objectID
	//create a filter to find the movie by id
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got dleated:", result.DeletedCount)
}

//dleate all movies in mongodb

func deleteAllMovies() int64 {
	//filter := bson.D{{}} // bson.m{{}}
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) // like siniors
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of movies got dleated:", result.DeletedCount)
	return result.DeletedCount
}

// get all movies from mongodb

func getAllMovies() []primitive.M {
	//find all movies in the collection
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return movies
}

// actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// create a new movie

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

//Mark a movie as watched

func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("movie marked as watched")
}

//Delete a movie

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("movie deleted")
}

//Delete all movies

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deletedCount := deleteAllMovies()
	json.NewEncoder(w).Encode(fmt.Sprintf("%d movies deleted", deletedCount))
}
