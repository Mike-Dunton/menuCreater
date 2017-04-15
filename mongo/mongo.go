package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var MongoManager *mongoSessionManager

type mongoSessionManager struct {
	sessions map[string]*mgo.Session
}

type DBCall func(*mgo.Collection) error

func Start(MongoDialString string) error {
	MongoManager = &mongoSessionManager{
		sessions: map[string]*mgo.Session{},
	}

	return CreateSession("main", MongoDialString)
}

func CreateSession(sessionName string, MongoDialString string) (err error) {
	// Establish a new session
	mongoSession, err := mgo.Dial(MongoDialString)
	if err != nil {
		fmt.Printf("MongoDB dial err %v\n", err)
		return err
	}
	MongoManager.sessions[sessionName] = mongoSession

	return err
}

// CopySession makes a clone of the specified session for client use
func CopySession(sessionId string) (clonedSession *mgo.Session, err error) {
	session := MongoManager.sessions[sessionId]

	// Clone the master session
	clonedSession = session.Clone()

	return clonedSession, err
}

// CloseSession puts the connection back into the pool
func CloseSession(mongoSession *mgo.Session) {
	mongoSession.Close()
}
