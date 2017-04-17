package controllers

import (
	"github.com/mike-dunton/menuCreator/models/user"
	"github.com/mike-dunton/menuCreator/mongo"
	"github.com/mike-dunton/menuCreator/services"
	"github.com/mike-dunton/menuCreator/services/userService"
)

type (
	// RecipeController contains common fields and behavior for all controllers
	UserController struct {
		services.Service
	}
)

func (userController *UserController) NewUser(newUser userModel.User) (int, *userModel.User, error) {
	user, err := userService.AddUser(&userController.Service, newUser)
	return 200, user, err
}

func (recipe *UserController) NewController() (err error) {
	recipe.Service.MongoSession, err = mongo.CopySession("main")
	return err
}
