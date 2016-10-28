package cmd

import (
	"../framework"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

const (
	invalid_song_format = "Invalid song number `%d`. Min: 1, max: %d"
)

func PickCommand(ctx framework.Context) {
	argsLen := len(ctx.Args)
	if argsLen == 0 {
		ctx.Reply("Usage: `music pick <result number>`")
		return
	}
	if argsLen > 5 {
		ctx.Reply("You cannot pick more than 5 songs at once.")
		return
	}
	identifier := ytSessionIdentifier(ctx.User, ctx.TextChannel)
	var ytSession ytSearchSession
	var ok bool
	if ytSession, ok = ytSessions[identifier]; !ok {
		ctx.Reply("You haven't searched for a song yet!")
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
	rLen := len(ytSession.results)
	var msg *discordgo.Message
	for i := 0; i < argsLen; i++ {
		num, err := strconv.Atoi(ctx.Args[i])
		if err != nil {
			ctx.Reply("An error occured!")
			fmt.Print("Error parsing int,", err)
			return
		}
		if num < 1 || num > rLen {
			ctx.Reply(fmt.Sprintf(invalid_song_format, num, rLen))
			return
		}
		result := ytSession.results[num-1]
		song, err := loadYTSong(ctx, result.Id)
		sess.Queue.Add(*song)
		if msg != nil {
			msg, err = ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, msg.Content+", `"+song.Title+"`")
		} else {
			msg = ctx.Reply("Added `" + song.Title + "`")
		}
	}
	ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, msg.Content+
		" to the song queue.\nUse **music play** to start playing the songs! To see the song queue, use **music queue**.")
}
