package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/mike-dunton/menuCreator/mongo"
	"github.com/mike-dunton/menuCreator/routes"
	"gopkg.in/mgo.v2"
)

var log = logrus.New()
var connectionConfig mongo.ConnectionConfig
var authConfig routes.AuthConfig

var MongoSession mgo.Session

func init() {
	connectionConfig = mongo.ConnectionConfig{
		DBAddrs: strings.Split(os.Getenv("MENU_CREATOR_MONGO_DB_ADDRS"), ","),
		AuthDB:  os.Getenv("MENU_CREATOR_MONGO_AUTH_DB"),
		User:    os.Getenv("MENU_CREATOR_MONGO_USER"),
		Pass:    os.Getenv("MENU_CREATOR_MONGO_PASS"),
	}

	authConfig = routes.AuthConfig{
		SecretKey: []byte(os.Getenv("MENU_CREATOR_AUTH_SECRET_KEY")),
	}
}

func main() {
	log.Info("Starting...")
	log.Infof("Dial String %v", connectionConfig)
	err := mongo.Start(connectionConfig)
	if err != nil {
		fmt.Printf("There was an error starting the mongo connection: %v\n", err)
		os.Exit(1)
	}
	// Start server on port 8080
	routes.GetMainRouter(authConfig).Run(":8080")
}
