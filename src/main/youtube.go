package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
)

type response struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
}

func getYoutubeUrl(id string) string {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "https://youtube.com/watch?v=" + id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("Run: ", err)
	}
	resp := new(response)
	json.Unmarshal(out.Bytes(), resp)
	return resp.Formats[0].Url
}
