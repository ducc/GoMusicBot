package media

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type response struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
	Title string `json:"title"`
}

type VideoResult struct {
	Media string
	Title string
}

func Youtube(id string) (*VideoResult, error) {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "https://youtube.com/watch?v="+id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting youtube info,", err)
		return nil, err
	}
	resp := new(response)
	json.Unmarshal(out.Bytes(), resp)
	url := resp.Formats[0].Url
	return &VideoResult{url, resp.Title}, nil
}
