package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Author  string             `bson:"author,omitempty"`
	Title   string             `bson:"title,omitempty"`
	Content string             `bson:"content,omitempty"`
}

type RequestCreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
