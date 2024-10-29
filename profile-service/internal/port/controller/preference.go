package controller

import (
	"go-notifier/commons/utils/logger"
	"go-notifier/profile-service/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PreferenceController struct {
	PreferenceService services.IPreferenceService
}

func NewPreferenceController(preferenceService services.IPreferenceService) *PreferenceController {
	return &PreferenceController{
		PreferenceService: preferenceService,
	}
}

func (n *PreferenceController) GetPreferences(ctx *gin.Context) {
	log := logger.GetLogger()
	methodName := "GetPreferences"
	log.Info(methodName, "started")

	userID := ctx.Param("user_id")

	preferences, err := n.PreferenceService.GetPreferences(userID)

	if err != nil {
		log.Error(methodName, "error getting preferences", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info(methodName, "preferences", preferences)

	ctx.JSON(http.StatusOK, preferences)

}
