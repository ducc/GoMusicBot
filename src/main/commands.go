package main

import (
	"bytes"
	"fmt"
	"os"
)

func helpCommand(ctx context) {
	cmds := cmdManager.commands
	buffer := bytes.NewBufferString("Commands: ")
	for key := range cmds {
		buffer.WriteString(key)
		buffer.WriteString(", ")
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}

func joinCommand(ctx context) {
	textChannel, _ := ctx.Discord.State.Channel(ctx.Message.ChannelID)
	if textChannel == nil {
		fmt.Println("channel is nil")
		return
	}
	guild, _ := ctx.Discord.State.Guild(textChannel.GuildID)
	if guild == nil {
		fmt.Println("guild is nil")
		return
	}
	voiceChannel := getCurrentVoiceChannel(ctx.Discord, ctx.User, guild)
	if voiceChannel == nil {
		fmt.Println("voicechannel is nil")
		return
	}
	channelId := voiceChannel.ID
	if chanManager.isChannel(channelId) {
		ctx.Reply("Already in that channel")
		return
	}
	chanManager.joinChannelDeafened(ctx.Discord, guild.ID, channelId)
	ctx.Reply("Joined channel!")
}

func leaveCommand(ctx context) {
	textChannel, _ := ctx.Discord.State.Channel(ctx.Message.ChannelID)
	if textChannel == nil {
		fmt.Println("channel is nil")
		return
	}
	guild, _ := ctx.Discord.State.Guild(textChannel.GuildID)
	if guild == nil {
		fmt.Println("guild is nil")
		return
	}
	voiceChannel := getCurrentVoiceChannel(ctx.Discord, ctx.User, guild)
	if voiceChannel == nil {
		fmt.Println("voicechannel is nil")
		return
	}
	channelId := voiceChannel.ID
	if !chanManager.isChannel(channelId) {
		ctx.Reply("Not in that channel")
		return
	}
	chanManager.leaveChannel(ctx.Discord, channelId)
	ctx.Reply("Left channel!")
}

func playCommand(ctx context) {
	textChannel, _ := ctx.Discord.State.Channel(ctx.Message.ChannelID)
	if textChannel == nil {
		fmt.Println("channel is nil")
		return
	}
	guild, _ := ctx.Discord.State.Guild(textChannel.GuildID)
	if guild == nil {
		fmt.Println("guild is nil")
		return
	}
	voiceChannel := getCurrentVoiceChannel(ctx.Discord, ctx.User, guild)
	if voiceChannel == nil {
		fmt.Println("voicechannel is nil")
		return
	}
	channelId := voiceChannel.ID
	if !chanManager.isChannel(channelId) {
		ctx.Reply("Not in that channel")
		return
	}
	var err error
	if len(ctx.Args) == 2 {
        url := getYoutubeUrl(ctx.Args[1])
        song := Song{url}
        err = chanManager.connections[channelId].connection.play(song)
	} else {
	    err = chanManager.connections[channelId].connection.play(Song{"music/filthy.m4a"})
	}
	if err != nil {
		fmt.Println("error playing,", err)
	}
}

func stopCommand(ctx context) {
	textChannel, _ := ctx.Discord.State.Channel(ctx.Message.ChannelID)
	if textChannel == nil {
		fmt.Println("channel is nil")
		return
	}
	guild, _ := ctx.Discord.State.Guild(textChannel.GuildID)
	if guild == nil {
		fmt.Println("guild is nil")
		return
	}
	voiceChannel := getCurrentVoiceChannel(ctx.Discord, ctx.User, guild)
	if voiceChannel == nil {
		fmt.Println("voicechannel is nil")
		return
	}
	channelId := voiceChannel.ID
	if !chanManager.isChannel(channelId) {
		ctx.Reply("Not in that channel")
		return
	}
	chanManager.connections[channelId].connection.stop()
	ctx.Reply("Stopped playing!")
}

func stopBotCommand(ctx context) {
	if ctx.User.ID != conf.OwnerId {
		return
	}
	ctx.Reply("Bye :)")
	ctx.Discord.Close()
	os.Exit(-1)
}
