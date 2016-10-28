package cmd

import (
	"../framework"
	"bytes"
	"fmt"
	"strconv"
)

const (
	song_format = "\n`%03d` %s"
    current_format = "__Current song__\n%s\n"
    invalid_page = "Invalid page `%d`. Min: `1`, max: `%d`"
    response_footer = "\n\nPage **%d** of **%d**. To view the next page, use `music queue %d`."
)

func QueueCommand(ctx framework.Context) {
    sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
    if sess == nil {
        ctx.Reply("Not in a voice channel! To make the bot join one, use `music join`.")
        return
    }
    queue := sess.Queue
    q := queue.Get()
    if len(q) == 0 && queue.Current() == nil {
        ctx.Reply("Song queue is empty! Add a song with `music add`.")
        return
    }
    buff := bytes.Buffer{}
    if queue.Current() != nil {
        buff.WriteString(fmt.Sprintf(current_format, queue.Current().Title))
    }
    queueLength := len(q)
    if len(ctx.Args) == 0 {
        var resp string
        if queueLength > 20 {
            resp = display(q[:20], buff, 2, 0)
        } else {
            resp = display(q[:queueLength], buff, 2, 0)
        }
        ctx.Reply(resp)
        return
    }
    page, err := strconv.Atoi(ctx.Args[0])
    if err != nil {
        ctx.Reply("Invalid page `" + ctx.Args[0] + "`. Usage: `music queue <page>`")
        return
    }
    pages := queueLength / 20;
    if page < 1 || page > (pages + 1) {
        ctx.Reply(fmt.Sprintf(invalid_page, page, pages + 1, ))
        return
    }
    var lowerBound int
    if page == 1 {
        lowerBound = 0
    } else {
        lowerBound = (page - 1) * 20
    }
    upperBound := page * 20;
    if upperBound > queueLength {
        upperBound = queueLength
    }
    slice := q[lowerBound:upperBound]
    ctx.Reply(display(slice, buff, page + 1, lowerBound))
}

func display(queue []framework.Song, buff bytes.Buffer, page, start int) string {
    for index, song := range queue {
        buff.WriteString(fmt.Sprintf(song_format, start+index+1, song.Title))
    }
    buff.WriteString(fmt.Sprintf("\n\nView the next page: `music queue %d`", page))
    buff.WriteString(fmt.Sprintf(response_footer))
    return buff.String()
}
