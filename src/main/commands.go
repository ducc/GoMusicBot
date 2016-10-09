package main

import "fmt"

func testCommand(ctx context) {
    ctx.reply("Testing, 1 2 3")
}

func joinCommand(ctx context) {
    textChannel, _ := ctx.discord.State.Channel(ctx.message.ChannelID)
    if textChannel == nil {
        fmt.Println("channel is nil")
        return
    }
    guild, _ := ctx.discord.State.Guild(textChannel.GuildID)
    if guild == nil {
        fmt.Println("guild is nil")
        return
    }
    voiceChannel := getCurrentVoiceChannel(ctx.discord, ctx.user, guild)
    if voiceChannel == nil {
        fmt.Println("voicechannel is nil")
        return
    }
    channelId := voiceChannel.ID
    if chanManager.isChannel(channelId) {
        ctx.reply("Already in that channel")
        return
    }
    chanManager.joinChannelDeafened(ctx.discord, guild.ID, channelId)
    ctx.reply("Joined channel!")
}

func leaveCommand(ctx context) {
    textChannel, _ := ctx.discord.State.Channel(ctx.message.ChannelID)
    if textChannel == nil {
        fmt.Println("channel is nil")
        return
    }
    guild, _ := ctx.discord.State.Guild(textChannel.GuildID)
    if guild == nil {
        fmt.Println("guild is nil")
        return
    }
    voiceChannel := getCurrentVoiceChannel(ctx.discord, ctx.user, guild)
    if voiceChannel == nil {
        fmt.Println("voicechannel is nil")
        return
    }
    channelId := voiceChannel.ID
    if !chanManager.isChannel(channelId) {
        ctx.reply("Not in that channel")
        return
    }
    chanManager.leaveChannel(ctx.discord, channelId)
    ctx.reply("Left channel!")
}

func playCommand(ctx context) {
    textChannel, _ := ctx.discord.State.Channel(ctx.message.ChannelID)
    if textChannel == nil {
        fmt.Println("channel is nil")
        return
    }
    guild, _ := ctx.discord.State.Guild(textChannel.GuildID)
    if guild == nil {
        fmt.Println("guild is nil")
        return
    }
    voiceChannel := getCurrentVoiceChannel(ctx.discord, ctx.user, guild)
    if voiceChannel == nil {
        fmt.Println("voicechannel is nil")
        return
    }
    channelId := voiceChannel.ID
    if !chanManager.isChannel(channelId) {
        ctx.reply("Not in that channel")
        return
    }
    var err error;
    if (len(ctx.args) == 2) {
        err = chanManager.connections[channelId].connection.play(Song{"music/" + ctx.args[1]})
    } else {
        err = chanManager.connections[channelId].connection.play(Song{"music/filthy.m4a"})
    }
    if err != nil {
        fmt.Println("error playing,", err)
    }
}

func stopCommand(ctx context) {
    textChannel, _ := ctx.discord.State.Channel(ctx.message.ChannelID)
    if textChannel == nil {
        fmt.Println("channel is nil")
        return
    }
    guild, _ := ctx.discord.State.Guild(textChannel.GuildID)
    if guild == nil {
        fmt.Println("guild is nil")
        return
    }
    voiceChannel := getCurrentVoiceChannel(ctx.discord, ctx.user, guild)
    if voiceChannel == nil {
        fmt.Println("voicechannel is nil")
        return
    }
    channelId := voiceChannel.ID
    if !chanManager.isChannel(channelId) {
        ctx.reply("Not in that channel")
        return
    }
    chanManager.connections[channelId].connection.stop()
    ctx.reply("Stopped playing!")
}