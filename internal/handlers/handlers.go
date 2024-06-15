package handlers

import (
	"log/slog"

	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	err := s.UpdateGameStatus(1, config.BotStatus())
	if err != nil {
		slog.Warn("failed to update game status", "error", err)
	}

	/* after client is ready clear channels */
	discord.DeleteBotMessagesInChannel(config.TicketOpenChannel())
	discord.DeleteBotMessagesInChannel(config.BugOpenChannel())

	ticketEmbed := discord.CreateInteractionEmbed(discord.ButtonInteraction{
		InteractionLabel:    "Open Ticket!",
		InteractionStyle:    discordgo.PrimaryButton,
		InteractionCustomID: "open_ticket",
		InteractionEmoji:    "📩",
		Title:               "New Ticket",
		Description:         "Open a new ticket",
		Color:               0x00ff00,
	})

	bugEmbed := discord.CreateInteractionEmbed(discord.ButtonInteraction{
		InteractionLabel:    "Open Bug!",
		InteractionStyle:    discordgo.PrimaryButton,
		InteractionCustomID: "open_bug",
		InteractionEmoji:    "📩",
		Title:               "Bug raport",
		Description:         "Open a new bug raport",
		Color:               0x00ff00,
	})

	discord.SendChannelMessageSendComplex(config.TicketOpenChannel(), ticketEmbed)
	discord.SendChannelMessageSendComplex(config.BugOpenChannel(), bugEmbed)
}

func InteractionHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.GuildID != config.ServerGuildID() {
		return
	}

	slog.Info("Interaction detected", "interaction", interaction, "userId", interaction.Member.User.ID, "username", interaction.Member.User.Username)

	switch interaction.Type {
	case discordgo.InteractionMessageComponent:
		componentData := interaction.MessageComponentData()
		slog.Info("Component Interaction", "customId", componentData.CustomID, "values", componentData.Values)

		// openInteractionModal(discord, interaction, componentData)
	case discordgo.InteractionApplicationCommand:
		commandData := interaction.ApplicationCommandData()
		slog.Info("Application Command Interaction", "commandName", commandData.Name, "options", commandData.Options)
	case discordgo.InteractionModalSubmit:
		// Interaction with a modal submission
		modalData := interaction.ModalSubmitData()
		slog.Info("Modal Submit Interaction", "customId", modalData.CustomID, "components", modalData.Components)

		// doCreateNewPrivateChannelToCategory(discord, "906482313624444988", "testi-ticket", interaction)

	default:
		slog.Error("Unknown interaction type")
	}

}
