package cmd

import (
    "../framework"
    "bytes"
    "strconv"
    "fmt"
)

const num_format = "%03d"

func formatNum(input int) string {
    return fmt.Sprintf(num_format, input)
}

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
    if len(q) == 0 {
        ctx.Reply("Song queue is empty! Add a song with `music add`.")
        return
    }
    buff := bytes.NewBufferString("__Song queue (" + strconv.Itoa(len(q)) + " songs)__")
    for index, song := range q {
        buff.WriteString("\n`" + formatNum(index + 1) + "` " + song.Title)
    }
    ctx.Reply(buff.String())
}