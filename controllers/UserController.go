package controllers

import (
	"github.com/Sirupsen/logrus"
	"github.com/mike-dunton/menuCreator/errors"
	"github.com/mike-dunton/menuCreator/models/user"
	"github.com/mike-dunton/menuCreator/mongo"
	"github.com/mike-dunton/menuCreator/services"
	"github.com/mike-dunton/menuCreator/services/userService"

	"net/http"
)

var log = logrus.New()

type (
	// RecipeController contains common fields and behavior for all controllers
	UserController struct {
		services.Service
	}
)

func (userController *UserController) NewUser(newUser userModel.User) (int, *userModel.User, error) {
	existingUser, err := userService.GetUserByEmail(&userController.Service, newUser.Email)
	if existingUser != nil {
		log.Infof("User Exists")
		return http.StatusPreconditionFailed, nil, errors.ErrUserExists
	}
	user, err := userService.AddUser(&userController.Service, newUser)
	return 200, user, err
}

func (recipe *UserController) NewController() (err error) {
	recipe.Service.MongoSession, err = mongo.CopySession("main")
	return err
}
