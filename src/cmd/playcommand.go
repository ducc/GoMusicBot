package cmd

import (
	"../framework"
)

func PlayCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
		return
	}
	queue := sess.Queue
	if !queue.HasNext() {
		ctx.Reply("Queue is empty! Add songs with `music add`.")
		return
	}
	go queue.Start(sess, func(msg string) {
		ctx.Reply(msg)
	})
}
