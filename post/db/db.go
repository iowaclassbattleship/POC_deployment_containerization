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
const databaseName = "post"
const collectionName = "post"

func getClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func GetPosts() ([]models.Post, error) {
	client := getClient()
	defer client.Disconnect(context.TODO())

	cursor, err := client.Database(databaseName).Collection(collectionName).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var posts []models.Post
	err = cursor.All(context.TODO(), &posts)
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}

func GetPostByID(id primitive.ObjectID) (models.Post, error) {
	client := getClient()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"_id": id}

	result := client.Database(databaseName).Collection(collectionName).FindOne(context.TODO(), filter)

	var post models.Post
	err := result.Decode(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func CreatePost(body models.PostRequestBody) (primitive.ObjectID, error) {
	client := getClient()
	defer client.Disconnect(context.TODO())

	coll := client.Database(databaseName).Collection(collectionName)
	doc := bson.D{
		{Key: "author", Value: body.Author},
		{Key: "title", Value: body.Title},
		{Key: "content", Value: body.Content},
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func UpdatePost(id primitive.ObjectID, body models.PostRequestBody) error {
	client := getClient()
	defer client.Disconnect(context.TODO())

	setTo := bson.D{
		{Key: "author", Value: body.Author},
		{Key: "title", Value: body.Title},
		{Key: "content", Value: body.Content},
	}

	postCollection := client.Database(databaseName).Collection(collectionName)
	update := bson.D{{Key: "$set", Value: setTo}}

	_, err := postCollection.UpdateByID(context.TODO(), id, update)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id primitive.ObjectID) error {
	client := getClient()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"_id": id}

	result, err := client.Database(databaseName).Collection(collectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("Post with id %s could not be deleted", id)
	}

	return nil
}
