# MediaServer

Media Server for streaming videos over local network written in Go. It has no
dependencies other than the Go standard library.

## Build
To build in your computer, first clone this repository into your Go path

```
cd $GOPATH/src
git clone https://github.com/GAumala/MediaServer
```

Once you have the repository in your computer go to the root directory and
use the `build` command.

```
cd MediaServer
go build MediaServer.go
```

This command creates an executable named `MediaServer`. As long as the `public`
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

## Development

To test your local changes to the MediaServer go to the root directory of the
repository and use the `run` command.

```
GODEBUG=1 go run MediaServer.go
```

The `GODEBUG` must be used because `run` creates an executable on a different
directory, so it doesn't have access to the HTML templates in the `public`
directory.
