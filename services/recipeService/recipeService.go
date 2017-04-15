package recipeService

import (
	"encoding/json"

	"github.com/mike-dunton/menuCreater/models/recipe"
	"github.com/mike-dunton/menuCreater/services"
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

var mockRecipes = []byte("[{\"id\":\"123\",\"name\":\"Recipe1\",\"ingredients\":[{\"name\":\"Ingredient1\",\"isOptional\":false}]},{\"id\":\"321\",\"name\":\"Recipe2\",\"ingredients\":[{\"name\":\"Ingredient3\",\"isOptional\":false}]}]")

func init() {
	Config.Database = "menuCreater"
	Config.Collection = "recipes"
}

// ListRecipes retrieves all Recipes
func ListRecipes(service *services.Service) (*[]recipeModel.Recipe, error) {

	var recipes []recipeModel.Recipe
	executeFunc := func(collection *mgo.Collection) error {
		log.Info("Getting Recipes")
		//return collection.Find(nil).All(&recipes)
		return json.Unmarshal(mockRecipes, &recipes)
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
		//Insert New Recipe
		return nil
	}

	if err := service.DBAction(Config.Database, Config.Collection, executeFunc); err != nil {
		return nil, err
	}
	log.Info("Recipe Added")
	return &recipeModel.Recipe{
		Name: "Fake",
		Ingredients: []recipeModel.RecipeIngredients{
			recipeModel.RecipeIngredients{
				Name:       "FakeIngredient1",
				IsOptional: false,
			},
		},
	}, nil
}
