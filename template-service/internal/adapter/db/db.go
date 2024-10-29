package db

import (
	"fmt"
	"go-notifier/commons/utils/db"
	"go-notifier/commons/utils/logger"

	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbObject *mongo.Database

func Init() {
	log := logger.GetLogger()

	// Get database connection details
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")

	// Construct the database URL based on the presence of username/password
	var dbUrl string
	if username == "" && password == "" {
		dbUrl = fmt.Sprintf("mongodb://%s:27017", host) // No auth for local
	} else {
		dbUrl = fmt.Sprintf("mongodb://%s:%s@%s:27017", username, password, host) // With auth
	}
	log.Info("Database URL: ", dbUrl)

	// Initialize MongoDB client
	client, err := db.NewDBClient(log, dbUrl)
	if err != nil {
		log.Fatal("Error ", err.Error(), " error connecting mongo db")
	}

	dbObject = client.Database(viper.GetString("database.name"))
	_, err = mongodb.WithInstance(client, &mongodb.Config{
		DatabaseName: viper.GetString("database.name"),
	})
	if err != nil {
		log.Error("Failed to configure MongoDB instance:", err)
	}
}

func GetDatabase() *mongo.Database {
	return dbObject
}
