package services

import (
	"context"

	"github.com/emerald-lan/simple-todo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	collection := db.Collection("tasks")
	return &TaskService{collection}
}

func (s *TaskService) CreateOne(task models.Task) (*mongo.InsertOneResult, error) {
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