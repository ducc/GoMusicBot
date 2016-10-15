package main

import (
    "fmt"
    "bytes"
    "github.com/bwmarrin/discordgo"
    "time"
    "runtime"
    "github.com/dustin/go-humanize"
    "strconv"
)

func testCommand(ctx context) {
	ctx.Reply("Testing, 1 2 3")
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
		err = chanManager.connections[channelId].connection.play(Song{"music/" + ctx.Args[1]})
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

// credit to github.com/iopred/bruxism for the stats command below

var startTime = time.Now()
var userString *string

func getDurationString(duration time.Duration) string {
    return fmt.Sprintf(
        "%0.2d:%02d:%02d",
        int(duration.Hours()),
        int(duration.Minutes())%60,
        int(duration.Seconds())%60,
    )
}

func write(buff *bytes.Buffer, str ...string) {
    for _, s := range str {
        buff.WriteString(s)
    }
}

func infoCommand(ctx context) {
    if userString == nil {
        usr, err := ctx.Discord.User(conf.OwnerId)
        if err != nil {
            fmt.Println("error getting user ", conf.OwnerId, err)
            return
        }
        str := usr.Username + "#" + usr.Discriminator
        userString = &str
    }
    stats := runtime.MemStats{}
    runtime.ReadMemStats(&stats)
    buffer := bytes.NewBufferString("```")
    write(buffer, "owner: ", *userString)
    write(buffer, "\ngo version: ", runtime.Version())
    write(buffer, "\ndiscordgo version: ", discordgo.VERSION)
    write(buffer, "\nuptime: ", getDurationString(time.Now().Sub(startTime)))
    buffer.WriteString(fmt.Sprintf("\nmemory used: %s / %s (%s garbage collected)", humanize.Bytes(stats.Alloc),
        humanize.Bytes(stats.Sys), humanize.Bytes(stats.TotalAlloc)))
    write(buffer, "\nconcurrent tasks: ", strconv.Itoa(runtime.NumGoroutine()))
    write(buffer, "\ncurrent shard: ", strconv.Itoa(ctx.Discord.ShardID))
    write(buffer, "\nshard count: ", strconv.Itoa(ctx.Discord.ShardCount))
    buffer.WriteString("```")
    ctx.Reply(buffer.String())
}