package router

import (
	"go-notifier/notification-service/internal/port/constants"
	"go-notifier/notification-service/internal/port/factory"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	//get all controllers
	factory := factory.NewPortFactory()

	//health check
	router.GET(constants.HEALTH, factory.HealthController.HealthCheck)

	api := router.Group(constants.API)
	{
		v1 := api.Group(constants.V1)
		{
			v1.POST("/send", factory.NotificationController.SendNotification)
		}
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
