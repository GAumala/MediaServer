# MediaServer

[![Build Status](https://travis-ci.org/GAumala/MediaServer.svg?branch=master)](https://travis-ci.org/GAumala/MediaServer)

Media Server for streaming videos over local network written in Go. The server
has no dependencies other than the Go standard library. The web client uses
[Video.js](http://videojs.com/)

## Build
Currently, only Linux is supported. To build in your computer, you need to have
installed all the build dependencies:

- git
- go
- make
- wget

Once you have everything installed, grab the code frome github.

```
go get github.com/GAumala/MediaServer
```

Once you have the repository in your computer go to the root directory and
use the `make` command.

```
cd $GOPATH/src/github.com/GAumala/MediaServer
make
```

This command downloads the necessary Video.js files if they are missing and
creates an executable named `MediaServer`. As long as the `public`
folder is its same directory it will work properly.

## Usage

You can start the server with:

```
./MediaServer
```

If you open your browser and visit [localhost:8080](http://localhost:8080) you
will find a simple HTML file listing all your available videos. By default it
will serve the .mp4 videos in the path `/home/<USER>/Videos`. If you want to
configure things like the port or the directories where to find videos you can
pass a JSON file as parameter:

```
{
  "port": 5678,
  "verbose": true,
  "videoDirs": [ 
    "~/Videos/",
    "/home/gabriel/MEGAsync Downloads",
    "~/Downloads/" ]
}
```

The file specifies a custom port 5678 and 3 directories that should be scanned 
to serve videos: `~/Videos/`, `~/Downloads/`, and `~/MEGAsync Downloads/`. To 
use it just execute:

```
./MediaServer paths.json
```

If you need more help about using this tool or how to manage your video library checkout [the wiki](https://github.com/GAumala/MediaServer/wiki).

### Raw files

By default the videos are streamed in a web client using video.js. If you want
to have access to the raw video files change your URL to [localhost:8080/?player=1](
http://localhost:8080/?player=0). This is useful when you want to use different
video players that support video streaming over network, such as mpv.

