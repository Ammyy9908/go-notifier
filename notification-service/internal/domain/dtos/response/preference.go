package response

import "time"

// UserPreferences represents the user preferences document.
type UserPreferences struct {
	ID          string      `json:"_id" bson:"_id"`                   // MongoDB ObjectId
	UserID      string      `json:"user_id" bson:"user_id"`           // Unique identifier for the user
	Preferences Preferences `json:"preferences" bson:"preferences"`   // User's notification preferences
	LastUpdated time.Time   `json:"last_updated" bson:"last_updated"` // Timestamp of the last update
}

// Preferences holds the user's notification preferences for channels and types.
type Preferences struct {
	Channels          Channels          `json:"channels" bson:"channels"`                     // Preferred channels for notifications
	NotificationTypes NotificationTypes `json:"notification_types" bson:"notification_types"` // Preferences by notification type
}

// Channels defines preferences for each notification channel (email, sms, web push).
type Channels struct {
	Email   ChannelSettings `json:"email" bson:"email"`
	SMS     ChannelSettings `json:"sms" bson:"sms"`
	WebPush ChannelSettings `json:"web_push" bson:"web_push"`
}

// ChannelSettings represents the settings for a specific notification channel.
type ChannelSettings struct {
	Enabled   bool   `json:"enabled" bson:"enabled"`                             // Whether this channel is enabled
	Frequency string `json:"frequency,omitempty" bson:"frequency,omitempty"`     // Frequency of notifications (e.g., "daily", "instant")
	TimeOfDay string `json:"time_of_day,omitempty" bson:"time_of_day,omitempty"` // Preferred time for scheduled notifications
}

// NotificationTypes holds preferences for specific types of notifications.
type NotificationTypes struct {
	Marketing     NotificationTypeSettings `json:"marketing" bson:"marketing"`
	Transactional NotificationTypeSettings `json:"transactional" bson:"transactional"`
	Reminders     NotificationTypeSettings `json:"reminders" bson:"reminders"`
}

// NotificationTypeSettings represents settings for a specific notification type.
type NotificationTypeSettings struct {
	Enabled   bool   `json:"enabled" bson:"enabled"`                             // Whether notifications of this type are enabled
	TimeOfDay string `json:"time_of_day,omitempty" bson:"time_of_day,omitempty"` // Preferred time for reminders
}
