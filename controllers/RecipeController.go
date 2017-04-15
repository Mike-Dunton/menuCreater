package controllers

import (
	recipeModel "github.com/mike-dunton/menuCreator/models/recipe"
	"github.com/mike-dunton/menuCreator/mongo"
	"github.com/mike-dunton/menuCreator/services"
	recipe "github.com/mike-dunton/menuCreator/services/recipeService"
)

type (
	// RecipeController contains common fields and behavior for all controllers
	RecipeController struct {
		services.Service
	}
)

func (recipeController *RecipeController) ListRecipes() (int, *[]recipeModel.Recipe, error) {
	recipes, err := recipe.ListRecipes(&recipeController.Service)
	return 200, recipes, err
}

func (recipeController *RecipeController) NewRecipe(newRecipe recipeModel.Recipe) (int, *[]recipeModel.Recipe, error) {
	recipes, err := recipe.AddRecipe(&recipeController.Service, newRecipe)
	return 200, recipes, err
}

func (recipe *RecipeController) NewController() (err error) {
	recipe.Service.MongoSession, err = mongo.CopySession("main")
	return err
}
