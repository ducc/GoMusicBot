package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	BotToken   string `json:"bot_token"`
	OwnerId    string `json:"owner_id"`
	UseSharding      bool   `json:"use_sharding"`
	ShardId    int    `json:"shard_id"`
	ShardCount int    `json:"shard_count"`
}

func loadConfig(filename string) *config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config,", err)
		return nil
	}
	var conf config
	json.Unmarshal(body, &conf)
	return &conf
}
