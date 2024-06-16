package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/Aik-10/go-discordbot/internal/utils"
)

func HandleChannelArchive(channelID string, archiveChannelID string) {
	messages, err := discord.Session.ChannelMessages(channelID, 100, "", "", "")
	if err != nil {
		utils.Logger.Error("Failed to get messages", "error", err)
	}

	file, err := os.Create("tmp/" + channelID + ".txt")
	if err != nil {
		utils.Logger.Error("Failed to create file", "error", err)
		return
	}
	defer file.Close()

	for i := len(messages)/2 - 1; i >= 0; i-- {
		opp := len(messages) - 1 - i
		messages[i], messages[opp] = messages[opp], messages[i]
	}

	for _, message := range messages {
		logMessage := fmt.Sprintf("%s - %s <@%s>: %s\n", message.Timestamp.UTC(), message.Author.Username, message.Author.ID, message.Content)

		if message.Author.Bot {
			logMessage = fmt.Sprintf("%s\n", message.Content)
		}

		utils.Logger.Info(logMessage)

		if _, err := file.WriteString(logMessage); err != nil {
			utils.Logger.Error("Failed to write message to file", "error", err)
			return
		}
	}

	channel, err := discord.Session.Channel(channelID)
	if err != nil {
		utils.Logger.Error("Failed to get channel information", "error", err)
		return
	}

	channelInfo := fmt.Sprintf("\n\nChannel: %s\nClosed at: %s\n", channel.Name, time.Now().UTC())
	if _, err := file.WriteString(channelInfo); err != nil {
		utils.Logger.Error("Failed to write message to file", "error", err)
		return
	}

	utils.Logger.Info("All messages have been archived to file.")

	archiveMessageContent := fmt.Sprintf("<@%s> - %s", "228494142236393472", archiveChannelID)
	SendChannelFileToArchive(channelID, archiveChannelID, archiveMessageContent)
}

func SendChannelFileToArchive(channelID string, archiveChannelID string, content string) {
	file, err := os.Open("tmp/" + channelID + ".txt")
	if err != nil {
		utils.Logger.Error("Failed to open file", "error", err)
		return
	}
	defer file.Close()

	message, err := discord.Session.ChannelFileSendWithMessage(archiveChannelID, content, "archive.txt", file)
	if err != nil {
		utils.Logger.Error("Failed to send file", "error", err)
		return
	}

	utils.Logger.Info("File has been sent to archive channel", "messageId", message.ID)

	err = file.Close()
	if err != nil {
		utils.Logger.Error("Failed to close file", "error", err)
		return
	}
}

func HandleChannelDeletion(channelID string) {
	body, err := discord.Session.ChannelDelete(channelID)
	if err != nil {
		utils.Logger.Error("Failed to delete channel", "error", err)
		return
	}

	utils.Logger.Info("Channel has been deleted", "name", body.Name)

	err = os.Remove("tmp/" + channelID + ".txt")
	if err != nil {
		utils.Logger.Error("Failed to delete file", "error", err)
	}
}
