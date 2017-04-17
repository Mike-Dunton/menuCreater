package routes

import (
	"github.com/gin-gonic/gin"
)

// GetMainRouter - Handlers for route management
func GetMainRouter(authConfig AuthConfig) *gin.Engine {

	// Create a default gin router
	router := gin.Default()
	authMiddleware := initAuthMiddleware(authConfig)

	router.POST("/addUser", addUser)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authMiddleware.LoginHandler)
	}

	recipeGroup := router.Group("/recipe")
	recipeGroup.Use(authMiddleware.MiddlewareFunc())
	{
		recipeGroup.GET("/", getAllRecipes)
		recipeGroup.GET("/:id", getRecipeById)
		recipeGroup.PUT("/", putRecipe)
	}

	// return the gin.Engine
	return router
}
