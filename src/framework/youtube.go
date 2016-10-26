package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
)

type (
	response struct {
		Formats []struct {
			Url string `json:"url"`
		} `json:"formats"`
		Title string `json:"title"`
	}

	VideoResult struct {
		Media string
		Title string
	}

	YTSearchContent struct {
		Id           string `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ChannelTitle string `json:"channel_title"`
		Duration     string `json:"duration"`
	}

	ytApiResponse struct {
		Error   bool              `json:"error"`
		Content []YTSearchContent `json:"content"`
	}

	Youtube struct {
		Conf *Config
	}
)

func (youtube Youtube) Get(id string) (*VideoResult, error) {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "https://youtube.com/watch?v="+id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting youtube info,", err)
		return nil, err
	}
	var resp response
	json.Unmarshal(out.Bytes(), &resp)
	u := resp.Formats[0].Url
	return &VideoResult{u, resp.Title}, nil
}

func (youtube Youtube) buildUrl(query string) (*string, error) {
	base := youtube.Conf.ServiceUrl + "/v1/youtube/search"
	address, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("search", query)
	address.RawQuery = params.Encode()
	str := address.String()
	return &str, nil
}

func (youtube Youtube) Search(query string) ([]YTSearchContent, error) {
	addr, err := youtube.buildUrl(query)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(*addr)
	if err != nil {
		return nil, err
	}
	var apiResp ytApiResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	return apiResp.Content, nil
}
