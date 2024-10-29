package controller

import (
	"go-notifier/commons/utils/logger"
	"go-notifier/commons/utils/validator"
	"go-notifier/notification-service/internal/domain/dtos/request"
	"go-notifier/notification-service/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	rv validator.IRequestValidator
	ns services.INotificationService
}

func NewNotificationController(rv validator.IRequestValidator, ns services.INotificationService) *NotificationController {
	return &NotificationController{
		rv: rv,
		ns: ns,
	}
}

func (n *NotificationController) SendNotification(ctx *gin.Context) {
	log := logger.GetLogger()
	methodName := "SendNotification"
	log.Info(methodName, "started")

	var Payload *request.NotificationRequest
	if err := ctx.ShouldBindJSON(&Payload); err != nil {
		log.Error(methodName, "error binding json", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := n.rv.Validate(Payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, domainError := n.ns.SendNotification(Payload)
	if domainError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": domainError})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
