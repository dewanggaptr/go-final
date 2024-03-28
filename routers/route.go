package routes

import (
	"final-project/controllers"
	// order "challenge-api/pkg/controllers"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func StartApp() *gin.Engine  {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	return r
}