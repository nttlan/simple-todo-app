package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Title     string             `bson:"title,required" json:"title"`
    Completed bool               `bson:"completed" json:"completed"`
}