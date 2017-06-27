package main

import (
    "os"
    "github.com/GAumala/MediaServer/filesys"
    "github.com/GAumala/MediaServer/net"
)

func main() {
    var configFile string
    args := os.Args
    debug := os.Getenv("GODEBUG") != ""

    if len(args) > 1 {
        configFile = args[1]
    }

    net.RunServer(debug, filesys.FindAllVideos(configFile))
}
