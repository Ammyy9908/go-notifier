package dto

type NotificationDTO struct {
	Recipient RecipientDTO         `json:"recipient"` // Contains recipient information
	Title     string               `json:"title"`     // Title of the notification message
	Body      string               `json:"body"`      // Body of the notification message with placeholders filled
	Priority  string               `json:"priority"`  // Notification priority (e.g., "high", "standard")
	Metadata  NotificationMetadata `json:"metadata"`  // Additional metadata for the notification
}

// RecipientDTO represents the structure for recipient information.
type RecipientDTO struct {
	UserID      string `json:"user_id"`      // Unique identifier for the user
	Email       string `json:"email"`        // Email address of the recipient
	PhoneNumber string `json:"phone_number"` // Phone number of the recipient
	DeviceToken string `json:"device_token"` // Device token for push notifications
}

// NotificationMetadata represents additional metadata for the notification.
type NotificationMetadata struct {
	CampaignID    string `json:"campaign_id"`    // ID of the campaign for tracking
	ScheduledTime string `json:"scheduled_time"` // Scheduled time for notification delivery
}
