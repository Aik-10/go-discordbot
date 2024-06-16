package handlers

import (
	"fmt"
	"log/slog"

	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/bwmarrin/discordgo"
)

// HandleModalSubmitData routes the modal submission based on CustomID
func HandleModalSubmitData(CustomID string, interaction *discordgo.InteractionCreate) {
	switch CustomID {
	case "bug_report":
		handleBugReport(interaction)
	case "ticket_modal":
		handleTicket(interaction)
	default:
		slog.Error("Unknown modal type")
	}
}

// getComponentValues extracts values from modal components
func getComponentValues(components []discordgo.MessageComponent) map[string]string {
	values := make(map[string]string)
	for _, comp := range components {
		if row, ok := comp.(*discordgo.ActionsRow); ok {
			for _, component := range row.Components {
				if input, ok := component.(*discordgo.TextInput); ok {
					values[input.CustomID] = input.Value
				}
			}
		}
	}
	return values
}

func handleTicket(interaction *discordgo.InteractionCreate) {
	handleInteraction(interaction, "ticket-1", config.TicketCategory(), "ticket_title", "ticket_reason")
}

func handleBugReport(interaction *discordgo.InteractionCreate) {
	handleInteraction(interaction, "bug-report", config.BugReportCategoryID(), "bug_report_title", "bug_report_reason")
}

func handleInteraction(interaction *discordgo.InteractionCreate, channelName, categoryID, titleKey, bodyKey string) {
	userId := interaction.Member.User.ID

	channelId, err := discord.CreatePrivateChannel(discord.PrivateChannelData{
		ChannelName:   channelName,
		CategoryID:    categoryID,
		CreatorUserID: userId,
		Topic:         channelName + " channel",
	})
	if err != nil {
		slog.Error("Failed to create channel", "error", err)
		respondWithError(interaction, "Failed to create channel. Please try again later.")
		return
	}

	modalValues := getComponentValues(interaction.ModalSubmitData().Components)
	reportData := ModalData{
		Title: modalValues[titleKey],
		Body:  modalValues[bodyKey],
	}

	sendMessageToChannel(TicketData{
		Title:     reportData.Title,
		Body:      reportData.Body,
		ChannelId: channelId,
		UserId:    userId,
	}, interaction)
}

func sendMessageToChannel(data TicketData, interaction *discordgo.InteractionCreate) {
	message := fmt.Sprintf("<@%s> - %s\n\n__**%s**__\n\n*%s*", data.UserId, "hex", data.Title, data.Body)
	openingMessage, err := discord.SendChannelMessage(data.ChannelId, message)
	if err != nil {
		slog.Error("Failed to send channel message", "error", err)
		respondWithError(interaction, "Failed to send message to channel. Please try again later.")
		return
	}

	err = discord.Session.ChannelMessagePin(data.ChannelId, openingMessage.ID)
	if err != nil {
		slog.Error("Failed to pin message", "error", err)
		respondWithError(interaction, "Failed to pin message in channel. Please try again later.")
		return
	}

	successResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ticket created successfully!",
		},
	}

	err = discord.Session.InteractionRespond(interaction.Interaction, successResponse)
	if err != nil {
		slog.Error("Failed to send interaction response", "error", err)
	}
}

// respondWithError sends an error message as an interaction response
func respondWithError(interaction *discordgo.InteractionCreate, message string) {
	errorResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	}
	err := discord.Session.InteractionRespond(interaction.Interaction, errorResponse)
	if err != nil {
		slog.Error("Failed to send error interaction response", "error", err)
	}
}

// ModalData holds data from modal submissions
type ModalData struct {
	Body  string
	Title string
}

// TicketData holds data for sending messages to channels
type TicketData struct {
	Title     string
	Body      string
	ChannelId string
	UserId    string
}
