package routes

import (
	"github.com/gin-gonic/gin"
)

// GetMainRouter - Handlers for route management
func GetMainRouter() *gin.Engine {

	// Create a default gin router
	router := gin.Default()
	recipeGroup := router.Group("/recipe")
	{
		recipeGroup.GET("/", getAllRecipes)
		recipeGroup.GET("/:id", getRecipeById)
		recipeGroup.PUT("/", putRecipe)
	}

	// return the gin.Engine
	return router
}
