package flowbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type KbrdBtn struct {
	Text string
	Data string
}

type Kbrd [][]KbrdBtn

func NewKbrd(k *Kbrd) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, r := range *k {
		var row []tgbotapi.InlineKeyboardButton
		for _, b := range r {
			row = append(row, tgbotapi.NewInlineKeyboardButtonData(b.Text, b.Data))
		}
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
	}
	return &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}
