package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)

type config struct {
    BotToken string `json:"bot_token"`
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