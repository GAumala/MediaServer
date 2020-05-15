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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getPublicFilePath(templateFileName string) string {
	filename := path.Join(os.Getenv("GOPATH"),
		"src/github.com/GAumala/MediaServer/public/",
		templateFileName)
	if !fileExists(filename) { // Use executable file path if GOPATH not exists
		filename = path.Join(getProjectDir(), "public/", templateFileName)
	}
	return filename
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	requestedAsset := r.URL.Path
	if requestedAsset == "/"+videojs {
		http.ServeFile(w, r, getPublicFilePath(videojs))
	} else if requestedAsset == "/"+videocss {
		http.ServeFile(w, r, getPublicFilePath(videocss))
	} else {
		http.NotFound(w, r)
	}

}
