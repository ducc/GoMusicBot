package cmd

import (
    "../framework"
    "../media"
    "fmt"
    "strconv"
)

func loadYTSong(id string) (*framework.Song, error) {
    result, err := media.Youtube(id)
    if err != nil {
        return nil, err
    }
    return &framework.Song{Media: result.Media, Title: result.Title, Id: id}, nil
}

func AddCommand(ctx framework.Context) {
    if len(ctx.Args) == 0 {
        ctx.Reply("Add command usage: `music add <song>`\nValid inputs: `youtube url`, `soundcloud url`, " +
                "`youtube id`, `soundcloud id`")
        return
    }
    sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
    if sess == nil {
        ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
        return
    }
    msg := ctx.Reply("Adding songs to queue...")
    for _, arg := range ctx.Args {
        song, err := loadYTSong(arg)
        if err != nil {
            ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "An error occured")
            fmt.Println("Error loading yt song,", err)
            return
        }
        sess.Queue.Add(*song)
        ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "Added `" + song.Title + "` to the song queue.")
    }
    ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "Added " + strconv.Itoa(len(ctx.Args)) +
            " songs to the queue. Use `music play` to start playing the songs!")
}