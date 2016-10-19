package cmd

import (
    "../framework"
    "strings"
    "os"
)

func AdminCommand(ctx framework.Context) {
    if ctx.User.ID != ctx.Conf.OwnerId {
        return
    }
    if len(ctx.Args) == 0 {
        ctx.Reply("Usage: music admin <subcommand>\nSubcommands: stop")
        return
    }
    switch strings.ToLower(ctx.Args[0]) {
    case "stop":
        ctx.Reply("Bye :wave:")
        ctx.Discord.Close()
        os.Exit(-1)
        break
    default:
        ctx.Reply("Invalid subcommand!")
    }
}
