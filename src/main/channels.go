package main

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
)

type channel struct {
    guildId string
    channelId string
    connection connection
}

type channelManager struct {
    connections map[string]*channel
}

func newChannelManager() *channelManager {
    return &channelManager{make(map[string]*channel)}
}

func (manager channelManager) isChannel(channelId string) bool {
    if _, ok := manager.connections[channelId]; ok {
        return true
    }
    return false
}

func (manager channelManager) getChannelByGuild(guildId string) *channel {
    for _, channel := range manager.connections {
        if channel.guildId == guildId {
            return channel
        }
    }
    return nil
}

func (manager channelManager) joinChannelDeafened(discord *discordgo.Session, guildId, channelId string) (*channel, error) {
    return manager.joinChannel(discord, guildId, channelId, false, true)
}

func (manager channelManager) joinChannel(discord *discordgo.Session, guildId, channelId string, muted, deafened bool) (*channel, error) {
    voice, err := discord.ChannelVoiceJoin(guildId, channelId, muted, deafened)
    if err != nil {
        return nil, err
    }
    ch := &channel{guildId, channelId, *newConnection(voice)}
    manager.connections[channelId] = ch
    fmt.Println("Added " + channelId + " to chanManager")
    return ch, nil
}

func (manager channelManager) leaveChannel(discord *discordgo.Session, channelId string) {
    ch := manager.connections[channelId]
    ch.connection.stop()
    ch.connection.voiceConnection.Disconnect()
    delete(manager.connections, channelId)
    fmt.Println("Removed " + channelId + " from chanManager")
}