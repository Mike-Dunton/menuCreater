package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/mike-dunton/menuCreator/controllers"
	"github.com/mike-dunton/menuCreator/models/recipe"
	"github.com/mike-dunton/menuCreator/mongo"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

var log = logrus.New()

// recipes middlewear
func getRecipeById(c *gin.Context) {
	recipeID := c.Param("id")
	if bson.IsObjectIdHex(recipeID) {
		recipeControler := &controllers.RecipeController{}
		err := recipeControler.NewController()
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		defer mongo.CloseSession(recipeControler.Service.MongoSession)
		code, body, _ := recipeControler.GetRecipeById(bson.ObjectIdHex(recipeID))
		c.JSON(code, body)
	} else {
		c.JSON(400, "no")
		return
	}
}

// recipes middlewear
func getAllRecipes(c *gin.Context) {
	recipeControler := &controllers.RecipeController{}
	err := recipeControler.NewController()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	code, body, _ := recipeControler.ListRecipes()
	mongo.CloseSession(recipeControler.Service.MongoSession)
	if code != 200 {
		log.WithFields(logrus.Fields{
			"code": code,
			"body": body,
		}).Warn("/status is not 200")
	}
	c.JSON(code, body)
}

func putRecipe(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var newRecipe recipeModel.Recipe
	err := json.Unmarshal(body, &newRecipe)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Unable to unmarshal your requst")
		c.JSON(400, "Unable to unmarshal your requst")
		return
	}
	err = validator.Validate(newRecipe)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Validation Error")
		c.JSON(400, err.Error())
		return
	}
	recipeControler := &controllers.RecipeController{}
	err = recipeControler.NewController()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	code, respBody, _ := recipeControler.NewRecipe(newRecipe)
	mongo.CloseSession(recipeControler.Service.MongoSession)
	c.JSON(code, respBody)
}
