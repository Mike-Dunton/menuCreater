package routes

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/mike-dunton/menuCreator/mongo"
	"github.com/mike-dunton/menuCreator/services"
	"github.com/mike-dunton/menuCreator/services/userService"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	SecretKey []byte
}

func initAuthMiddleware(authConfig AuthConfig) (authMiddleware *jwt.GinJWTMiddleware) {
	// the jwt middleware
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "menuCreatorZone",
		Key:        authConfig.SecretKey,
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Infof("Failed to hash Password %v", err)
				return email, false
			}
			authMongoSession, err := mongo.CopySession("main")
			defer mongo.CloseSession(authMongoSession)
			authService := services.Service{
				MongoSession: authMongoSession,
			}
			userToAuth, err := userService.GetUserByEmail(&authService, email)
			if err != nil {
				log.Infof("Failed to Get by email %v", err)
				return email, false
			}
			err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(userToAuth.Password))
			if err != nil {
				log.Infof("passwords do not match %v", err)
				return email, false
			}
			return email, true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		TimeFunc:    time.Now,
	}
	return authMiddleware
}
