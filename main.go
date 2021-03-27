package main

import (
	"context"
	"encoding/json"

	"github.com/enriquelira1994/go-rest-api/helper"
	"github.com/enriquelira1994/go-rest-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Data array
	var datos []models.Data

	// bson.M{},  we passed empty filter. So we want to get all datos.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var dato models.Data
		// & character returns the memory address of the following variable.
		err := cur.Decode(&dato) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		datos = append(datos, dato)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(datos) // encode similar to serialize process.
}

func createData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Data

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}


// var client *mongo.Client

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/data", getData).Methods("GET")
	r.HandleFunc("/api/data", createData).Methods("POST")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
