package cmd

import (
	"../framework"
)

func JoinCommand(ctx framework.Context) {
	if ctx.Sessions.GetByGuild(ctx.Guild.ID) != nil {
		ctx.Reply("Already connected! Use `music leave` for the bot to disconnect.")
		return
	}
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		ctx.Reply("You must be in a voice channel to use the bot!")
		return
	}
	sess, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, framework.JoinProperties{
		Muted:    false,
		Deafened: true,
	})
	if err != nil {
		ctx.Reply("An error occured!")
		return
	}
	ctx.Reply("Joined <#" + sess.ChannelId + ">!")
}
