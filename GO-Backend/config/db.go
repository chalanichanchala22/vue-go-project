package config

import (
    "context" //context is used to manage timeouts or cancel operations (like connecting to MongoDB
    "log" //log is used to print messages to the console.
    "os" //os lets you read environment variables 
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client //stores the MongoDB client

func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI is not set in .env file")
    }

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    Client = client
    log.Println("Connected to MongoDB")
}

func GetDatabase() *mongo.Database {
    dbName := os.Getenv("MONGO_DB")
    if dbName == "" {
        log.Fatal("MONGO_DB is not set in .env file")
    }
    return Client.Database(dbName)
}