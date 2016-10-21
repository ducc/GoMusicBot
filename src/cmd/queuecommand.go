package cmd

import (
    "../framework"
    "bytes"
    "strconv"
)

func QueueCommand(ctx framework.Context) {
    sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
    if sess == nil {
        ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
        return
    }
    queue := sess.Queue
    q := queue.Get()
    buff := bytes.NewBufferString("__Song queue (" + strconv.Itoa(len(q)) + " songs)__")
    for index, song := range q {
        buff.WriteString("\n`" + strconv.Itoa(index + 1) + "` " + song.Title)
    }
    ctx.Reply(buff.String())
}