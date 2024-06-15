package discord

import "github.com/bwmarrin/discordgo"

type ButtonInteraction struct {
	Label       string
	Style       discordgo.ButtonStyle
	CustomID    string
	Emoji       string
	Title       string
	Description string
	Color       int
}

func CreateInteractionEmbed(data ButtonInteraction) *discordgo.MessageSend {
	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    data.Label,
					Style:    data.Style,
					CustomID: data.CustomID,
					Emoji: &discordgo.ComponentEmoji{
						Name: data.Emoji,
					},
				},
			},
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       data.Title,
		Description: data.Description,
		Color:       data.Color,
	}

	return &discordgo.MessageSend{
		Content:    "",
		Embed:      embed,
		Components: components,
	}
}
