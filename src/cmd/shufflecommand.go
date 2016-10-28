package cmd

import (
    "../framework"
    "math/rand"
)

func ShuffleCommand(ctx framework.Context) {
    sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
    if sess == nil {
        ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
        return
    }
    queue := sess.Queue
    if !queue.HasNext() {
        ctx.Reply("Queue is empty! Add songs with `music add`.")
        return
    }
    dest := shuffleLoop(queue.Get(), 3)
    queue.Set(dest)
    ctx.Reply("Shuffled the song queue!")
}

func shuffleLoop(list []framework.Song, i int) []framework.Song {
    for x := 0; x < i; x++ {
        list = shuffle(list)
    }
    return list
}

func shuffle(list []framework.Song) []framework.Song {
    dest := make([]framework.Song, len(list))
    perm := rand.Perm(len(list))
    for i, v := range perm {
        dest[v] = list[i]
    }
    return dest
}