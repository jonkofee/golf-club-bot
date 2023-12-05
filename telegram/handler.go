package telegram

import (
	"github.com/jonkofee/golf-club-bot/gpt"
	"github.com/jonkofee/golf-club-bot/telegram/types"
	"log/slog"
)

type Handler struct {
	*Client
	Gpt *gpt.Client
}

func (h Handler) Start(ch chan types.Update) {
	for update := range ch {
		h.handle(update)
	}
}

func (h Handler) handle(data types.Update) {
	slog.Info(`Handle update`, data)

	if data.Message == nil {
		return
	}

	switch {
	case data.Message.NewChatMember != nil:
		h.handleNewChatMember(data)
		return
	}
}

func (h Handler) handleNewChatMember(data types.Update) {
	text, _ := h.Gpt.GetResponse(`Представим, что ты встречаешь новопришедшего участника чата. Сперва поприветствуй его, попроси представиться и рассказать о себе. Далее ознакомь его с нашими правилами, их обязательно нужно соблюдать. Уточни, что если правила не соблюдать, то получишь бан! Маты разрешены. Вот они: нельзя оскорблять собеседника, нельзя затрагивать политические темы, ты фанат Golf 8 или соплатформенного автомобиля. В конце НАСТОЙЧИВО попроси отправить фотографию своей машины. Весь текст пиши в неформальном стиле, но без оскорблений. Ты не ассистен, ты живой человек и зовут тебя Матулазем`)

	h.Client.SendMessage(data.Message.Chat.Id, text, data.Message.Id)
	h.Client.SendSticker(data.Message.Chat.Id, `CAACAgQAAx0CfCksJQADEmVuQthGd2cCFH5DYGcFxqR6EkxbAAI-AQACqCEhBrJy8YE-YrIMMwQ`, 0)
}
