package services

import "gopkg.in/mgo.v2"

type (
	Service struct {
		MongoSession string
	}
	DBCall func(*mgo.Collection) error
)

func (service *Service) DBAction(databaseName string, collectionName string, ExecuteFunction DBCall) (err error) {
	err = ExecuteFunction(&mgo.Collection{})
	if err != nil {
		return err
	}
	return nil
}
