package main

import (
	"log"
	"time"

	"github.com/Aik-10/go-discordbot/internal/bot"
	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/utils"
	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.GetSentryDNS(),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	utils.InitLogger()

	bot.Start()
}
