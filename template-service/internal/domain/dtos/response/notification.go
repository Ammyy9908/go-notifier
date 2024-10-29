package response

import "time"

type SendNotificationResponse struct {
	Status      string       `json:"status"`          // Status of the request, e.g., "Queued", "Sent", "Failed"
	MessageID   string       `json:"message_id"`      // Unique identifier for the notification message
	RecipientID string       `json:"recipient_id"`    // ID of the recipient for tracking
	Tracking    TrackingInfo `json:"tracking"`        // Tracking details if enabled
	Error       *ErrorDetail `json:"error,omitempty"` // Error details if the notification failed
}

type TrackingInfo struct {
	DeliveryStatus string    `json:"delivery_status,omitempty"` // Status of delivery, e.g., "Delivered", "Pending"
	ReadStatus     string    `json:"read_status,omitempty"`     // Status of the read confirmation, e.g., "Read", "Unread"
	DeliveredAt    time.Time `json:"delivered_at,omitempty"`    // Timestamp of successful delivery
	ReadAt         time.Time `json:"read_at,omitempty"`         // Timestamp of read confirmation
}

// ErrorDetail represents any error encountered during the notification process.
type ErrorDetail struct {
	Code      string `json:"code"`      // Error code, e.g., "DELIVERY_FAILURE", "INVALID_RECIPIENT"
	Message   string `json:"message"`   // Human-readable error message
	Retryable bool   `json:"retryable"` // Indicates if the error can be retried
}
