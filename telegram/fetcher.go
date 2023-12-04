package telegram

import (
	"github.com/jonkofee/golf-club-bot/telegram/types"
	"log/slog"
	"time"
)

type Fetcher struct {
	*Client
}

func (f Fetcher) Start(ch chan types.Update) {
	ticker := time.NewTicker(time.Second / 2)

	for {
		select {
		case <-ticker.C:
			updates, err := f.Client.GetUpdates()
			if err != nil {
				slog.Error(err.Error())
			}

			for _, update := range updates {
				ch <- update
			}
		}
	}
}
