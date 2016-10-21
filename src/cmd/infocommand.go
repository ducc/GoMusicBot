package cmd

import (
    "time"
    "fmt"
    "runtime"
    "bytes"
    "github.com/bwmarrin/discordgo"
    "github.com/dustin/go-humanize"
    "strconv"
    "../framework"
)

// credit to github.com/iopred/bruxism for this command

var startTime = time.Now()
var userString *string

func getDurationString(duration time.Duration) string {
    return fmt.Sprintf(
        "%0.2d:%02d:%02d",
        int(duration.Hours()),
        int(duration.Minutes())%60,
        int(duration.Seconds())%60,
    )
}

func InfoCommand(ctx framework.Context) {
    if userString == nil {
        usr, err := ctx.Discord.User(ctx.Conf.OwnerId)
        if err != nil {
            fmt.Println("error getting user ", ctx.Conf.OwnerId, err)
            return
        }
        str := usr.Username + "#" + usr.Discriminator
        userString = &str
    }
    stats := runtime.MemStats{}
    runtime.ReadMemStats(&stats)
    buffer := bytes.NewBufferString("```")
    write(buffer, "owner: ", *userString)
    write(buffer, "\ngo version: ", runtime.Version())
    write(buffer, "\ndiscordgo version: ", discordgo.VERSION)
    write(buffer, "\nuptime: ", getDurationString(time.Now().Sub(startTime)))
    buffer.WriteString(fmt.Sprintf("\nmemory used: %s / %s (%s garbage collected)", humanize.Bytes(stats.Alloc),
        humanize.Bytes(stats.Sys), humanize.Bytes(stats.TotalAlloc)))
    write(buffer, "\nconcurrent tasks: ", strconv.Itoa(runtime.NumGoroutine()))
    write(buffer, "\ncurrent shard: ", strconv.Itoa(ctx.Discord.ShardID))
    write(buffer, "\nshard count: ", strconv.Itoa(ctx.Discord.ShardCount))
    buffer.WriteString("```")
    ctx.Reply(buffer.String())
}