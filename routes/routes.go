package routes

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/mike-dunton/menuCreater/controllers"
)

var log = logrus.New()

// GetMainRouter - Handlers for route management
func GetMainRouter() *gin.Engine {

	// Create a default gin router
	router := gin.Default()
	// GET /status
	router.GET("/recipes", getRecipes)

	// return the gin.Engine
	return router
}

// recipes middlewear
func getRecipes(c *gin.Context) {
	recipeControler := &controllers.RecipeController{}
	err := recipeControler.NewController()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	code, body, _ := recipeControler.ListRecipes()
	if code != 200 {
		log.WithFields(logrus.Fields{
			"code": code,
			"body": body,
		}).Warn("/status is not 200")
	}
	c.JSON(code, body)
}
