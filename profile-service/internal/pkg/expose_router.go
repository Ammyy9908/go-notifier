package pkg

import (
	"go-notifier/profile-service/internal/port/router"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := router.NewRouter()
	return router
}
