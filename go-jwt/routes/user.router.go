package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ramalloc/go-jwt/controllers"
	"github.com/ramalloc/go-jwt/middlewares"
)


func UserRoutes(incomingRoutes *gin.Engine)  {
    incomingRoutes.Use(middlewares.Authenticate())
    incomingRoutes.GET("/users", controllers.GetUsers())
    incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}