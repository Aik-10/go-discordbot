package discord

import (
	"log/slog"

	"github.com/Aik-10/go-discordbot/internal/config"
	"github.com/bwmarrin/discordgo"
)

type PrivateChannelData struct {
	ChannelName   string
	CategoryID    string
	CreatorUserID string
	Topic         string
}

func CreatePrivateChannel(data PrivateChannelData) string {
	guildId := config.ServerGuildID()

	permissionOverwrites := []*discordgo.PermissionOverwrite{
		{
			ID:    data.CreatorUserID,
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages | discordgo.PermissionReadMessageHistory | discordgo.PermissionAttachFiles | discordgo.PermissionAddReactions,
		},
		{
			Deny: discordgo.PermissionViewChannel,
			ID:   guildId,
		},
	}

	// for _, roleID := range config.GetManagerRoleIDs() {
	// 	permissionOverwrites = append(permissionOverwrites, &discordgo.PermissionOverwrite{
	// 		ID:    roleID,
	// 		Allow: discordgo.PermissionViewChannel,
	// 	})
	// }

	createdChannel, err := Session.GuildChannelCreateComplex(guildId, discordgo.GuildChannelCreateData{
		Name:                 data.ChannelName,
		Type:                 discordgo.ChannelTypeGuildText,
		ParentID:             data.CategoryID,
		Topic:                data.Topic,
		PermissionOverwrites: permissionOverwrites,
	})

	if err != nil {
		slog.Error("Failed to create channel", "error", err)
	}

	return createdChannel.ID
}
