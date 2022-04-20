package api

import (
	"context"

	"github.com/ChatKNU/pkg/types"
)

//go:generate mockery -name Manager -outpkg chatmocks -output ./chatmocks -dir .
type Manager interface {
	SendMessage(ctx context.Context, userID string, eventID string, msg types.IncomingChatMessage) error
	GetMessages(ctx context.Context, eventID string, query types.ChatQuery) (*types.ChatPage, error)
}
