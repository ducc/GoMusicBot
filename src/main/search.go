package main

import (
	"fmt"
	"net/http"
	"net/url"
    "encoding/json"
)

func buildYtUrl(query string) (*string, error) {
	base := conf.ServiceUrl + "/v1/youtube/search"
	addr, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
    params := url.Values{}
    params.Add("search", query)
    addr.RawQuery = params.Encode()
	str := addr.String()
	return &str, nil
}

type content struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ChannelTitle string `json:"channel_title"`
	Duration     string `json:"duration"`
}

type apiResponse struct {
	Error   bool      `json:"error"`
	Content []content `json:"content"`
}

func searchYoutube(query string) ([]content, error) {
	addr, err := buildYtUrl(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(*addr)
	resp, err := http.Get(*addr)
	if err != nil {
		return nil, err
	}
    apiResp := new(apiResponse)
    json.NewDecoder(resp.Body).Decode(apiResp)
	return apiResp.Content, nil
}
