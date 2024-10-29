package controller

import (
	"go-notifier/commons/utils/logger"
	"go-notifier/template-service/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	TemplateService services.ITemplateService
}

func NewTemplateController(templateService services.ITemplateService) *TemplateController {
	return &TemplateController{
		TemplateService: templateService,
	}
}

func (n *TemplateController) GetTemplate(ctx *gin.Context) {
	log := logger.GetLogger()
	methodName := "GetTemplate"
	log.Info(methodName, "started")

	userID := ctx.Param("template_id")

	template, err := n.TemplateService.GetTemplate(userID)

	if err != nil {
		log.Error(methodName, "error getting preferences", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info(methodName, "template", template)

	ctx.JSON(http.StatusOK, template)

}
