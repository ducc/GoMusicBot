package cmd

import (
	"../framework"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	TYPE_UNKNOWN    = 0
	TYPE_YOUTUBE    = 1
	TYPE_SOUNDCLOUD = 2
)

type identifyResult struct {
	t uint8
	v string
}

func cont(input, check string) bool {
	return strings.Contains(input, check)
}

func identifyInput(input string) *identifyResult {
	lower := strings.ToLower(input)
	if cont(lower, "youtube.com") {
		u, err := url.Parse(input)
		if err != nil {
			fmt.Println("oops parsing inp", err)
			return nil
		}
		id := u.Query().Get("v")
		return &identifyResult{TYPE_YOUTUBE, id}
	} else if cont(lower, "youtu.be") {
		ind := strings.Index(input, "be/") + 3
		return &identifyResult{TYPE_YOUTUBE, input[ind:]}
	} else {
		return &identifyResult{TYPE_UNKNOWN, input}
	}
}

func loadYTSong(ctx framework.Context, id string) (*framework.Song, error) {
	result, err := ctx.Youtube.Get(id)
	if err != nil {
		return nil, err
	}
	return framework.NewSong(result.Media, result.Title, id), nil
}

func AddCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply("Add command usage: `music add <song>`\nValid inputs: `youtube url`, `soundcloud url`, " +
			"`youtube id`, `soundcloud id`")
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
	msg := ctx.Reply("Adding songs to queue...")
	for _, arg := range ctx.Args {
		identified := identifyInput(arg)
		if identified == nil {
			ctx.Reply("Could not identify input `" + arg + "`.")
			continue
		}
		id := identified.v
		var song *framework.Song
		var err error
		switch identified.t {
		case TYPE_UNKNOWN:
			// try youtube xd
			song, err = loadYTSong(ctx, id)
			break
		case TYPE_YOUTUBE:
			song, err = loadYTSong(ctx, id)
			break
		case TYPE_SOUNDCLOUD:
			ctx.Reply("Soundcloud not yet supported.")
			break
		default:
			ctx.Reply(fmt.Sprintf("Invalid type `%d`.", identified.t))
			return
		}
		if err != nil {
			ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "An error occured")
			fmt.Println("Error loading yt song,", err)
			return
		}
		sess.Queue.Add(*song)
		ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "Added `"+song.Title+"` to the song queue.")
	}
	ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "Added "+strconv.Itoa(len(ctx.Args))+
		" songs to the queue. Use `music play` to start playing the songs! To see the song queue, use `music queue`.")
}
