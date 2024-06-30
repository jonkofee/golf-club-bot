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
		err := h.handle(update)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func (h Handler) handle(data types.Update) (err error) {
	slog.Info(`Handle update`, data)

	if data.Message == nil {
		return
	}

	switch {
	case data.Message.NewChatMember != nil:
		err = h.handleNewChatMember(data)
		return
	}

	return
}

func (h Handler) handleNewChatMember(data types.Update) (err error) {
	text, err := h.Gpt.GetResponse(`Зовут тебя Матузалем. Весь текст пиши в неформальном стиле, но без оскорблений. Представим, что ты встречаешь новопришедшего участника чата . Сперва поприветствуй его, попроси представиться и рассказать о себе. Далее ознакомь его с нашими правилами, их обязательно нужно соблюдать. Уточни, что если правила не соблюдать, то получишь бан! Вот они: нельзя оскорблять собеседника, нельзя затрагивать политические темы, ты фанат Golf 8 или соплатформенного автомобиля. Маты разрешены. В конце НАСТОЙЧИВО попроси отправить фотографию своей машины.`)
	if err != nil {
		return err
	}

	err = h.Client.SendMessage(data.Message.Chat.Id, text, data.Message.Id)
	if err != nil {
		return err
	}

	err = h.Client.SendVoice(data.Message.Chat.Id, `AwACAgIAAxkDAAMcZoEtY6rkU2AqCho1HSm4BWHSCH8AAgFLAALK9AlIK5iRMfGoXso1BA`)
	if err != nil {
		return err
	}

	err = h.Client.SendSticker(data.Message.Chat.Id, `CAACAgQAAx0CfCksJQADEmVuQthGd2cCFH5DYGcFxqR6EkxbAAI-AQACqCEhBrJy8YE-YrIMMwQ`)
	if err != nil {
		return err
	}

	return nil
}
