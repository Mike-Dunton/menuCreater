package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/mike-dunton/menuCreater/mongo"
	"github.com/mike-dunton/menuCreater/routes"
	"gopkg.in/mgo.v2"
)

var log = logrus.New()
var MongoDialString string
var MongoSession mgo.Session

func init() {
	MongoDialString = os.Getenv("MENU_CREATOR_MONGO_DIAL_STRING")
}

func main() {
	log.Info("Starting...")
	err := mongo.Start(MongoDialString)
	if err != nil {
		fmt.Printf("There was an error starting the mongo connection: %v\n", err)
		os.Exit(1)
	}
	// Start server on port 8080
	routes.GetMainRouter().Run(":8080")
}
