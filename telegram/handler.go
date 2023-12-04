package telegram

import (
	"github.com/jonkofee/golf-club-bot/telegram/types"
	"log/slog"
)

type Handler struct {
	*Client
}

func (h Handler) Start(ch chan types.Update) {
	for update := range ch {
		h.handle(update)
	}
}

func (h Handler) handle(data types.Update) {
	slog.Info(`Handle update`, data)

	switch {
	case data.Message.NewChatMember != nil:
		h.handleNewChatMember(data)
		return
	}
}

func (h Handler) handleNewChatMember(data types.Update) {
	text := `Привет, добро пожаловать!) Скидывай свою лайбу`

	h.Client.SendMessage(data.Message.Chat.Id, text, data.Message.Id)
	h.Client.SendSticker(data.Message.Chat.Id, `CAACAgQAAx0CfCksJQADEmVuQthGd2cCFH5DYGcFxqR6EkxbAAI-AQACqCEhBrJy8YE-YrIMMwQ`, 0)
}
