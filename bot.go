package flowbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type FlowBot struct {
	*tgbotapi.BotAPI
	chatStore   *ChatStore
	updates     tgbotapi.UpdatesChannel
	timeout     int
	timeoutText string
}

func (bot *FlowBot) HandleUpdates(handler func(*Chat)) {

	for update := range bot.updates {
		var chatID int64
		if update.Message != nil {
			chatID = update.Message.Chat.ID
		}
		if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID

		}
		chat, ok := bot.chatStore.Get(chatID)
		if !ok {
			chat = NewChat(chatID, bot)
			go handler(chat)
		}
		chat.ch <- &update
	}

}

func NewFlowBot(token string, timeout int, timeoutText string) (*FlowBot, error) {

	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	updates, err := api.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return &FlowBot{
		api,
		NewChatStore(),
		updates,
		timeout,
		timeoutText,
	}, nil
}
