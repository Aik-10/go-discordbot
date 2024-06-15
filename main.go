package main

import (
	bot "lentokone/ticketbot/Bot"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	bot.BotToken = goDotEnvVariable("DISCORD_TOKEN")
	bot.Run()
}
