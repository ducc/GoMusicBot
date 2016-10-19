package cmd

import (
	"../framework"
)

func JoinCommand(ctx framework.Context) {
	vc := ctx.VoiceChannel
	if vc != nil {
		ctx.Reply("Already connected! Use `music leave` for the bot to disconnect.")
		return
	}
	vc = ctx.GetVoiceChannel()
	_, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, framework.JoinProperties{
		Muted:    false,
		Deafened: true,
	})
	if err != nil {
		ctx.Reply("An error occured!")
		return
	}
	ctx.Reply("ok :)")
}
