package cmd

import (
	"../framework"
)

func StopCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
    if sess.Queue.HasNext() {
        sess.Queue.Clear()
    }
	sess.Stop()
}
