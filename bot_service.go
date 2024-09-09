package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BotService handles the main bot logic
type BotService struct {
	BotAPI    *tgbotapi.BotAPI
	UpdateSvc UpdateService
}

// UpdateService defines an interface for handling updates
type UpdateService interface {
	HandleUpdate(update tgbotapi.Update)
}
