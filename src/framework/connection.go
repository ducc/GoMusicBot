package framework

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

type Connection struct {
	voiceConnection *discordgo.VoiceConnection
	send            chan []int16
	lock            sync.Mutex
	sendpcm         bool
	stopRunning     bool
	playing         bool
}

func NewConnection(voiceConnection *discordgo.VoiceConnection) *Connection {
	connection := new(Connection)
	connection.voiceConnection = voiceConnection
	return connection
}
func (connection Connection) Disconnect() {
	connection.voiceConnection.Disconnect()
}
