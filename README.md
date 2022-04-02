[![Go](https://github.com/prplecake/DiscordMusicCollector/actions/workflows/go.yml/badge.svg)](https://github.com/prplecake/DiscordMusicCollector/actions/workflows/go.yml)
[![CodeFactor](https://www.codefactor.io/repository/github/prplecake/discordmusiccollector/badge)](https://www.codefactor.io/repository/github/prplecake/discordmusiccollector)

# DiscordMusicCollector

DMC listens in whichever channels you allow it to and watches messages
for links to popular music services. It adds these links to a database
which can be used for any number of things, including displaying a
webpage of all music shared in a server, or to automatically build
playlists.

## Getting Started

Basically you need to clone the repo, create your configuration, and run
the thing.

```text
git clone https://github.com/prplecake/DiscordMusicCollector
cd DiscordMusicCollector
cp config.sample.yaml config.yaml
$EDITOR config.yaml
go build cmd/DiscordMusicCollector
./DiscordMusicCollector
```
