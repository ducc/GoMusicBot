package oldmain

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		ctx.Reply("Could not find your current voice channel. Are you in a voice channel that the bot can join?")
		return
	}
	channelId := voiceChannel.ID
	if chanManager.isChannel(channelId) {
		ctx.Reply("The bot is already in that voice channel! To make the bot leave, use `music leave`!")
		return
	}
	chanManager.joinChannelDeafened(ctx.Discord, guild.ID, channelId)
	ctx.Reply("Joined the <#" + channelId + "> voice channel!")
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
		ctx.Reply("Could not find your current voice channel. Are you in a voice channel that the bot can join?")
		return
	}
	channelId := voiceChannel.ID
	if !chanManager.isChannel(channelId) {
		ctx.Reply("Not currently in the voice channel <#" + channelId + ">. For the bot to join, use `music join`!")
		return
	}
	if len(ctx.Args) != 2 {
		ctx.Reply("Usage: music play <yt video id>")
		return
	}
	con := chanManager.connections[channelId]
	if con.connection.playing {
		ctx.Reply("Already playing a song in <#" + channelId + ">! To stop playing, use `music stop`. " +
			"To queue a song, use `music queue <song>`.")
		return
	}
	videoId := ctx.Args[1]
	msg := ctx.Reply("Getting song info for `" + videoId + "`...")
	vResult, err := getYoutubeUrl(videoId)
	if err != nil {
		ctx.Discord.ChannelMessageEdit(textChannel.ID, msg.ID, "Something went wrong! Try a different song.")
		return
	}
	song := Song{vResult.media, vResult.title}
	ctx.Discord.ChannelMessageEdit(textChannel.ID, msg.ID, "Now playing **"+song.name+"** - <https://youtu.be/"+
		videoId+">!")
	err = con.connection.play(song)
	if err != nil {
		ctx.Discord.ChannelMessageEdit(textChannel.ID, msg.ID, "Something went wrong! Try a different song.")
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
		ctx.Reply("Could not find your current voice channel. Are you in a voice channel that the bot can join?")
		return
	}
	channelId := voiceChannel.ID
	if !chanManager.isChannel(channelId) {
		ctx.Reply("Not currently in the voice channel <#" + channelId + ">. For the bot to join, use `music join`!")
		return
	}
	chanManager.connections[channelId].connection.stop()
	ctx.Reply("Stopped playing music in <#" + channelId + ">!")
}

func stopBotCommand(ctx context) {
	if ctx.User.ID != conf.OwnerId {
		return
	}
	ctx.Reply("Bye :)")
	ctx.Discord.Close()
	os.Exit(-1)
}

func searchCommand(ctx context) {
	if len(ctx.Args) < 2 {
		ctx.Reply("Usage: music search <query>")
		return
	}
	query := strings.Join(ctx.Args[1:], " ")
	contents, err := searchYoutube(query)
	if err != nil {
		ctx.Reply("Something went wrong!")
		fmt.Println("err searching yt,", err)
		return
	}
	if contents == nil || len(contents) < 1 {
		ctx.Reply("No results found!")
		return
	}
	buffer := bytes.NewBufferString("Search results:")
	for index, content := range contents {
		write(buffer, "\n", strconv.Itoa(index+1), ". ", content.Title, " - ", content.ChannelTitle, " (", content.Id, ")")
	}
	ctx.Reply(buffer.String())
}
