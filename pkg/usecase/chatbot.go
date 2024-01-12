package usecase

import (
	"context"

	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	services "NurtureNest/pkg/usecase/interface"
)

type chatBotUseCase struct {
	chatBotUseCase interfaces.ChatBotRepository
}

func NewChatBotUseCase(repo interfaces.ChatBotRepository) services.ChatBotUseCase {
	return &chatBotUseCase{
		chatBotUseCase: repo,
	}
}

func (c *chatBotUseCase) GenerateChatBot(ctx context.Context, model domain.ChatBot) (domain.ChatBot, error) {
	generate, err := c.chatBotUseCase.GenerateChatBot(ctx, model)

	return generate, err
}

func (c *chatBotUseCase) GetAllHistoryChatBot(ctx context.Context, user_id uint) ([]domain.ChatBot, error) {
	history, err := c.chatBotUseCase.GetAllHistoryChatBot(ctx, user_id)
	return history, err
}
