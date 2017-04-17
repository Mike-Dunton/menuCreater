package recipeService

import (
	"github.com/Sirupsen/logrus"
	"github.com/mike-dunton/menuCreator/models/recipe"
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
	Config.Collection = "recipes"
}

func GetRecipe(service *services.Service, recipeID bson.ObjectId) (*recipeModel.Recipe, error) {
	var recipe recipeModel.Recipe
	executeFunc := func(collection *mgo.Collection) error {
		log.Info("Getting Recipe")
		return collection.FindId(recipeID).One(&recipe)
	}
	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		return nil, err
	}
	log.Info("Got Recipe")
	return &recipe, nil
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

func AddRecipe(service *services.Service, newRecipe recipeModel.Recipe) (*recipeModel.Recipe, error) {
	executeFunc := func(collection *mgo.Collection) error {
		newRecipe.ID = bson.NewObjectId()
		return collection.Insert(newRecipe)
	}

	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		return nil, err
	}
	log.Info("Recipe Added")
	return &newRecipe, nil
}
