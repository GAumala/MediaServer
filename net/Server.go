package net

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

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

func htmlVideoHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

       	reqVid := videoDict[v]
	if reqVid.FilePath != "" {
		if config.Verbose {
			log.Println("Serving requested stream : ", v)
		}

		t := loadTemplate("playerTemplate.html")
		t.Execute(w, reqVid)
	} else {
		if config.Verbose {
			log.Println("Unable to find requested stream : ", v)
		}

		http.NotFound(w, r)
	}
}

func rawVideoHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")
	reqVid := videoDict[v]
	if reqVid.FilePath != "" {
		if config.Verbose {
			log.Println("Serving requested raw video: ", v)
		}

		http.ServeFile(w, r, reqVid.FilePath)
	} else {
		if config.Verbose {
			log.Println("Unable to find requested raw video: ", v)
		}

		http.NotFound(w, r)
	}
}

func watchHandler(w http.ResponseWriter, r *http.Request) {
  acceptHeader := r.Header.Get("Accept")
  if (strings.Contains(acceptHeader, "text/html")) {
    htmlVideoHandler(w, r)
  } else {
    rawVideoHandler(w, r)
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

	mux.Handle("/"+videojs, http.HandlerFunc(assetsHandler))
	mux.Handle("/"+videocss, http.HandlerFunc(assetsHandler))

	mux.Handle("/watch", http.HandlerFunc(watchHandler))
	mux.Handle("/vid", http.HandlerFunc(rawVideoHandler))
	log.Fatal(http.ListenAndServe(portAddr, mux))
}
