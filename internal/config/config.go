package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type configuration struct {
	BotToken             string
	BotStatus            string
	ServerGuildID        string
	TicketGategoryId     string
	BugCategoryId        string
	ManagerRoles         []string
	TicketCategory       string
	ticketOpenChannel    string
	ticketArchiveChannel string
	bugCategory          string
	bugOpenChannel       string
	bugArchiveChannel    string
}

var config *configuration

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func Load() {
	config = &configuration{
		BotToken:             goDotEnvVariable("DISCORD_TOKEN"),
		BotStatus:            goDotEnvVariable("DISCORD_BOT_STATUS"),
		ServerGuildID:        goDotEnvVariable("SERVER_GUILDID"),
		TicketGategoryId:     goDotEnvVariable("TICKET_CATEGORY"),
		BugCategoryId:        goDotEnvVariable("BUG_CATEGORY"),
		TicketCategory:       goDotEnvVariable("TICKET_CATEGORY"),
		ticketOpenChannel:    goDotEnvVariable("TICKET_OPEN_CHANNEL"),
		ticketArchiveChannel: goDotEnvVariable("TICKET_ARCHIVE_CHANNEL"),
		bugCategory:          goDotEnvVariable("BUG_CATEGORY"),
		bugOpenChannel:       goDotEnvVariable("BUG_OPEN_CHANNEL"),
		bugArchiveChannel:    goDotEnvVariable("BUG_ARCHIVE_CHANNEL"),
		ManagerRoles:         strings.Split(goDotEnvVariable("MANAGER_ROLE"), "|"),
	}
}

func BotToken() string {
	return config.BotToken
}

func BotStatus() string {
	return config.BotStatus
}

func ServerGuildID() string {
	return config.ServerGuildID
}

func GetManagerRoleIDs() []string {
	return config.ManagerRoles
}

func TicketOpenChannel() string {
	return config.ticketOpenChannel
}

func TicketArchiveChannel() string {
	return config.ticketArchiveChannel
}

func TicketCategory() string {
	return config.TicketCategory

}

func BugOpenChannel() string {
	return config.bugOpenChannel
}

func BugArchiveChannel() string {
	return config.bugArchiveChannel
}

func BugReportCategoryID() string {
	return config.BugCategoryId
}
