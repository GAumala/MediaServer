package filesys

import (
    "log"
    "os"
    "os/user"
    "path"
    "path/filepath"
    "sort"

    "MediaServer/data"
)

func getDefaultVideoDir() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    return usr.HomeDir + "/Videos/"
}

func findVideosInPath(root string, vids []data.VideoDir) []data.VideoDir {
    log.Println("looking for videos in: " + root)
    dirMap := make(map[string][]data.VideoInfo)
    filepath.Walk(root,
        func(pathStr string, info os.FileInfo, err error) error {
        if(data.IsStreamableVideoFormat(path.Ext(pathStr))) {
            log.Printf("found: %s\n", pathStr)
            name := path.Base(pathStr)
            dir := filepath.Dir(pathStr)
            dirMap[dir] = append(dirMap[dir], data.VideoInfo{pathStr, name})
        }
        return err
    })

    for key, value := range dirMap {
        vids = append(vids, data.VideoDir{key, value})
    }

    sort.Sort(data.VideoDirectories(vids))
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

/*FindAllVideos Walks throug the specified directories in the provided jsonFile
* looking for videos. if the jsonFile string is empty or an invalid path then
* only the /home/$USER/Videos directory will be searched.
* If the jsonFile string is an invalid path a waring will be logged.
* All the found videos are grouped by the container directories using the
* VideoDir struct.
* A slice of all the directories containing videos is returned. The slice is
* sorted using the directory path.
*/
func FindAllVideos(jsonFile string) []data.VideoDir {
    pathsToWalk := getPathsToWalk(jsonFile)

    vids := make([]data.VideoDir, 0, 0)
    for _, pathStr := range pathsToWalk {
        vids = findVideosInPath(pathStr, vids)
    }
    return vids
}
