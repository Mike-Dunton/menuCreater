package mongo

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var log = logrus.New()
var MongoManager *mongoSessionManager

type ConnectionConfig struct {
	DBAddrs []string
	AuthDB  string
	User    string
	Pass    string
}

type mongoSessionManager struct {
	sessions map[string]*mgo.Session
}

type DBCall func(*mgo.Collection) error

func Execute(mongoSession *mgo.Session, databaseName string, collectionName string, executeFunction DBCall) (err error) {
	log.Infof("Execute: Database[%s] Collection[%s]", databaseName, collectionName)

	// capture mongo collection
	collection := mongoSession.DB(databaseName).C(collectionName)
	// pass collection to calling  function
	err = executeFunction(collection)
	if err != nil {
		log.Infof("Execute: Database[%s] Collection[%s] Err %v", databaseName, collectionName, err)
		return err
	}
	log.Infof("Execute: Database[%s] Collection[%s] Complete", databaseName, collectionName)

	return err
}

func Start(config ConnectionConfig) error {
	MongoManager = &mongoSessionManager{
		sessions: map[string]*mgo.Session{},
	}

	return CreateSession("main", config)
}

func CreateSession(sessionName string, config ConnectionConfig) (err error) {
	// Establish a new session
	dialInfo := &mgo.DialInfo{
		Addrs:    config.DBAddrs,
		Database: config.AuthDB,
		Username: config.User,
		Password: config.Pass,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Timeout: time.Second * 10,
	}
	fmt.Printf("dialInfo: %v\n", dialInfo)
	mongoSession, err := mgo.DialWithInfo(dialInfo)
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
	log.Infof("Cloaning Session \n")
	clonedSession = session.Clone()

	return clonedSession, err
}

// CloseSession puts the connection back into the pool
func CloseSession(mongoSession *mgo.Session) {
	log.Infof("Closing Session \n")
	mongoSession.Close()
}
