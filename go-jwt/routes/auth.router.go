package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ramalloc/go-jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.Singnup())
	incomingRoutes.POST("/users/login", controllers.Login())
}