package main

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
)

type context struct {
    discord         *discordgo.Session
    message         *discordgo.MessageCreate
    textChannelId   string
    user            *discordgo.User
    args            []string
}

func newContext(discord *discordgo.Session, message *discordgo.MessageCreate, args []string) *context {
    ctx := new(context)
    ctx.discord = discord
    ctx.message = message
    ctx.textChannelId = message.ChannelID
    ctx.user = message.Author
    ctx.args = args
    return ctx
}

func (ctx context) reply(message string) *discordgo.Message {
    msg, err := ctx.discord.ChannelMessageSend(ctx.textChannelId, message)
    if err != nil {
        fmt.Println("Error whilst sending message,", err)
        return nil
    }
    return msg
}