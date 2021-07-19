package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// My sample db's collection has three fields, name-age-title
// So, I create a struct with three fields, which will be used to decode to get the data
//
type Employee struct {
	Name string
	Age string
	Title string
}


func main() {

	// Initialization || localhost mongodb and used (default) port
	//
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf( clientOptions ), "\n")
	// Create the client of mongodb connection
	//
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Connection to MongoDB failed ...")
		panic(err)
	}


	// Mandatory to iterate through all documents our collection
	//
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// The db and collections to use
	//
	col := client.Database("intratel").Collection("employees")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	iter, err := col.Find(context.TODO(), bson.D{})

    if err != nil {
        fmt.Println("Documents ERROR occured ...")
		// close / defer channel
        defer iter.Close(ctx)

    } else {
		// Just iterate through all entries
		//
		counter := 1
        for iter.Next(ctx) {

            var result bson.M
            err := iter.Decode(&result)

            if err != nil {
                fmt.Println("iter.Next() error:", err)
                panic(err)
               
            } else {
                // fmt.Println("\nresult type:", reflect.TypeOf(result))
                fmt.Println("Entry:", counter, " : " , result)
				counter += 1
            }
        }
    }
}

