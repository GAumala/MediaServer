package net
import (
    "log"
    "net/http"
    "os"
    "path/filepath"
)

const videocss = "video-js.css"
const videojs = "video.js"

func getProjectDir() string {
    pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
     if err != nil {
             log.Fatal(err)
     }
     return pwd
}

func getPublicDir() string {
    if(debug) {
        /*
        debug == true means that this process started with go run MediaServer.go
        in such scenario getProjectDir() returns a temp directory which does
        not have the html templates && assets. We can't get the absolute path of our
        Project directory, so let's return this instead.
        */
        return "public/"
    }
    return getProjectDir() +"/public/"
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
    publicDir := getPublicDir()
    requestedAsset := r.URL.Path
    if(requestedAsset == "/" + videojs) {
        http.ServeFile(w, r, publicDir + videojs)
    } else if (requestedAsset == "/" + videocss) {
        http.ServeFile(w, r, publicDir + videocss)
    } else {
        http.NotFound(w, r)
    }

}
