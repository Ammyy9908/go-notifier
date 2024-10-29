package impl

import (
	"encoding/json"
	"fmt"
	"go-notifier/commons/utils/api_client"
	"go-notifier/commons/utils/convertor"
	"go-notifier/commons/utils/injector"
	"go-notifier/commons/utils/logger"
	"go-notifier/commons/utils/queue_client"
	"go-notifier/notification-service/internal/domain/dtos/request"
	"go-notifier/notification-service/internal/domain/dtos/response"
)

type NotificationService struct {
	APIClient   api_client.IAPIClient
	QueueClient queue_client.IQueueInterface
}

func NewNotificationService(apiClient api_client.IAPIClient, queueClient queue_client.IQueueInterface) *NotificationService {
	return &NotificationService{
		APIClient:   apiClient,
		QueueClient: queueClient,
	}
}

func (n *NotificationService) SendNotification(payload *request.NotificationRequest) (interface{}, error) {
	log := logger.GetLogger()
	methodName := "SendNotification"
	log.Info(methodName, "started")

	api_response, api_err := n.APIClient.Get("/api/v1/preferences/" + payload.Recipients.UserID)
	if api_err != nil {
		log.Error(methodName, "error getting preferences", api_err)
		return nil, api_err
	}

	//Convert http.response to struct

	ResponseObject, BindError := convertor.TypeConverter[response.UserPreferences](api_response)
	if BindError != nil {
		log.Error(methodName, "error binding preferences", BindError)
		return nil, BindError
	}

	//Now Change the base url to template service

	n.APIClient.SetBaseURL("http://localhost:8082")

	template_response, template_err := n.APIClient.Get("/api/v1/templates/" + payload.Message.TemplateID)
	if template_err != nil {
		log.Error(methodName, "error getting template", template_err)
		return nil, template_err
	}

	template_object, template_bind_err := convertor.TypeConverter[response.TemplateDTO](template_response)
	if template_bind_err != nil {
		log.Error(methodName, "error binding template", template_bind_err)
		return nil, template_bind_err
	}

	log.Info(methodName, "template", template_object)

	//Check Notification Type

	enabledChannels := GetEnabledChannels(ResponseObject.Preferences.Channels)
	messageTitle := template_object.Title
	messageBody := injector.InjectPlaceholders(template_object.Body, payload.Message.Placeholders)

	finalPayload := map[string]interface{}{
		"recipient": payload.Recipients,
		"title":     messageTitle,
		"body":      messageBody,
		"priority":  payload.Priority,
		"metadata":  payload.Metadata,
	}
	payloadBytes, err := json.Marshal(finalPayload)
	if err != nil {
		log.Error(methodName, "error marshaling payload", err)
		return nil, err
	}

	log.Info(methodName, "Queue Client", n.QueueClient)
	for _, c := range enabledChannels {
		queueName := getQueueForChannel(c, payload.Priority)

		queue_error := n.QueueClient.Enqueue(queueName, payloadBytes)
		if queue_error != nil {
			log.Error(methodName, "error enqueuing message", queue_error)
			return nil, queue_error
		}
	}

	log.Info(methodName, "notification sent")
	n.APIClient.SetBaseURL("http://localhost:8081")
	return nil, nil
}

func GetEnabledChannels(channels response.Channels) []string {
	var enabledChannels []string

	if channels.Email.Enabled {
		enabledChannels = append(enabledChannels, "Email")
	}
	if channels.SMS.Enabled {
		enabledChannels = append(enabledChannels, "SMS")
	}
	if channels.WebPush.Enabled {
		enabledChannels = append(enabledChannels, "WebPush")
	}

	return enabledChannels
}

// getQueueForChannel returns the correct queue name based on priority
func getQueueForChannel(channel, priority string) string {
	if priority == "high" {
		return fmt.Sprintf("%s_high_priority", channel)
	}
	return fmt.Sprintf("%s_standard", channel)
}
