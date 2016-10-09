# MediaServer

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

Once you have everything installed, clone this repository into your Go path.

```
cd $GOPATH/src
git clone https://github.com/GAumala/MediaServer
```

Once you have the repository in your computer go to the root directory and
use the `make` command.

```
cd MediaServer
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
change it. you can pass a json file as parameter with an array of absolute paths
to directories of videos you want to serve. For example, you can create a
`paths.json` file like this:

```
[
  "/home/gabriel/Videos/",
  "/home/gabriel/Downloads/",
  "/home/gabriel/MEGAsync Downloads/"
]
```

The file specifies 3 directories that should be scanned to serve videos:
`~/Videos/`, `~/Downloads/`, and `~/MEGAsync Downloads/`. To use it just
execute:

```
./MediaServer paths.json
```

### Compatibility

This assumes that your are running a browser capable of playing/downloading
raw video files, such as Safari, Firefox or Google Chrome. Some browsers,
like the PS4's web browser, don't support this, but they do support streaming
with the HTML5 `video` tag. In those browsers, change your URL to
[localhost:8080/?player=1](http://localhost:8080/?player=1). With this URL, the
links point to an HTML file with a `video` tag instead of a raw video resource
file.
