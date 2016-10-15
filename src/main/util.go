package main

import (
	"github.com/bwmarrin/discordgo"
)

// credit: b1nzy airhorn.solutions
func getCurrentVoiceChannel(discord *discordgo.Session, user *discordgo.User, guild *discordgo.Guild) *discordgo.Channel {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == user.ID {
			channel, _ := discord.State.Channel(vs.ChannelID)
			return channel
		}
	}
	return nil
}
