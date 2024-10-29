package response

import "time"

type TemplateDTO struct {
	TemplateID  string    `json:"template_id"` // Unique identifier for the template
	Title       string    `json:"title"`       // Title of the notification
	Body        string    `json:"body"`        // Body of the message with placeholders
	Description string    `json:"description"` // Brief description of the template
	CreatedAt   time.Time `json:"created_at"`  // Timestamp for when the template was created
	UpdatedAt   time.Time `json:"updated_at"`  // Timestamp for the last update of the template
}
