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
will find a simple HTML file listing all your available videos.


![screenshot from 2018-12-09 09-00-21](https://user-images.githubusercontent.com/5729175/49698430-ad3bb080-fb91-11e8-9047-abb3ea2b96c0.png)
![screenshot from 2018-12-09 09-03-39](https://user-images.githubusercontent.com/5729175/49698431-ad3bb080-fb91-11e8-8f7d-7d27bf814d42.png)

Notice that the url for watching a video could be something like 
`http://localhost:8080/watch?v=99653847.mp4`. You should be able to copy that
link into any player like vlc or mpv and stream it without issues.

## Configuration

By default it will serve the .mp4 videos in the path `/home/<USER>/Videos`. 
If you want to configure things like the port or the directories where to find 
videos you can pass a JSON file as parameter:

``` JSON
{
  "port": 5678,
  "verbose": true,
  "videoDirs": [ 
    "/home/gabriel/Videos/",
    "MEGAsync Downloads",
    "/home/gabriel/Downloads/" 
  ]
}
```

The file specifies a custom port 5678 and 3 directories that should be scanned 
to serve videos. To use it just execute:

```
./MediaServer myConfig.json
```

If you need more help about using this tool or how to manage your video library checkout [the wiki](https://github.com/GAumala/MediaServer/wiki).

### Raw files

By default the videos are streamed in a web client using video.js. The server
will try to figure out whether to serve html or raw video files based on the 
client request's `Accept` header. If you want to force the HTML list to show
raw file links your URL to [localhost:8080/?player=0](
http://localhost:8080/?player=0). This is useful when you have an unsupported
video player.

