package cmd

import (
	"../framework"
)

func ClearCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
    if !sess.Queue.HasNext() {
        ctx.Reply("Queue is already empty")
        return
    }
    sess.Queue.Clear()
    ctx.Reply("Cleared the song queue")
}
