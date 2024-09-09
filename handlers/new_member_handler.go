package handlers

import (
    "log"
    "strings"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// NewMemberHandler handles new members joining the group
type NewMemberHandler struct {
    BotAPI *tgbotapi.BotAPI
}

// HandleUpdate handles the logic when a new member joins
func (n NewMemberHandler) HandleUpdate(update tgbotapi.Update) {
    if update.Message != nil && update.Message.NewChatMembers != nil {
        for _, newMember := range update.Message.NewChatMembers {
            // Send the image first
            photo := tgbotapi.NewPhoto(
                update.Message.Chat.ID,
                tgbotapi.FilePath("assets/img/nuaku.jpeg"), //
            )
            _, err := n.BotAPI.Send(photo)
            if err != nil {
                log.Printf("Failed to send image: %v", err)
                continue // If sending the image fails, don't proceed to send the welcome message
            }

            // Create the welcome message with escaped MarkdownV2 characters
            welcomeMsg := tgbotapi.NewMessage(update.Message.Chat.ID, 
                "Selamat datang *" + escapeMarkdownV2(newMember.FirstName) + "* di *NUAKU HEALTHY CHANNEL*\\!\n\n" +
                "Silahkan perkenalkan diri anda\\.\n\n" +
                "*NUAKU HEALTHY CHANNEL* adalah tempat anda mendapat informasi untuk meningkatkan kesehatan anda\\. Kami menawarkan program:\n" +
                "\\- \\*Bincang Sehat\\*\n" + // Escape hyphen and asterisk
                "\\- \\*Ruang Praktik Terapi\\*\n" +
                "\\- \\*Latihan Olah Nafas untuk Kesehatan\\*\n\n" +
                "_Semoga anda mendapatkan banyak manfaat di sini\\!_")
            
            // Create the inline keyboard with buttons that redirect to specific links
            inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
                tgbotapi.NewInlineKeyboardRow(
                    tgbotapi.NewInlineKeyboardButtonURL("Telegram", "https://t.me/NuakuHealthy"),
                ),
                tgbotapi.NewInlineKeyboardRow(
                    tgbotapi.NewInlineKeyboardButtonURL("Youtube Nuaku", "https://www.youtube.com/@NuakuHealthy"),
                ),
                tgbotapi.NewInlineKeyboardRow(
                    tgbotapi.NewInlineKeyboardButtonURL("Facebook", "https://www.facebook.com/people/NUAKU-Healthy-LINE/100057369253016/?_rdr"),
                ),
                tgbotapi.NewInlineKeyboardRow(
                    tgbotapi.NewInlineKeyboardButtonURL("Whatsapp", "https://wa.me/085294421993"),
                ),
            )
            
            // Set the parse mode to MarkdownV2 to support formatting
            welcomeMsg.ParseMode = "MarkdownV2"
            
            // Add the inline keyboard to the message
            welcomeMsg.ReplyMarkup = inlineKeyboard

            // Send the welcome message
            _, err = n.BotAPI.Send(welcomeMsg)
            if err != nil {
                log.Printf("Failed to send message: %v", err)
            }
        }
    }
}

// escapeMarkdownV2 is used to escape special characters in MarkdownV2 format
func escapeMarkdownV2(text string) string {
    // List of special characters in MarkdownV2 that need to be escaped
    specialChars := []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
    
    // Escape each special character by prepending it with a backslash
    for _, char := range specialChars {
        text = strings.ReplaceAll(text, char, "\\"+char)
    }
    
    return text
}
