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

var videos data.VideoDirectories
var ipAddr string
var debug bool
const port = ":8080"

func getProjectDir() string {
    pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
     if err != nil {
             log.Fatal(err)
     }
     return pwd
}

func getHTML(templateName string) string{
    var pwd string
    if(debug) {
        /*
        debug == true means that this process started with go run MediaServer.go
        in such scenario getProjectDir() returns a temp directory which does
        not have the html templates. We can't get the absolute path of our
        Project directory, so let's return this instead.
        */
        return "public/" + templateName
    }
    pwd = getProjectDir()
    log.Println(pwd + "/public/" + templateName)
    return pwd + "/public/" + templateName
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    indexTemplate := "listTemplate.html"

    streamOnly := r.URL.Query().Get("player")
    if(streamOnly == "1") {
        indexTemplate = "streamListTemplate.html"
    }

    t, _ := template.ParseFiles(getHTML(indexTemplate))
    t.Execute(w, videos)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
    pathStr := r.URL.Query().Get("p")
    log.Println("requested: stream", pathStr)

    reqVid := videos.FindVideo(pathStr)
    if(reqVid.FilePath != "") {
        t, _ := template.ParseFiles(getHTML("playerTemplate.html"))
        t.Execute(w, reqVid)
    } else {
        http.NotFound(w, r)
    }
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
    pathStr := r.URL.Query().Get("p")
    log.Println("requested: ", pathStr)

    reqVid := videos.FindVideo(pathStr)
    if(reqVid.FilePath != "") {
        http.ServeFile(w, r, pathStr)
    } else {
        http.NotFound(w, r)
    }
}

func initIPAddr() (ip string) {
    ip, err := externalIP()
	if err != nil {
		log.Println(err)
        ip = "localhost"
	}
    return
}

/*RunServer starts the HTTP server at port 8080 that will serve the videos
* in vidSlice. The goDebug value should be true only during development.
*/
func RunServer(goDebug bool, vidSlice []data.VideoDir){
    debug = goDebug
    videos = vidSlice
    ipAddr = initIPAddr()
    log.Println("Running media server at: http://" + ipAddr + port)

    mux := http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(indexHandler))
    mux.Handle("/vid", http.HandlerFunc(videoHandler))
    mux.Handle("/watch", http.HandlerFunc(streamHandler))
    http.ListenAndServe(port, mux)

}
