package framework

type SongQueue struct {
	list    []Song
	current *Song
	running bool
}

func (queue SongQueue) Get() []Song {
	return queue.list
}

func (queue *SongQueue) Set(list []Song) {
    queue.list = list
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
	queue.current = &song
	return song
}

func (queue *SongQueue) Clear() {
	queue.list = make([]Song, 0)
	queue.running = false
	queue.current = nil
}

func (queue *SongQueue) Start(sess *Session, callback func(string)) {
	queue.running = true
	for queue.HasNext() && queue.running {
		song := queue.Next()
		callback("Now playing `" + song.Title + "`.")
		sess.Play(song)
	}
	if !queue.running {
		callback("Stopped playing.")
	} else {
		callback("Finished queue.")
	}
}

func (queue *SongQueue) Current() *Song {
	return queue.current
}

func (queue *SongQueue) Pause() {
    queue.running = false
}

func newSongQueue() *SongQueue {
	queue := new(SongQueue)
	queue.list = make([]Song, 0)
	return queue
}
