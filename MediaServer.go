package main

import (
    "os"
    "MediaServer/filesys"
    "MediaServer/net"
)

func main() {
    var configFile string
    args := os.Args
    if len(args) > 1 {
        configFile = args[1]
    }

    net.RunServer(filesys.FindAllVideos(configFile))
}
