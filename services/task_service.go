package services

import (
	"context"

	"github.com/emerald-lan/simple-todo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	collection := db.Collection("tasks")
	return &TaskService{collection}
}

func (s *TaskService) Create(task models.Task) (*mongo.InsertOneResult, error) {
	result, err := s.collection.InsertOne(context.TODO(), task)
	return result, err
}

func (s *TaskService) FindAll() ([]models.Task, error) {
	filter := bson.D{}

	cursor, err := s.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []models.Task
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (c *TaskService) FindById(id primitive.ObjectID) (models.Task, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var result models.Task
	err := c.collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (c *TaskService) Update(task models.Task) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: task.ID}}
	update := bson.D{{Key: "$set", Value: task}}
	result, err := c.collection.UpdateOne(context.TODO(), filter, update)
	return result, err
}

func (c *TaskService) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := c.collection.DeleteOne(context.TODO(), filter)
	return result, err
}