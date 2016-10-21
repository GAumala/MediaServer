GO_FILES := $(shell find -type f -name '*.go')
HTML_FILES := $(shell find -type f -name '*.html')
all: MediaServer

MediaServer: $(GO_FILES) $(HTML_FILES) public/video.js public/video-js.css
	go build MediaServer.go

public/video-js.css:
	wget -O public/video-js.css http://vjs.zencdn.net/5.11.7/video-js.css

public/video.js:
	wget -O public/video.js http://vjs.zencdn.net/5.11.7/video.js

clean:
	rm MediaServer
