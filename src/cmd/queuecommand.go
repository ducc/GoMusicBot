package cmd

import (
	"../framework"
	"bytes"
	"fmt"
	"strconv"
)

const (
	song_format = "\n`%03d` %s"
)

func QueueCommand(ctx framework.Context) {
	if len(ctx.Args) > 0 {
		AddCommand(ctx)
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
	queue := sess.Queue
	q := queue.Get()
	if len(q) == 0 && queue.Current() == nil {
		ctx.Reply("Song queue is empty! Add a song with `music add`.")
		return
	}
	buff := bytes.NewBufferString("__Song queue (" + strconv.Itoa(len(q)) + " songs)__")
	if queue.Current() != nil {
		buff.WriteString(fmt.Sprintf(song_format, 0, queue.Current().Title))
	}
	for index, song := range q {
		buff.WriteString(fmt.Sprintf(song_format, index+1, song.Title))
	}
	ctx.Reply(buff.String())
}
