package services

import (
	"github.com/mike-dunton/menuCreator/mongo"
	"gopkg.in/mgo.v2"
)

type (
	Service struct {
		MongoSession *mgo.Session
	}
)

func (service *Service) DBAction(databaseName string, collectionName string, executeFunction mongo.DBCall) (err error) {
	return mongo.Execute(service.MongoSession, databaseName, collectionName, executeFunction)
}
