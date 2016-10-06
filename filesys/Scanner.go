package filesys

import (
    "log"
    //"io/ioutil"
    //"encoding/json"
    "os"
    "os/user"
    "path"
    "path/filepath"

    "MediaServer/data"
)

func getDefaultVideoDir() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    return usr.HomeDir + "/Videos/"
}

func findVideosInPath(root string, vids []data.VideoInfo) []data.VideoInfo {
    log.Println("looking for videos in: " + root)
    filepath.Walk(root,
        func(pathStr string, info os.FileInfo, err error) error {
        if(data.IsStreamableVideoFormat(path.Ext(pathStr))) {
            log.Printf("found: %s\n", pathStr)
            name := path.Base(pathStr)
            vids = append(vids, data.VideoInfo{pathStr, name})
        }
        return err
    })

    return vids
}

func getPathsToWalk(jsonFile string) []string {
    if(jsonFile != "") {
        jsonPaths, err := ParseJsonPathList(jsonFile)
        if(err!= nil) {
            log.Println(err)
        } else {
            return jsonPaths
        }
    }

    return []string{getDefaultVideoDir()}
}

func FindAllVideos(jsonFile string) []data.VideoInfo {
    pathsToWalk := getPathsToWalk(jsonFile)

    vids := make([]data.VideoInfo, 0, 0)
    for _, pathStr := range pathsToWalk {
        vids = findVideosInPath(pathStr, vids)
    }
    return vids
}
