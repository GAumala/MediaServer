package net

import (
	"log"
	"net/http"
	"os"
	"path"
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

func getPublicFilePath(templateFileName string) string {
	return path.Join(os.Getenv("GOPATH"),
		"src/github.com/GAumala/MediaServer/public/",
		templateFileName)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	requestedAsset := r.URL.Path
	if requestedAsset == "/"+videojs {
		http.ServeFile(w, r, (videojs))
	} else if requestedAsset == "/"+videocss {
		http.ServeFile(w, r, getPublicFilePath(videocss))
	} else {
		http.NotFound(w, r)
	}

}
