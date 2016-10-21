package framework

import (
	"../voice"
	"os/exec"
	"strconv"
)

type Song struct {
	Media string
	Title string
	Id    string
}

func (song Song) Ffmpeg() *exec.Cmd {
	return exec.Command("ffmpeg", "-i", song.Media, "-f", "s16le", "-ar", strconv.Itoa(voice.FRAME_RATE), "-ac",
		strconv.Itoa(voice.CHANNELS), "pipe:1")
}
