package main

import (
	"github.com/jonkofee/golf-club-bot/telegram"
	"github.com/jonkofee/golf-club-bot/telegram/types"
)

func main() {
	client := telegram.Create()
	ch := make(chan types.Update)

	fetcher := telegram.Fetcher{Client: client}
	go fetcher.Start(ch)

	handler := telegram.Handler{Client: client}
	handler.Start(ch)
}
