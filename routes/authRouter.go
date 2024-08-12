package routes

import (
	controller "golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup", controller.Signup())
	incomingRoutes.POST("/uesr/login", controller.Login())
}
