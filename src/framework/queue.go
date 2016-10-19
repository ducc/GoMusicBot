package framework

type SongQueue struct {
	list []Song
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
	queue.list = append(queue.list[1:], queue.list[1:]...)
	return song
}

func newSongQueue() *SongQueue {
	queue := new(SongQueue)
	queue.list = make([]Song, 0)
	return queue
}
