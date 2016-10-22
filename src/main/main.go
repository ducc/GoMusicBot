package main

import (
	"../cmd"
	"../framework"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	PREFIX = "music"
)

var (
	conf       *framework.Config
	CmdHandler *framework.CommandHandler
	Sessions   *framework.SessionManager
	botId      string
)

func main() {
	conf = framework.LoadConfig("config.json")
	CmdHandler = framework.NewCommandHandler()
	registerCommands()
	Sessions = framework.NewSessionManager()
	discord, err := discordgo.New(conf.BotToken)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}
	if conf.UseSharding {
		discord.ShardID = conf.ShardId
		discord.ShardCount = conf.ShardCount
	}
	usr, err := discord.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}
	botId = usr.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateStatus(0, "boyyyy")
		guilds := discord.State.Guilds
		fmt.Println("Ready with", len(guilds), "guilds.")
	})
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}
	fmt.Println("Started")
	<-make(chan struct{})
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botId || user.Bot {
		return
	}
	content := message.Content
	if len(content) <= len(PREFIX) {
		return
	}
	if content[:len(PREFIX)] != PREFIX {
		return
	}
	content = content[len(PREFIX)+1:]
	if len(content) < 1 {
		return
	}
	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := CmdHandler.Get(name)
	if !found {
		return
	}
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return
	}
	guild, err := discord.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return
	}
	ctx := framework.NewContext(discord, guild, channel, user, message, conf, CmdHandler, Sessions)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("help", cmd.HelpCommand)
	CmdHandler.Register("admin", cmd.AdminCommand)
	CmdHandler.Register("join", cmd.JoinCommand)
	CmdHandler.Register("leave", cmd.LeaveCommand)
	CmdHandler.Register("play", cmd.PlayCommand)
	CmdHandler.Register("stop", cmd.StopCommand)
    CmdHandler.Register("info", cmd.InfoCommand)
    CmdHandler.Register("add", cmd.AddCommand)
    CmdHandler.Register("skip", cmd.SkipCommand)
    CmdHandler.Register("queue", cmd.QueueCommand)
    CmdHandler.Register("eval", cmd.EvalCommand)
    CmdHandler.Register("debug", cmd.DebugCommand)
    CmdHandler.Register("clear", cmd.ClearCommand)
}
