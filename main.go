package main

import (
	"github.com/jonkofee/golf-club-bot/gpt"
	"github.com/jonkofee/golf-club-bot/telegram"
	"github.com/jonkofee/golf-club-bot/telegram/types"
)

func main() {
	gptClient := gpt.Create()

	telegramClient := telegram.Create()
	ch := make(chan types.Update)

	fetcher := telegram.Fetcher{Client: telegramClient}
	go fetcher.Start(ch)

	handler := telegram.Handler{Client: telegramClient, Gpt: gptClient}
	handler.Start(ch)
}
