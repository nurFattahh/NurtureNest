package repository

import (
	"NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	"context"

	"gorm.io/gorm"
)

type chatBotDatabase struct {
	DB *gorm.DB
}

func NewChatBotRepository(DB *gorm.DB) interfaces.ChatBotRepository {
	return &chatBotDatabase{DB}
}

func (c *chatBotDatabase) GenerateChatBot(ctx context.Context, model domain.ChatBot) (domain.ChatBot, error) {
	err := c.DB.Create(&model).Error

	return model, err
}

func (c *chatBotDatabase) GetAllHistoryChatBot(ctx context.Context, user_id uint) ([]domain.ChatBot, error) {
	var chats []domain.ChatBot
	err := c.DB.Where("user_id = ?", user_id).Find(&chats).Error

	return chats, err

}
