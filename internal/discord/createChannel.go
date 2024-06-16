package discord

import (
	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/Aik-10/go-discordbot/internal/utils"
	"github.com/bwmarrin/discordgo"
)

// PrivateChannelData holds data for creating a private channel
type PrivateChannelData struct {
	ChannelName   string
	CategoryID    string
	CreatorUserID string
	Topic         string
}

// CreatePrivateChannel creates a new private channel with specified permissions
func CreatePrivateChannel(data PrivateChannelData) (string, error) {
	guildID := config.ServerGuildID()

	permissionOverwrites := []*discordgo.PermissionOverwrite{
		{
			ID:    data.CreatorUserID,
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages | discordgo.PermissionReadMessageHistory | discordgo.PermissionAttachFiles | discordgo.PermissionAddReactions,
		},
		{
			ID:   guildID,
			Deny: discordgo.PermissionViewChannel,
		},
	}

	/* for _, roleID := range config.GetManagerRoleIDs() {
		permissionOverwrites = append(permissionOverwrites, &discordgo.PermissionOverwrite{
			ID:    roleID,
			Allow: discordgo.PermissionViewChannel,
		})
	} */

	channelData := discordgo.GuildChannelCreateData{
		Name:                 data.ChannelName,
		Type:                 discordgo.ChannelTypeGuildText,
		ParentID:             data.CategoryID,
		Topic:                data.Topic,
		PermissionOverwrites: permissionOverwrites,
	}

	createdChannel, err := Session.GuildChannelCreateComplex(guildID, channelData)
	if err != nil {
		utils.Logger.Error("Failed to create channel", "error", err)
		return "", err
	}

	return createdChannel.ID, nil
}
