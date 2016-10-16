package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	PREFIX = "music"
)

var (
	cmdManager  *commandManager
	chanManager *channelManager
	conf        *config
	botId       string
)

func main() {
	cmdManager = newCommandManager()
	registerCommands()
	chanManager = newChannelManager()
	conf = loadConfig("config.json")
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
		fmt.Println("Ready")
		discord.UpdateStatus(0, "boyyyy")
		guilds := discord.State.Guilds
		fmt.Println("num guilds:", len(guilds))
	})
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}
	fmt.Println("Started")
	<-make(chan struct{})
}

func registerCommands() {
	cmdManager.register("help", helpCommand)
	cmdManager.register("join", joinCommand)
	cmdManager.register("leave", leaveCommand)
	cmdManager.register("play", playCommand)
	cmdManager.register("stop", stopCommand)
	cmdManager.register("eval", evalCommand)
    cmdManager.register("info", infoCommand)
    cmdManager.register("stopbot", stopBotCommand)
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
	if !cmdManager.isCommand(name) {
		return
	}
	ctx := newContext(discord, message, args)
	cmdManager.commands[name](*ctx)
}
