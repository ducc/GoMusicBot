package main

import (
    "os/exec"
    "strconv"
)

type Song struct {
    filename    string
}

func (song Song) ffmpeg() *exec.Cmd {
    return exec.Command("ffmpeg", "-i", song.filename, "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac",
        strconv.Itoa(CHANNELS), "pipe:1")
}