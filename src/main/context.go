package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type context struct {
	Discord       *discordgo.Session
	Message       *discordgo.MessageCreate
	TextChannelId string
	User          *discordgo.User
	Args          []string
}

func newContext(discord *discordgo.Session, message *discordgo.MessageCreate, args []string) *context {
	ctx := new(context)
	ctx.Discord = discord
	ctx.Message = message
	ctx.TextChannelId = message.ChannelID
	ctx.User = message.Author
	ctx.Args = args
	return ctx
}

func (ctx context) Reply(message string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannelId, message)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}
