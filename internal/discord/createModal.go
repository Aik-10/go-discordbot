package discord

import (
	"github.com/bwmarrin/discordgo"
)

type ComponentModalData struct {
	CustomID    string
	Label       string
	Style       discordgo.TextInputStyle
	Placeholder string
	MinLength   int
	MaxLength   int
	Required    bool
}

type ModalData struct {
	Title      string
	CustomID   string
	Components []ComponentModalData
}

func CreateInputModal(data ModalData) *discordgo.InteractionResponse {
	components := make([]discordgo.MessageComponent, 0, len(data.Components))

	for _, component := range data.Components {
		textInput := &discordgo.TextInput{
			CustomID:    component.CustomID,
			Label:       component.Label,
			Style:       component.Style,
			Placeholder: component.Placeholder,
			MinLength:   component.MinLength,
			MaxLength:   component.MaxLength,
			Required:    component.Required,
		}

		components = append(components, discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{textInput},
		})
	}

	modal := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID:   data.CustomID,
			Title:      data.Title,
			Components: components,
		},
	}

	return modal
}
