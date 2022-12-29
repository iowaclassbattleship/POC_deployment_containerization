package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"poc/post/models"
)

const uri = "mongodb://root:example@mongo:27017/?maxPoolSize=20&w=majority"

func getClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func TestInsert() {
	client := getClient()
	defer client.Disconnect(context.TODO())

	coll := client.Database("post").Collection("post")

	docs := []interface{}{
		models.Post{Author: "John Cena", Title: "The Art of the fart", Content: "Lol"},
		models.Post{Author: "John Cena", Title: "The Art of the fart", Content: "Lol"},
		models.Post{Author: "John Cena", Title: "The Art of the fart", Content: "Lol"},
		models.Post{Author: "John Cena", Title: "The Art of the fart", Content: "Lol"},
		models.Post{Author: "John Cena", Title: "The Art of the fart", Content: "Lol"},
	}

	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPosts() {
	client := getClient()
	defer client.Disconnect(context.TODO())

	cursor, err := client.Database("post").Collection("post").Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Post
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}
}

func CreatePost(body models.RequestCreatePost) error {
	return nil
}
