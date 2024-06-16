package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/Aik-10/go-discordbot/internal/handlers"
	"github.com/Aik-10/go-discordbot/internal/utils"
)

func Start() {
	config.Load()

	discord.InitSession()
	discord.InitConnection()

	addHandlers()

	defer discord.Session.Close()

	utils.Logger.Info("Bot is running. Press Ctrl + C to exit.")

	handlers.ReadyHandler(discord.Session)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func addHandlers() {
	discord.Session.AddHandler(handlers.InteractionHandler)
}
