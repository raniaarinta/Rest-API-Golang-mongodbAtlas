package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Fullname string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
}
type Destination struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Spotname        string             `json:"spotname,omitempty" bson:"spotname,omitempty"`
	Popularity      int                `json:"popularity,omitempty" bson:"popularity,omitempty"`
	Characteristics string             `json:"characteristics,omitempty" bson:"characteristics,omitempty"`
	Category        string             `json:"category ,omitempty" bson:"category ,omitempty"`
	Cost            float64            `json:"cost,omitempty" bson:"cost,omitempty"`
	Location        string             `json:"location,omitempty" bson:"location,omitempty"`
}
type TravelProvider struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Orgname     string             `json:"orgname,omitempty" bson:"orgname,omitempty"`
	OrgWebsite  string             `json:"orgwebsite,omitempty" bson:"orgwebsite,omitempty"`
	Phonenumber string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Details     string             `json:"details,omitempty" bson:"details,omitempty"`
}
type EventAttractions struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Eventname      string             `json:"eventname,omitempty" bson:"eventname,omitempty"`
	EventStartDate string             `json:"EventStartDate,omitempty" bson:"EventStartDate,omitempty"`
	EventEndDate   string             `json:"EventEndDate,omitempty" bson:"EventEndDate,omitempty"`
	EventDesc      string             `json:"EventDesc,omitempty" bson:"EventDesc,omitempty"`
	EventPrice     float64            `json:"EventPrice,omitempty" bson:"EventPrice,omitempty"`
	EventLocation  string             `json:"EventLocation,omitempty" bson:"EventLocation,omitempty"`
}

//user
func createUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myUser User
	_ = json.NewDecoder(request.Body).Decode(&myUser)
	collection := client.Database("tourism").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, myUser)
	json.NewEncoder(response).Encode(result)
}
func getUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var myUser User
	collection := client.Database("tourism").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&myUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myUser)
}
func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myUser []User
	collection := client.Database("tourism").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var U User
		cursor.Decode(&U)
		myUser = append(myUser, U)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myUser)
}

//destination
func createDestination(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myDestination Destination
	_ = json.NewDecoder(request.Body).Decode(&myDestination)
	collection := client.Database("tourism").Collection("destinations")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, myDestination)
	json.NewEncoder(response).Encode(result)
}
func getDestinations(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myDestination []Destination
	collection := client.Database("tourism").Collection("destinations")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var D Destination
		cursor.Decode(&D)
		myDestination = append(myDestination, D)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myDestination)

}
func getDestination(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var myDestination Destination
	collection := client.Database("tourism").Collection("destinations")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&myDestination)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myDestination)
}

//event
func createEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myEvent EventAttractions
	_ = json.NewDecoder(request.Body).Decode(&myEvent)
	collection := client.Database("tourism").Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, myEvent)
	json.NewEncoder(response).Encode(result)
}
func getEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myEvent []EventAttractions
	collection := client.Database("tourism").Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var E EventAttractions
		cursor.Decode(&E)
		myEvent = append(myEvent, E)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myEvent)

}
func getEvent(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var myEvents EventAttractions
	collection := client.Database("tourism").Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&myEvents)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myEvents)
}

//travelprovider
func createProvider(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myProvider TravelProvider
	_ = json.NewDecoder(request.Body).Decode(&myProvider)
	collection := client.Database("tourism").Collection("travelproviders")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, myProvider)
	json.NewEncoder(response).Encode(result)
}
func getProviders(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var myProvider []TravelProvider
	collection := client.Database("tourism").Collection("travelproviders")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var P TravelProvider
		cursor.Decode(&P)
		myProvider = append(myProvider, P)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myProvider)
}
func getProvider(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var myProvider TravelProvider
	collection := client.Database("tourism").Collection("travelproviders")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&myProvider)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(myProvider)
}

// reviews
func createReview(response http.ResponseWriter, request *http.Request) {}
func getReviews(response http.ResponseWriter, request *http.Request)   {}
func getReview(response http.ResponseWriter, request *http.Request)    {}

//tours
func createTour(response http.ResponseWriter, request *http.Request) {}
func getTours(response http.ResponseWriter, request *http.Request)   {}
func getTour(response http.ResponseWriter, request *http.Request)    {}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	string_conection_db := "mongodb+srv://<username>:<password>@<cluster_name>/<db_name>?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(string_conection_db)
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	//destination
	router.HandleFunc("/createDestination", createDestination).Methods("POST")
	router.HandleFunc("/getDestination", getDestinations).Methods("GET")
	router.HandleFunc("/getDestination/{id}", getDestination).Methods("GET")
	//user
	router.HandleFunc("/createUser", createUser).Methods("POST")
	router.HandleFunc("/getUser", getUsers).Methods("GET")
	router.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	//event
	router.HandleFunc("/createEvent", createEvent).Methods("POST")
	router.HandleFunc("/getEvent", getEvents).Methods("GET")
	router.HandleFunc("/getEvent/{id}", getEvents).Methods("GET")
	//tours
	router.HandleFunc("/createTours", createTour).Methods("POST")
	router.HandleFunc("/getTour", getTours).Methods("GET")
	router.HandleFunc("/getTour/{id}", getTour).Methods("GET")
	//reviews
	router.HandleFunc("/createReviews", createReview).Methods("POST")
	router.HandleFunc("/getReview", getReviews).Methods("GET")
	router.HandleFunc("/getReview/{id}", getReview).Methods("GET")
	//travelprovider
	router.HandleFunc("/createProvider", createProvider).Methods("POST")
	router.HandleFunc("/getProvider", getProviders).Methods("GET")
	router.HandleFunc("/getProvider/{id}", getProvider).Methods("GET")

	http.ListenAndServe(":8000", router)

}
