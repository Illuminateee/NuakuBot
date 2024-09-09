package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BotInitializer handles bot initialization
type BotInitializer struct{}

// InitializeBot initializes the bot using the provided token
func (b BotInitializer) InitializeBot(token string) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(token)
}
