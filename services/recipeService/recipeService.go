package recipeService

import (
	"github.com/mike-dunton/menuCreator/models/recipe"
	"github.com/mike-dunton/menuCreator/services"
	"gopkg.in/mgo.v2"

	"github.com/Sirupsen/logrus"
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
	Config.Collection = "recipes"
}

// ListRecipes retrieves all Recipes
func ListRecipes(service *services.Service) (*[]recipeModel.Recipe, error) {

	var recipes []recipeModel.Recipe
	executeFunc := func(collection *mgo.Collection) error {
		log.Info("Getting Recipes")
		return collection.Find(nil).All(&recipes)
	}

	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		if err != nil {
			return nil, err
		}
	}

	log.Info("Got Recipes")
	return &recipes, nil
}

func AddRecipe(service *services.Service, newRecipe recipeModel.Recipe) (*[]recipeModel.Recipe, error) {
	executeFunc := func(collection *mgo.Collection) error {
		return collection.Insert(newRecipe)
	}

	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		return nil, err
	}
	log.Info("Recipe Added")
	return ListRecipes(service)
}
