package types

import "time"

type IncomingChatMessage struct {
	EventID  string
	SenderID string
	IsBot    bool
	LocalID  string
	Text     string
}

type ChatMessage struct {
	ID        string    `firestore:"id"`
	Timestamp time.Time `firestore:"timestamp"`
	SenderID  string    `firestore:"sender_id"`
	IsBot     bool      `firestore:"is_bot"`
	Text      string    `firestore:"text"`
	URL       string    `firestore:"url"`
}

type ChatPage struct {
	Messages      []ChatMessage
	Attendees     []Attendee
	LastTimestamp *time.Time
}

type Attendee struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ChatQuery struct {
	StartFrom int64 `form:"start_from"`
	Limit     int   `form:"limit"`
}
