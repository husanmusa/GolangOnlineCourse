package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lesson31/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("go_online").Collection("students")
	name := "John Doe"

	_, err = coll.InsertOne(context.Background(), bson.D{{"name", "John Doe"}})
	if err != nil {
		log.Fatal(err)
	}
	a := []model.User{{Name: "John Doe", Age: 20}}

	_, err = coll.InsertMany(context.Background(), a)

	var result bson.D
	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
