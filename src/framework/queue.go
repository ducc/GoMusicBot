package framework

import (
	"fmt"
	"strconv"
)

type SongQueue struct {
	list    []Song
	running bool
}

func (queue SongQueue) Get() []Song {
	return queue.list
}

func (queue *SongQueue) Add(song Song) {
	queue.list = append(queue.list, song)
}

func (queue SongQueue) HasNext() bool {
	return len(queue.list) > 0
}

func (queue *SongQueue) Next() Song {
	song := queue.list[0]
	queue.list = queue.list[1:]
	fmt.Print("queue entries:")
	for i, entry := range queue.list {
		fmt.Print("\n" + strconv.Itoa(i) + ") " + entry.Title)
	}
	fmt.Println()
	return song
}

func (queue *SongQueue) Clear() {
	queue.list = make([]Song, 0)
    queue.running = false
}

func (queue *SongQueue) Start(ctx Context, sess *Session) {
    queue.running = true
	for queue.HasNext() && queue.running {
		song := queue.Next()
		ctx.Reply("Now playing `" + song.Title + "`.")
		sess.Play(song)
	}
    if !queue.running {
        ctx.Reply("Stopped playing.")
    } else {
        ctx.Reply("Finished queue.")
    }
}

func newSongQueue() *SongQueue {
	queue := new(SongQueue)
	queue.list = make([]Song, 0)
	return queue
}
