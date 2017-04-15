package services

import (
	"github.com/mike-dunton/menuCreater/mongo"
	"gopkg.in/mgo.v2"
)

type (
	Service struct {
		MongoSession *mgo.Session
	}
)

func (service *Service) DBAction(databaseName string, collectionName string, ExecuteFunction mongo.DBCall) (err error) {
	err = ExecuteFunction(&mgo.Collection{})
	if err != nil {
		return err
	}
	return nil
}
