package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonkofee/golf-club-bot/telegram/types"
	"io"
	"log"
	"log/slog"
	"maps"
	"net/http"
	"os"
)

const url = `https://api.telegram.org/bot%s/%s`

var lastReceivedUpdateId int64

type Client struct {
	Key string
}

func Create() *Client {
	key := os.Getenv(`TELEGRAM_BOT_KEY`)
	if len(key) == 0 {
		slog.Error(`env TELEGRAM_BOT_KEY is not set`)
	}

	return &Client{Key: key}
}

func (c Client) SendMessage(chatId int64, text string, replyToMessageId int64) error {
	return c.request(`sendMessage`, map[string]interface{}{
		`chat_id`:             chatId,
		`text`:                text,
		`reply_to_message_id`: replyToMessageId,
	})
}

func (c Client) SendSticker(chatId int64, sticker string, replyToMessageId int64) error {
	return c.request(`sendSticker`, map[string]interface{}{
		`chat_id`:             chatId,
		`sticker`:             sticker,
		`reply_to_message_id`: replyToMessageId,
	})
}

func (c Client) request(method string, data map[string]interface{}) error {
	bodyData := map[string]interface{}{}

	maps.Copy(bodyData, data)

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(bodyData)

	resp, err := http.Post(fmt.Sprintf(url, c.Key, method), `application/json`, buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c Client) GetUpdates() (result []types.Update, err error) {
	respData := &struct {
		Result []types.Update
	}{}
	reqUrl := fmt.Sprintf(url, c.Key, `getUpdates`)

	if lastReceivedUpdateId > 0 {
		reqUrl += fmt.Sprintf(`?offset=%d`, lastReceivedUpdateId+1)
	}

	resp, err := http.Get(reqUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bodyBytes, respData); err != nil {
		log.Fatalf("Parse response failed, reason: %v \n", err)
	}

	for _, update := range respData.Result {
		if update.Id > lastReceivedUpdateId {
			lastReceivedUpdateId = update.Id
		}
	}

	return respData.Result, nil
}
