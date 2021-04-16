# DiscordMusicCollector

DMC listens in whichever channels you allow it to and watches messages
for links to popular music services. It adds these links to a database
which can be used for any number of things, including displaying a
webpage of all music shared in a server, or to automatically build
playlists.

## Usage

```
$ git clone https://git.sr.ht/~mjorgensen/DiscordMusicCollector
$ cd DiscordMusicCollector
$ go build DiscordMusicCollector.go
$ ./DiscordMusicCollector -t YOUR_BOT_TOKEN
```

## Resources

Discussion and paches are welcome and should be directed to my public
inbox: [~mjorgensen@lists.sr.ht][lists]. Please use `--subject-prefix
PATCH DiscordMusicCollector` for clarity when sending patches.

Bugs and issues are tracked in the tracker: 
[~mjorgensen/DiscordMusicCollector][todo].

[lists]:https://lists.sr.ht/~mjorgensen/public-inbox
[todo]:https://todo.sr.ht/~mjorgensen/DiscordMusicCollector
