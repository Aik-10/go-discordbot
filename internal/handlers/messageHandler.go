package handlers

import (
	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/Aik-10/go-discordbot/internal/utils"
	"github.com/bwmarrin/discordgo"
)

func MessageInteractionHandler(CustomID string, interaction *discordgo.InteractionCreate) {
	switch CustomID {
	case "open_ticket":
		ButtonInteractHandleTicket(interaction)
	case "open_bug":
		ButtonInteractHandleBug(interaction)
	case "close_channel":
		CloseChannelButton(interaction)
	default:
		utils.Logger.Error("Unknown interaction type", "CustomID", CustomID)
	}
}

func getParentCategoryByChannelID(channelID string) (*discordgo.Channel, error) {
	channel, err := discord.Session.Channel(channelID)
	if err != nil {
		utils.Logger.Error("Failed to get channel", "error", err)
		return nil, err
	}

	if channel.ParentID != "" {
		parentCategory, err := discord.Session.Channel(channel.ParentID)
		if err != nil {
			utils.Logger.Error("Failed to get parent category", "error", err)
			return nil, err
		}

		return parentCategory, nil
	}

	return nil, nil
}

func CloseChannelButton(interaction *discordgo.InteractionCreate) {
	category, err := getParentCategoryByChannelID(interaction.ChannelID)
	if category == nil || err != nil {
		utils.Logger.Error("Failed to get parent category")
		return
	}

	switch category.ID {
	case config.TicketCategory():
		HandleChannelArchive(interaction.ChannelID, config.TicketArchiveChannel())
		HandleChannelDeletion(interaction.ChannelID)
	case config.BugReportCategoryID():
		HandleChannelArchive(interaction.ChannelID, config.BugArchiveChannel())
		HandleChannelDeletion(interaction.ChannelID)
	default:
		utils.Logger.Error("Unknown category", "category", category.ID)
	}
}

func ButtonInteractHandleBug(interaction *discordgo.InteractionCreate) {
	modal := discord.CreateInputModal(discord.ModalData{
		Title:    "Bug raport",
		CustomID: "bug_report",
		Components: []discord.ComponentModalData{
			{
				CustomID:    "bug_report_title",
				Label:       "Bug raportin aihe",
				Style:       discordgo.TextInputShort,
				Placeholder: "esim. Hahmon poisto",
				MinLength:   5,
				MaxLength:   50,
				Required:    true,
			},
			{
				CustomID:    "bug_report_reason",
				Label:       "Bug raportin syy",
				Style:       discordgo.TextInputShort,
				Placeholder: "esim. Hahmo ei poistu",
				MinLength:   5,
				MaxLength:   1500,
				Required:    true,
			},
		},
	})

	err := discord.Session.InteractionRespond(interaction.Interaction, modal)
	if err != nil {
		utils.Logger.Error("Failed to respond to interaction", "error", err)
	}
}

func ButtonInteractHandleTicket(interaction *discordgo.InteractionCreate) {
	modal := discord.CreateInputModal(discord.ModalData{
		Title:    "Avaa uusi ticket.",
		CustomID: "ticket_modal",
		Components: []discord.ComponentModalData{
			{
				CustomID:    "ticket_title",
				Label:       "Tiketin aihe",
				Style:       discordgo.TextInputShort,
				Placeholder: "esim. Hahmon poisto",
				MinLength:   5,
				MaxLength:   50,
				Required:    true,
			},
			{
				CustomID:    "ticket_reason",
				Label:       "Tiketin syy",
				Style:       discordgo.TextInputParagraph,
				Placeholder: "esim. Haluan poistaa hahmoni x, koska en pelaa sillä enää.",
				MinLength:   20,
				MaxLength:   1500,
				Required:    true,
			},
		},
	})

	err := discord.Session.InteractionRespond(interaction.Interaction, modal)
	if err != nil {
		utils.Logger.Error("Failed to respond to interaction", "error", err)
	}
}
