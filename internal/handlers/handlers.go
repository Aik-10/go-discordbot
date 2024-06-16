package handlers

import (
	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/discord"
	"github.com/Aik-10/go-discordbot/internal/utils"
	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(s *discordgo.Session) {
	err := s.UpdateGameStatus(1, config.BotStatus())
	if err != nil {
		utils.Logger.Warn("failed to update game status", "error", err)
	}

	/* after client is ready clear channels */
	discord.DeleteBotMessagesInChannel(config.TicketOpenChannel())
	discord.DeleteBotMessagesInChannel(config.BugOpenChannel())

	ticketEmbed := discord.CreateInteractionEmbed(discord.ButtonInteraction{
		Label:       "Open Ticket!",
		Style:       discordgo.PrimaryButton,
		CustomID:    "open_ticket",
		Emoji:       "📩",
		Title:       "New Ticket",
		Description: "Open a new ticket",
		Color:       0x00ff00,
	})

	bugEmbed := discord.CreateInteractionEmbed(discord.ButtonInteraction{
		Label:       "Open Bug!",
		Style:       discordgo.PrimaryButton,
		CustomID:    "open_bug",
		Emoji:       "📩",
		Title:       "Bug raport",
		Description: "Open a new bug raport",
		Color:       0x00ff00,
	})

	discord.SendChannelMessageSendComplex(config.TicketOpenChannel(), ticketEmbed)
	discord.SendChannelMessageSendComplex(config.BugOpenChannel(), bugEmbed)
}

func InteractionHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.GuildID != config.ServerGuildID() {
		return
	}

	utils.Logger.Info("Interaction detected", "interaction", interaction, "userId", interaction.Member.User.ID, "username", interaction.Member.User.Username)

	switch interaction.Type {
	case discordgo.InteractionMessageComponent:
		componentData := interaction.MessageComponentData()
		utils.Logger.Info("Component Interaction", "customId", componentData.CustomID, "values", componentData.Values)

		MessageInteractionHandler(componentData.CustomID, interaction)
		// openInteractionModal(discord, interaction, componentData)
	case discordgo.InteractionApplicationCommand:
		commandData := interaction.ApplicationCommandData()
		utils.Logger.Info("Application Command Interaction", "commandName", commandData.Name, "options", commandData.Options)
	case discordgo.InteractionModalSubmit:
		// Interaction with a modal submission
		modalData := interaction.ModalSubmitData()
		utils.Logger.Info("Modal Submit Interaction", "customId", modalData.CustomID, "components", modalData.Components)
		HandleModalSubmitData(modalData.CustomID, interaction)
		// doCreateNewPrivateChannelToCategory(discord, "906482313624444988", "testi-ticket", interaction)

	default:
		utils.Logger.Error("Unknown interaction type")
	}
}
