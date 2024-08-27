package tests

import (
	"testing"

	"github.com/emerald-lan/simple-todo-app/models"
	"github.com/emerald-lan/simple-todo-app/services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    
    mt.Run("Create Task", func(mt *mtest.T) {
        task := models.Task{
            Title: "Test Task",
        }

        mt.AddMockResponses(mtest.CreateSuccessResponse())
        taskService := services.NewTaskService(mt.DB)

        result, err := taskService.Create(task)
        assert.NoError(t, err)
        assert.NotNil(t, result)
    })
}

func TestFindAllTasks(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("Find All Tasks", func(mt *mtest.T) {
        expectedTasks := []models.Task{
            {ID: primitive.NewObjectID(), Title: "Task 1"},
            {ID: primitive.NewObjectID(), Title: "Task 2"},
        }

        first := mtest.CreateCursorResponse(1, "tasks.tasks", mtest.FirstBatch, bson.D{
            {Key: "_id", Value: expectedTasks[0].ID}, {Key: "title", Value: expectedTasks[0].Title},
        })
        second := mtest.CreateCursorResponse(1, "tasks.tasks", mtest.NextBatch, bson.D{
            {Key: "_id", Value: expectedTasks[1].ID}, {Key: "title", Value: expectedTasks[1].Title},
        })
        killCursors := mtest.CreateCursorResponse(0, "tasks.tasks", mtest.NextBatch)

        mt.AddMockResponses(first, second, killCursors)

        taskService := services.NewTaskService(mt.DB)

        tasks, err := taskService.FindAll()
        assert.NoError(t, err)
        assert.Equal(t, expectedTasks, tasks)
    })
}

func TestFindByIdTask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("Find Task by ID", func(mt *mtest.T) {
        id := primitive.NewObjectID()
        expectedTask := models.Task{ID: id, Title: "Test Task"}

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "tasks.tasks", mtest.FirstBatch, bson.D{
            {Key: "_id", Value: expectedTask.ID}, {Key: "title", Value: expectedTask.Title},
        }))

        taskService := services.NewTaskService(mt.DB)

        task, err := taskService.FindById(id)
        assert.NoError(t, err)
        assert.Equal(t, expectedTask, task)
    })
}

func TestUpdateTask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("Update Task", func(mt *mtest.T) {
        id := primitive.NewObjectID()
        task := models.Task{ID: id, Title: "Updated Task"}

        mt.AddMockResponses(mtest.CreateSuccessResponse())
        taskService := services.NewTaskService(mt.DB)

        result, err := taskService.Update(task)
        assert.NoError(t, err)
        assert.NotNil(t, result)
    })
}

func TestDeleteTask(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

    mt.Run("Delete Task", func(mt *mtest.T) {
        id := primitive.NewObjectID()

        mt.AddMockResponses(mtest.CreateSuccessResponse())
        taskService := services.NewTaskService(mt.DB)

        result, err := taskService.Delete(id)
        assert.NoError(t, err)
        assert.NotNil(t, result)
    })
}
