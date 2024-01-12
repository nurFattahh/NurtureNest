package interfaces

import (
	"NurtureNest/pkg/domain"
	"context"
)

type ChatBotRepository interface {
	GenerateChatBot(ctx context.Context, model domain.ChatBot) (domain.ChatBot, error)
	GetAllHistoryChatBot(ctx context.Context, user_id uint) ([]domain.ChatBot, error)
}
