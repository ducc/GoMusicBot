package cmd

import (
	"../framework"
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const result_format = "\n`%d` %s - %s (%s)"

var ytSessions ytSearchSessions = make(ytSearchSessions)

type (
	ytSearchSessions map[string]ytSearchSession

	ytSearchSession struct {
		results []framework.YTSearchContent
	}
)

func ytSessionIdentifier(user *discordgo.User, channel *discordgo.Channel) string {
	return user.ID + channel.ID
}

func formatDuration(input string) string {
	return parseISO8601(input).String()
}

func YoutubeCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply("Usage: `music youtube <search query>`")
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
	query := strings.Join(ctx.Args, " ")
	results, err := ctx.Youtube.Search(query)
	if err != nil {
		ctx.Reply("An error occured!")
		fmt.Println("Error searching youtube,", err)
		return
	}
    if len(results) == 0 {
        ctx.Reply("No results found for your query `" + query + "`.")
        return
    }
	buffer := bytes.NewBufferString("__Search results__ for `" + query + "`:\n")
	for index, result := range results {
		buffer.WriteString(fmt.Sprintf(result_format, index+1, result.Title, result.ChannelTitle,
			formatDuration(result.Duration)))
	}
	buffer.WriteString("\n\nTo pick a song, use `music pick <number>`.")
	ytSessions[ytSessionIdentifier(ctx.User, ctx.TextChannel)] = ytSearchSession{results}
	ctx.Reply(buffer.String())
}
