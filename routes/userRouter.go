package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/syed/go-jwt-authentication/controllers"
	"github.com/syed/go-jwt-authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
}
