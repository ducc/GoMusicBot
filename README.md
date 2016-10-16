# GoMusicBot
A simple discord music bot using the [DiscordGo](https://github.com/bwmarrin/discordgo) library.

## Requirements
- go compiler
- gcc compiler
- ffmpeg.exe
- ffprobe.exe
- youtube-dl.exe
- python

## Install
1. Clone repository
2. Put `ffmpeg.exe` in the directory
3. Put `ffprobe.exe` in the directory
4. Put `youtube-dl.exe` in the directory
5. Create the `config.json` file and place this inside of the directory, following the template below
6. Create a `music` directory and place music files in here (any supported by ffmpeg)
7. Run `go build ./src/main` in the cloned directory
8. Execute `main.exe`

## Config
Create a `config.json` file in the project root using this template:
```json
{
	"bot_token": "Bot your-bot-token",
	"owner_id": "your user id",
	"use_sharding": false,
	"shard_id": 0,
	"shard_count": 1
}
```

## Commands
| Command           | Description                                                   |
|-------------------|---------------------------------------------------------------|
| music help        | shows all available commands                                  |
| music join        | joins your current voice channel (must be in a voice channel) |
| music play <file> | plays a music file                                            |
| music stop        | stops playing the current song                                |
| music leave       | leaves the voice channel                                      |
| music eval <code> | runs javascript (bot owner only)                              |
| music info        | shows bot info and statistics                                 |
| music stopbot     | stops the bot (bot owner only)                                |

## Support
This is not for public use. If you don't know GO I do not recommend using this and I will not provide support for that.

## Pull requests & issues
Go ahead, just bear in mind I've only been doing go for a few days.

## Credit
Developers of discordgo <3

https://github.com/iopred/bruxism for stats command