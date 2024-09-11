package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "nuakubot/handlers" // Replace with your actual path
)

// Initialize and start the bot
func main() {
    // Get the bot token directly from environment variables
    botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if botToken == "" {
        log.Fatal("Bot token is not set. Please check the environment variables.")
    }

    // Initialize the bot
    botInitializer := BotInitializer{}
    bot, err := botInitializer.InitializeBot(botToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    // BotService handles updates and delegates logic to handlers
    botService := BotService{
        BotAPI:    bot,
        UpdateSvc: handlers.NewMemberHandler{BotAPI: bot}, // Correctly reference the handler
    }

    // Listen for updates
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        botService.UpdateSvc.HandleUpdate(update)
    }
}
