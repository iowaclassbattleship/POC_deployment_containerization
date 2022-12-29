package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetPostByID(id primitive.ObjectID) (models.Post, error) {
	client := getClient()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"_id": id}

	result := client.Database("post").Collection("post").FindOne(context.TODO(), filter)

	var post models.Post
	err := result.Decode(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func CreatePost(body models.RequestCreatePost) primitive.ObjectID {
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

	return result.InsertedID.(primitive.ObjectID)
}

func DeletePost(id primitive.ObjectID) (int64, error) {
	client := getClient()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"_id": id}

	result, err := client.Database("post").Collection("post").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("Post with id %s could not be deleted", id)
	}

	return result.DeletedCount, nil
}
