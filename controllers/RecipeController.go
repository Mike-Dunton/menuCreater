package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/mike-dunton/menuCreater/mongo"
	"github.com/mike-dunton/menuCreater/services"
	recipe "github.com/mike-dunton/menuCreater/services/recipeService"
)

type (
	// RecipeController contains common fields and behavior for all controllers
	RecipeController struct {
		services.Service
	}
)

func (recipeController *RecipeController) ListRecipes() (int, string, error) {
	recipes, err := recipe.ListRecipes(&recipeController.Service)
	fmt.Println(recipes)
	fmt.Println(err)
	result, err := json.Marshal(recipes)
	return 200, string(result), nil
}

func (recipe *RecipeController) NewController() (err error) {
	recipe.Service.MongoSession, err = mongo.CopySession("main")
	return err
}
