package net

import (
    //"fmt"
    "log"
    "html/template"
    "net/http"
    //"sort"
    "path/filepath"
    "os"

    "MediaServer/data"
)

var videos []data.VideoDir
var ipAddr string
const debug = false
const port = ":8080"

func getProjectDir() string {
    pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
     if err != nil {
             log.Fatal(err)
     }
     return pwd
}

func getHtml(templateName string) string{
    var pwd string
    if(debug) {
        return "public/" + templateName
    }
    pwd = getProjectDir()
    log.Println(pwd + "/public/" + templateName)
    return pwd + "/public/" + templateName
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles(getHtml("listTemplate.html"))
    t.Execute(w, videos)
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
    pathStr := r.URL.Query().Get("p")
    log.Println("requested: ", pathStr)
    if(videoPathExists(pathStr)) {
        http.ServeFile(w, r, pathStr)
    }
}

func videoPathExists(videoPath string) bool {
    /*
    totalVideos := len(videos)
    dirName := path.filepath.DirName(videoPath)
    expectedPos := sort.Search(len(videos), func(index int) bool {
        return videos[index].FilePath >= videoPath
    })
    return expectedPos < totalVideos
    */
    return true
}

func initIpAddr() (ip string) {
    ip, err := externalIP()
	if err != nil {
		log.Fatal(err)
	}

    return
}

func RunServer(vidSlice []data.VideoDir){
    videos = vidSlice
    ipAddr = initIpAddr()
    log.Println("Running media server at: http://" + ipAddr + port)

    mux := http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(indexHandler))
    mux.Handle("/vid", http.HandlerFunc(videoHandler))
    http.ListenAndServe(port, mux)

}
