package filesys

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/GAumala/MediaServer/data"
)

func findVideosInPath(verbose bool, root string) []data.VideoDir {
	if verbose {
		log.Println("looking for videos in: " + root)
	}

	dirMap := make(map[string][]data.VideoInfo)

	filepath.Walk(root,
		func(pathStr string, info os.FileInfo, err error) error {
			if data.IsStreamableVideoFormat(path.Ext(pathStr)) {
				if verbose {
					log.Printf("found: %s\n", pathStr)
				}
				name := path.Base(pathStr)
				dir := filepath.Dir(pathStr)
				dirMap[dir] = append(dirMap[dir], data.VideoInfo{pathStr, name})
			}
			return err
		})

	vids := []data.VideoDir{}
	for dirPath, videos := range dirMap {
		vids = append(vids, data.NewVideoDir(dirPath, videos))
	}
	return vids
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
func FindAllVideos(config data.Config) data.VideoDirectories {
	vids := make([]data.VideoDir, 0, 0)
	for _, pathStr := range config.VideoDirs {
		vids = append(vids, findVideosInPath(config.Verbose, pathStr)...)
	}
	return vids
}
