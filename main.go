package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }

    // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
          panic(err)
        }
    }()

    // Send a ping to confirm a successful connection
    err = client.Ping(context.TODO(), readpref.Primary())
    if err != nil {
        panic(err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func main() {
	router := gin.Default()
    fmt.Println("Hello, World!")

	router.Run(":8080")
}