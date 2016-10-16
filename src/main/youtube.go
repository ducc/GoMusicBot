package main

import (
	"os/exec"
    "fmt"
    "strings"
)

func downloadYT(ctx context, id string) (*string, error) {
	cmd := exec.Command("py", "youtube-dl -4 -x --extract-audio --audio-format mp3 --prefer-ffmpeg -o yt\\%(id)s.%(ext)s -- " + id)
    err := cmd.Run()
    if err != nil {
        fmt.Println("error running,", err)
    }
    str := strings.Join(cmd.Args, " ")
    fmt.Println("executed command: " + str)
	f := "yt\\" + id + ".mp3"
	return &f, nil
}