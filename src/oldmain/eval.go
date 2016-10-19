package oldmain

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robertkrimen/otto"
	"strings"
)

var vm = otto.New()

func evalCommand(ctx context) {
	if conf.OwnerId != ctx.User.ID {
		return
	}
	js := strings.Join(ctx.Args[1:], " ")
	vm.Set("ctx", ctx)
	vm.Set("getGuild", Guild)
	vm.Set("getEmojis", Emojis)
	val, err := vm.Run(js)
	if err != nil {
		ctx.Reply(err.Error())
		return
	}
	if val.IsNull() {
		return
	}
	ctx.Reply("`" + val.String() + "`")
}

func Guild(ctx context, id string) *discordgo.Guild {
	guild, err := ctx.Discord.State.Guild(id)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return nil
	}
	return guild
}

func Emojis(ctx context, guildId string) []string {
	guild := Guild(ctx, guildId)
	arr := make([]string, 0)
	for _, emoj := range guild.Emojis {
		arr = append(arr, "<:"+emoj.Name+":"+emoj.ID+">")
	}
	return arr
}
