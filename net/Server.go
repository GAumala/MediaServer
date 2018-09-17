package net

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/GAumala/MediaServer/data"
	"github.com/GAumala/MediaServer/filesys"
)

var config *data.Config
var videos data.VideoDirectories

func loadTemplate(name string) *template.Template {
	templatePath := getPublicFilePath(name)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalln("Failed to load template at ", templatePath, ". Error:\n", err)
	}
	return t
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate := "streamListTemplate.html"

	streamOnly := r.URL.Query().Get("player")
	if streamOnly == "0" {
		indexTemplate = "listTemplate.html"
	}

	videos = filesys.FindAllVideos(*config)

	t := loadTemplate(indexTemplate)
	t.Execute(w, videos)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.URL.Query().Get("p")
	if config.Verbose {
		log.Println("requested: stream", pathStr)
	}

	reqVid := videos.FindVideo(pathStr)
	if reqVid.FilePath != "" {
		t := loadTemplate("playerTemplate.html")
		t.Execute(w, reqVid)
	} else {
		http.NotFound(w, r)
	}
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	pathStr := r.URL.Query().Get("p")
	if config.Verbose {
		log.Println("requested: ", pathStr)
	}

	reqVid := videos.FindVideo(pathStr)
	if reqVid.FilePath != "" {
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

/*RunServer starts the HTTP server at port 8080 using the specified config.
 */
func RunServer(c *data.Config) {
	config = c
	ipAddr := initIPAddr()
	portAddr := fmt.Sprintf(":%d", config.Port)

	fmt.Printf("Running media server at: http://%s%s\n", ipAddr, portAddr)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(indexHandler))
	mux.Handle("/vid", http.HandlerFunc(videoHandler))

	mux.Handle("/"+videojs, http.HandlerFunc(assetsHandler))
	mux.Handle("/"+videocss, http.HandlerFunc(assetsHandler))

	mux.Handle("/watch", http.HandlerFunc(streamHandler))
	log.Fatal(http.ListenAndServe(portAddr, mux))
}
