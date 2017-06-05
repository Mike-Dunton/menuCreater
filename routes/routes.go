package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mike-dunton/menuCreator/validators"
	"gopkg.in/gin-contrib/cors.v1"
)

// GetMainRouter - Handlers for route management
func GetMainRouter(authConfig AuthConfig) *gin.Engine {
	// Create a default gin router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	authMiddleware := initAuthMiddleware(authConfig)

	router.POST("/addUser", addUser)
	router.POST("/validate", validateUserName)

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
	validators.InitValidator()
	return router
}
