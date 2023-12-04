package main

import (
	"github.com/jonkofee/golf-club-bot/telegram"
	"github.com/jonkofee/golf-club-bot/telegram/types"
	"log/slog"
	"os"
)

func main() {
	telegramBotKey := os.Getenv(`TELEGRAM_BOT_KEY`)
	if len(telegramBotKey) == 0 {
		slog.Error(`env TELEGRAM_BOT_KEY is not set`)
	}

	client := telegram.Client{Key: telegramBotKey}

	ch := make(chan types.Update)

	fetcher := telegram.Fetcher{Client: &client}
	go fetcher.Start(ch)

	handler := telegram.Handler{Client: &client}
	handler.Start(ch)
}
