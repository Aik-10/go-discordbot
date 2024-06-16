package discord

import (
	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/utils"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func InitSession() {
	var err error
	Session, err = discordgo.New("Bot " + config.BotToken()) // Initializing discord session
	if err != nil {
		utils.Logger.Error("failed to create discord session", "error", err)
	}

	Session.Identify.Intents = discordgo.IntentsAll
}

func InitConnection() {
	if err := Session.Open(); err != nil { // Creating a connection
		utils.Logger.Error("failed to create websocket connection to discord", "error", err)
		return
	}
}
