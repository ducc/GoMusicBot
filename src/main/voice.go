package main

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
    "bufio"
    "sync"
    "github.com/layeh/gopus"
    "encoding/binary"
    "io"
    "errors"
)

const (
    CHANNELS    int = 2
    FRAME_RATE  int = 48000
    FRAME_SIZE  int = 960
    MAX_BYTES   int = (FRAME_SIZE * 2) * 2
)

type connection struct {
    voiceConnection *discordgo.VoiceConnection
    send            chan []int16
    lock            sync.Mutex
    sendpcm         bool
    stopRunning     bool
    playing         bool
}

func newConnection(voiceConnection *discordgo.VoiceConnection) *connection  {
    connection := new(connection)
    connection.voiceConnection = voiceConnection
    return connection
}

func (connection *connection) sendPCM(voice *discordgo.VoiceConnection, pcm <-chan []int16) {
    connection.lock.Lock()
    if connection.sendpcm || pcm == nil {
        connection.lock.Unlock()
        return
    }
    connection.sendpcm = true
    connection.lock.Unlock()
    defer func() {
        connection.sendpcm = false
    }()
    encoder, err := gopus.NewEncoder(FRAME_RATE, CHANNELS, gopus.Audio)
    if err != nil {
        fmt.Println("NewEncoder error,", err)
        return
    }
    for {
        receive, ok := <-pcm
        if !ok {
            fmt.Println("PCM channel closed")
            return
        }
        opus, err := encoder.Encode(receive, FRAME_SIZE, MAX_BYTES)
        if err != nil {
            fmt.Println("Encoding error,", err)
            return
        }
        if !voice.Ready || voice.OpusSend == nil {
            fmt.Printf("Discordgo not ready for opus packets. %+v : %+v", voice.Ready, voice.OpusSend)
            return
        }
        voice.OpusSend <- opus
    }
}

func (connection *connection) play(song Song) error {
    if (connection.playing) {
        return errors.New("song already playing")
    }
    connection.stopRunning = false
    ffmpeg := song.ffmpeg()
    out, err := ffmpeg.StdoutPipe()
    if err != nil {
        return err
    }
    buffer := bufio.NewReaderSize(out, 16384)
    err = ffmpeg.Start()
    if err != nil {
        return err
    }
    connection.playing = true
    defer func() {
        connection.playing = false
    }()
    connection.voiceConnection.Speaking(true)
    defer connection.voiceConnection.Speaking(false)
    if connection.send == nil {
        connection.send = make(chan []int16, 2)
    }
    go connection.sendPCM(connection.voiceConnection, connection.send)
    for {
        if (connection.stopRunning) {
            ffmpeg.Process.Kill()
        }
        audioBuffer := make([]int16, FRAME_SIZE * CHANNELS)
        err = binary.Read(buffer, binary.LittleEndian, &audioBuffer)
        if err == io.EOF || err == io.ErrUnexpectedEOF {
            return nil
        }
        if err != nil {
            return err
        }
        connection.send <- audioBuffer
    }
    return nil
}

func (connection *connection) stop() {
    connection.stopRunning = true
    connection.playing = false
}