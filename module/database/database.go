package database

import (
	"card-game-golang/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"path/filepath"
	"time"
)

// db ...
var db *mongo.Database

// Connect ...
func Connect() {
	envVars := config.GetEnv()

	curDir, _ := os.Getwd()
	dbName := envVars.Database.Name

	// database name when run test in controller
	if filepath.Base(curDir) == "controller" {
		dbName = envVars.Database.TestName
	}

	// configuring client to use the correct URI, but not yet connecting to it
	client, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.Uri))
	if err != nil {
		log.Fatal(err)
	}

	// timeout duration trying to connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect ...
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// ping ...
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	// db
	db = client.Database(dbName)
	fmt.Println("Database Connected to", dbName)
}
