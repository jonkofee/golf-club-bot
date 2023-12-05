package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

const url = `https://api.openai.com/v1/chat/completions`

type Client struct {
	Key string
}

func Create() *Client {
	key := os.Getenv(`OPENAI_API_KEY`)
	if len(key) == 0 {
		slog.Error(`env OPENAI_API_KEY is not set`)
	}

	return &Client{Key: key}
}

func (c Client) GetResponse(request string) (response string, err error) {
	respData := &struct {
		Choices []struct {
			Message struct {
				Content string
			}
		}
	}{}

	bodyData := map[string]interface{}{
		`model`: `gpt-3.5-turbo`,
		`messages`: [1]map[string]string{{
			`role`:    `user`,
			`content`: request,
		}},
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(bodyData)

	req, _ := http.NewRequest(`POST`, url, buffer)
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, c.Key))
	req.Header.Set(`Content-Type`, `application/json`)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	if err = json.Unmarshal(bodyBytes, respData); err != nil {
		return
	}

	if len(respData.Choices) == 0 {
		return "", errors.New(`empty Choices`)
	}

	response = respData.Choices[0].Message.Content

	return
}
