package cmd

import (
	"../framework"
	"bytes"
)

func HelpCommand(ctx framework.Context) {
	cmds := ctx.CmdHandler.GetCmds()
	buffer := bytes.NewBufferString("Commands: ")
	for key := range cmds {
		if len(key) == 1 {
			continue
		}
		buffer.WriteString(key)
		buffer.WriteString(", ")
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}
