package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var BotToken string
var ServerGuildID string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// discord.AddHandler(newMessage)
	discord.AddHandler(newInteraction)

	discord.Open()
	defer discord.Close() // close session, after function termination

	fmt.Println("Logged in as", discord.State.User.Username)

	doSendOpenTicketMessagesToChannel(discord)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

type Data struct {
	Title       string
	Description string
	Color       int
	Components  []discordgo.MessageComponent
}

func doClearChannelMessagesByBot(discord *discordgo.Session, channelID string) {
	messages, err := discord.ChannelMessages(channelID, 100, "", "", "")
	checkNilErr(err)

	for _, message := range messages {
		if (message.Author.ID == discord.State.User.ID) || (message.Author.Bot) {
			err = discord.ChannelMessageDelete(channelID, message.ID)
			fmt.Println("Deleted message", message.ID, "from channel", channelID)

			checkNilErr(err)
		}
	}
}

func doSendOpenTicketMessagesToChannel(discord *discordgo.Session) {
	ticketOpenChannel := "1251494812876673064"
	bugOpenChannel := "1251494904471748618"

	TicketComponents := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    "Open Ticket!",
					Style:    discordgo.PrimaryButton,
					CustomID: "open_ticket",
					Emoji: &discordgo.ComponentEmoji{
						Name: "📩",
					},
				},
			},
		},
	}

	BugComponents := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    "Open Bug raport!",
					Style:    discordgo.PrimaryButton,
					CustomID: "open_bug",
					Emoji: &discordgo.ComponentEmoji{
						Name: "🚀",
					},
				},
			},
		},
	}

	doClearChannelMessagesByBot(discord, ticketOpenChannel)
	doClearChannelMessagesByBot(discord, bugOpenChannel)

	doSendMessageToChannel(discord, ticketOpenChannel, Data{
		Title:       "New Ticket Report",
		Description: "A new ticket has been submitted",
		Components:  TicketComponents,
		Color:       0x00ff00,
	})

	doSendMessageToChannel(discord, bugOpenChannel, Data{
		Title:       "New Bug Report",
		Description: "A new bug report has been submitted",
		Components:  BugComponents,
		Color:       0x00ff00,
	})
}

func doSendMessageToChannel(discord *discordgo.Session, channelID string, data Data) {
	embed := &discordgo.MessageEmbed{
		Title:       data.Title,
		Description: data.Description,
	}

	embed.Color = data.Color

	_, err := discord.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content:    "",
		Embed:      embed,
		Components: data.Components,
	})
	checkNilErr(err)
}

func isChannelInCategory(discord *discordgo.Session, channelID string, categoryIDs []string) bool {
	channel, err := discord.Channel(channelID)
	checkNilErr(err)

	for _, categoryID := range categoryIDs {
		if channel.ParentID == categoryID {
			return true
		}
	}

	return false
}

func getGuildChannelsFromCategoryId(discord *discordgo.Session, categoryID string) []*discordgo.Channel {
	channels, err := discord.GuildChannels(ServerGuildID)
	checkNilErr(err)

	var categoryChannels []*discordgo.Channel
	for _, channel := range channels {
		if channel.ParentID == categoryID {
			categoryChannels = append(categoryChannels, channel)
		}
	}

	return categoryChannels
}

func doCreateNewPrivateChannelToCategory(discord *discordgo.Session, categoryID string, channelName string, interaction *discordgo.InteractionCreate) {

	permissionOverwrites := []*discordgo.PermissionOverwrite{
		{
			ID:    interaction.Member.User.ID,
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages | discordgo.PermissionReadMessageHistory | discordgo.PermissionAttachFiles | discordgo.PermissionAddReactions,
		},
		{
			Deny: discordgo.PermissionViewChannel,
			ID:   ServerGuildID,
		},
	}

	for _, roleID := range adminRoles {
		permissionOverwrites = append(permissionOverwrites, &discordgo.PermissionOverwrite{
			ID:    roleID,
			Allow: discordgo.PermissionViewChannel,
		})
	}

	createdChannel, err := discord.GuildChannelCreateComplex(ServerGuildID, discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildText,
		ParentID:             categoryID,
		Topic:                "Private channel for communication with the server staff.",
		PermissionOverwrites: permissionOverwrites,
	})
	checkNilErr(err)

	doSendMessageToPrivateChannelAfterCreation(discord, createdChannel.ID, interaction)
}

func doSendMessageToPrivateChannelAfterCreation(discord *discordgo.Session, channelID string, interaction *discordgo.InteractionCreate) {
	components := interaction.ModalSubmitData().Components
	var ticketTitle, ticketReason string

	for _, component := range components {
		fmt.Println(component)
	}

	fmt.Println(components)
	fmt.Println(ticketTitle, ticketReason)

	openingMessage, err := discord.ChannelMessageSend(channelID, fmt.Sprintf("<@%s> - %s\n\n__**%s**__\n\n*%s*", interaction.Member.User.ID, "hex", ticketTitle, ticketReason))
	checkNilErr(err)

	err = discord.ChannelMessagePin(channelID, openingMessage.ID)
	checkNilErr(err)

	successResponse := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ticket created successfully!",
		},
	}

	err = discord.InteractionRespond(interaction.Interaction, successResponse)
	checkNilErr(err)
}

/* func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.GuildID != ServerGuildID {
		return
	}

	categoryIDs := []string{
		"906482313624444988",
		"1183812513125638306",
	}

	fmt.Println(isChannelInCategory(discord, message.ChannelID, categoryIDs))

	categoryChannels := getGuildChannelsFromCategoryId(discord, "906482313624444988")

	for _, channel := range categoryChannels {
		fmt.Println(channel.ID, channel.Name)
	}

	fmt.Println(message.Author.ID, message.Author.Username)
	fmt.Println(message.ChannelID, message.GuildID, message.Message.Attachments, message.Message.Components, message.Message.Embeds, message.Message.MentionEveryone, message.Message.MentionRoles, message.Message.Mentions)
} */

func newInteraction(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.GuildID != ServerGuildID {
		return
	}

	fmt.Println("New interaction ", interaction.Member.User.ID, "(", interaction.Member.User.Username, ")")

	switch interaction.Type {
	case discordgo.InteractionMessageComponent:
		// Interaction with a message component (button, select menu, etc.)
		componentData := interaction.MessageComponentData()
		fmt.Printf("Component Interaction: customId = %s, values = %v\n", componentData.CustomID, componentData.Values)

		openInteractionModal(discord, interaction, componentData)
	case discordgo.InteractionApplicationCommand:
		commandData := interaction.ApplicationCommandData()
		fmt.Printf("Application Command Interaction: commandName = %s, options = %v\n", commandData.Name, commandData.Options)
	case discordgo.InteractionModalSubmit:
		// Interaction with a modal submission
		modalData := interaction.ModalSubmitData()
		fmt.Printf("Modal Submit Interaction: customId = %s, components = %v\n", modalData.CustomID, modalData.Components)

		doCreateNewPrivateChannelToCategory(discord, "906482313624444988", "testi-ticket", interaction)

		// Process the submitted data
		// processModalSubmission(discord, interaction, modalData)
	default:
		fmt.Println("Unknown interaction type")
	}
}

func openInteractionModal(discord *discordgo.Session, interaction *discordgo.InteractionCreate, componentData discordgo.MessageComponentInteractionData) {
	// componentData.CustomID

	modal := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "ticket-modal",
			Title:    "Avaa uusi ticket.",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.TextInput{
							CustomID:    "ticket_title",
							Label:       "Tiketin aihe",
							Style:       discordgo.TextInputShort,
							Placeholder: "esim. Hahmon poisto",
							MinLength:   5,
							MaxLength:   50,
							Required:    true,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.TextInput{
							CustomID:    "ticket_reason",
							Label:       "Tiketin syy",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "esim. Haluan poistaa hahmoni x, koska en pelaa sillä enää.",
							MinLength:   20,
							MaxLength:   1500,
							Required:    true,
						},
					},
				},
			},
		},
	}

	err := discord.InteractionRespond(interaction.Interaction, modal)
	checkNilErr(err)
}
