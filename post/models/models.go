package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Author  string             `json:"author" bson:"author,omitempty"`
	Title   string             `json:"title" bson:"title,omitempty"`
	Content string             `json:"content" bson:"content,omitempty"`
}

type RequestCreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
