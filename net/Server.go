package net

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/GAumala/MediaServer/data"
	"github.com/GAumala/MediaServer/filesys"
)

var config *data.Config
var videoList data.VideoDirectories
var videoDict data.VideoDict

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

	videoList, videoDict = filesys.FindAllVideos(*config)

	t := loadTemplate(indexTemplate)
	t.Execute(w, videoList)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")
	videoKey, err := strconv.ParseUint(v, 10, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqVid := videoDict[uint32(videoKey)]
	if reqVid.FilePath != "" {
		if config.Verbose {
			log.Println("Serving requested stream : ", videoKey)
		}

		t := loadTemplate("playerTemplate.html")
		t.Execute(w, reqVid)
	} else {
		if config.Verbose {
			log.Println("Unable to find requested stream : ", videoKey)
		}

		http.NotFound(w, r)
	}
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")
	videoKey, err := strconv.ParseUint(v, 10, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqVid := videoDict[uint32(videoKey)]
	if reqVid.FilePath != "" {
		if config.Verbose {
			log.Println("Serving requested raw video: ", videoKey)
		}

		http.ServeFile(w, r, reqVid.FilePath)
	} else {
		if config.Verbose {
			log.Println("Unable to find requested raw video: ", videoKey)
		}

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
