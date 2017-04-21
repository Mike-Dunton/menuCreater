package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/mike-dunton/menuCreator/controllers"
	"github.com/mike-dunton/menuCreator/errors"
	"github.com/mike-dunton/menuCreator/models/user"
	"github.com/mike-dunton/menuCreator/mongo"
	"gopkg.in/validator.v2"
)

// user middlewear
func addUser(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var user userModel.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Unable to unmarshal your requst")
		c.String(400, "Unable to unmarshal your requst")
		return
	}
	err = validator.Validate(user)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Validation Error")
		c.JSON(409, errors.ErrValidatingSignUp)
		return
	}
	userController := &controllers.UserController{}
	err = userController.NewController()
	defer mongo.CloseSession(userController.Service.MongoSession)
	if err != nil {
		c.String(409, err.Error())
		return
	}
	code, respBody, err := userController.NewUser(user)
	if err != nil {
		c.String(code, err.Error())
		return
	}
	c.JSON(code, respBody)
}
