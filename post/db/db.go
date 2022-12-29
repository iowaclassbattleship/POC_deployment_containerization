package db

import (
	"context"
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

func GetPosts() []models.Post {
	client := getClient()
	defer client.Disconnect(context.TODO())

	cursor, err := client.Database("post").Collection("post").Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var posts []models.Post
	err = cursor.All(context.TODO(), &posts)
	if err != nil {
		log.Fatal(err)
	}

	return posts
}

func GetPostByID(id string) models.Post {
	client := getClient()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"_id": id}

	result := client.Database("post").Collection("post").FindOne(context.TODO(), filter)
	var post models.Post
	result.Decode(&post)

	return post
}

func CreatePost(body models.RequestCreatePost) interface{} {
	client := getClient()
	defer client.Disconnect(context.TODO())

	coll := client.Database("post").Collection("post")
	doc := bson.D{
		{Key: "author", Value: "anybody"},
		{Key: "title", Value: body.Title},
		{Key: "content", Value: body.Content},
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

func DeletePost(id string) {
	client := getClient()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"_id": id}

	client.Database("post").Collection("post").DeleteOne(context.TODO(), filter)
}
