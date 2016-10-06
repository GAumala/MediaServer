# MediaServer

Media Server for streaming videos over local network written in Go. It has no
dependencies other than the Go standard library.

## Build
To build in your computer, clone this repository and in the root directory
execute: 

```
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

