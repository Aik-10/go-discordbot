package discord

import "github.com/bwmarrin/discordgo"

type ButtonInteraction struct {
	InteractionLabel    string
	InteractionStyle    discordgo.ButtonStyle
	InteractionCustomID string
	InteractionEmoji    string
	Title               string
	Description         string
	Color               int
}

func CreateInteractionEmbed(data ButtonInteraction) *discordgo.MessageSend {
	Components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    data.InteractionLabel,
					Style:    data.InteractionStyle,
					CustomID: data.InteractionCustomID,
					Emoji: &discordgo.ComponentEmoji{
						Name: data.InteractionEmoji,
					},
				},
			},
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       data.Title,
		Description: data.Description,
	}

	embed.Color = data.Color

	return &discordgo.MessageSend{
		Content:    "",
		Embed:      embed,
		Components: Components,
	}

}
