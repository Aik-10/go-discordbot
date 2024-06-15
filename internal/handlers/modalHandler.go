package handlers

import (
	"fmt"
	"log/slog"

	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/bwmarrin/discordgo"
)

func HandleModalSubmitData(CustomID string, interaction *discordgo.InteractionCreate) {
	switch CustomID {
	case "bug_report":
		HandleBugReport(interaction)
	case "ticket_modal":
		HandleTicket(interaction)
	default:
		slog.Error("Unknown modal type")
	}
}

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

func HandleTicket(interaction *discordgo.InteractionCreate) {
	userId := interaction.Member.User.ID

	channelId := discord.CreatePrivateChannel(discord.PrivateChannelData{
		ChannelName:   "ticket-1",
		CategoryID:    config.TicketCategory(),
		CreatorUserID: userId,
		Topic:         "Ticket asdasd",
	})

	modalValues := getComponentValues(interaction.ModalSubmitData().Components)
	var reportData ModalData
	for key, value := range modalValues {
		switch key {
		case "ticket_title":
			reportData.Body = value
		case "ticket_reason":
			reportData.Title = value
		}
	}

	sendMessageToChannel(TicketData{
		Title:     reportData.Title,
		Body:      reportData.Body,
		ChannelId: channelId,
		UserId:    userId,
	}, interaction)
}

type TicketData struct {
	Title     string
	Body      string
	ChannelId string
	UserId    string
}

func sendMessageToChannel(data TicketData, interaction *discordgo.InteractionCreate) {
	message := fmt.Sprintf("<@%s> - %s\n\n__**%s**__\n\n*%s*", data.UserId, "hex", data.Title, data.Body)
	openingMessage := discord.SendChannelMessage(data.ChannelId, message)

	discord.Session.ChannelMessagePin(data.ChannelId, openingMessage.ID)

	successResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ticket created successfully!",
		},
	}

	discord.Session.InteractionRespond(interaction.Interaction, successResponse)
}

type ModalData struct {
	Body  string
	Title string
}

func HandleBugReport(interaction *discordgo.InteractionCreate) {
	userId := interaction.Member.User.ID

	channelId := discord.CreatePrivateChannel(discord.PrivateChannelData{
		ChannelName:   "bug-report",
		CategoryID:    config.BugReportCategoryID(),
		CreatorUserID: userId,
		Topic:         "Bug report channel",
	})

	modalValues := getComponentValues(interaction.ModalSubmitData().Components)
	var reportData ModalData
	for key, value := range modalValues {
		switch key {
		case "bug_report_reason":
			reportData.Body = value
		case "bug_report_title":
			reportData.Title = value
		}
	}

	sendMessageToChannel(TicketData{
		Title:     reportData.Title,
		Body:      reportData.Body,
		ChannelId: channelId,
		UserId:    userId,
	}, interaction)
}
