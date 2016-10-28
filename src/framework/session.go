package framework

import (
	"github.com/bwmarrin/discordgo"
)

type (
	Session struct {
		Queue              *SongQueue
		guildId, ChannelId string
		connection         *Connection
	}

	SessionManager struct {
		sessions map[string]*Session
	}

	JoinProperties struct {
		Muted    bool
		Deafened bool
	}
)

func newSession(guildId, channelId string, connection *Connection) *Session {
	session := new(Session)
	session.Queue = newSongQueue()
	session.guildId = guildId
	session.ChannelId = channelId
	session.connection = connection
	return session
}

func (sess Session) Play(song Song) error {
	return sess.connection.Play(song.Ffmpeg())
}

func (sess *Session) Stop() {
	sess.connection.Stop()
}

func NewSessionManager() *SessionManager {
	return &SessionManager{make(map[string]*Session)}
}

func (manager SessionManager) GetByGuild(guildId string) *Session {
	for _, sess := range manager.sessions {
		if sess.guildId == guildId {
			return sess
		}
	}
	return nil
}

func (manager SessionManager) GetByChannel(channelId string) (*Session, bool) {
	sess, found := manager.sessions[channelId]
	return sess, found
}

func (manager *SessionManager) Join(discord *discordgo.Session, guildId, channelId string,
	properties JoinProperties) (*Session, error) {
	vc, err := discord.ChannelVoiceJoin(guildId, channelId, properties.Muted, properties.Deafened)
	if err != nil {
		return nil, err
	}
	sess := newSession(guildId, channelId, NewConnection(vc))
	manager.sessions[channelId] = sess
	return sess, nil
}

func (manager *SessionManager) Leave(discord *discordgo.Session, session Session) {
	session.connection.Stop()
	session.connection.Disconnect()
	delete(manager.sessions, session.ChannelId)
}
