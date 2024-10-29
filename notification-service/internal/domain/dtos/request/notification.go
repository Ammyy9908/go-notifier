package request

import "time"

type Recipient struct {
	UserID      string `json:"user_id" validate:"required"`
	Email       string `json:"email,omitempty" validate:"email"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"e164"` // E.164 format for phone numbers
	DeviceToken string `json:"device_token,omitempty"`
}

// Message represents the message details for the /send request.
type Message struct {
	TemplateID   string            `json:"template_id" validate:"required"`
	Placeholders map[string]string `json:"placeholders,omitempty"` // Map of placeholder values
}

// RetryPolicy defines the retry settings for the notification.
type RetryPolicy struct {
	MaxRetries   int    `json:"max_retries" validate:"required,gte=0"`
	RetryBackoff string `json:"retry_backoff" validate:"required,oneof=fixed exponential"`
}

// Tracking defines the tracking options for the notification.
type Tracking struct {
	TrackDelivery bool `json:"track_delivery"`
	TrackRead     bool `json:"track_read"`
}

// Metadata contains additional optional data for the notification.
type Metadata struct {
	CampaignID    string    `json:"campaign_id,omitempty"`
	ScheduledTime time.Time `json:"scheduled_time,omitempty"`
}

type NotificationRequest struct {
	Recipients  Recipient   `json:"recipient" validate:"required"`
	Message     Message     `json:"message" validate:"required"`
	RetryPolicy RetryPolicy `json:"retry_policy,omitempty"`
	Tracking    Tracking    `json:"tracking,omitempty"`
	Metadata    Metadata    `json:"metadata,omitempty"`
	Priority    string      `json:"priority,omitempty"`
}
