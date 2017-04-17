package userService

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/mike-dunton/menuCreator/models/user"
	"github.com/mike-dunton/menuCreator/services"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logrus.New()

type (
	// buoyConfiguration contains settings for running the buoy service.
	recipeConfiguration struct {
		Database   string
		Collection string
	}
)

var Config recipeConfiguration

func init() {
	Config.Database = "menuCreator"
	Config.Collection = "users"
}

func GetUserByEmail(service *services.Service, Email string) (*userModel.User, error) {
	var user userModel.User
	executeFunc := func(collection *mgo.Collection) error {
		return collection.Find(bson.M{"email": Email}).One(&user)
	}
	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		return nil, err
	}
	log.Info("Got User")
	return &user, nil
}

func AddUser(service *services.Service, newUser userModel.User) (*userModel.User, error) {
	executeFunc := func(collection *mgo.Collection) error {
		newUser.ID = bson.NewObjectId()
		newUser.CreatedAt = time.Now()
		return collection.Insert(newUser)
	}

	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		return nil, err
	}
	log.Info("User Added")
	return &newUser, nil
}
