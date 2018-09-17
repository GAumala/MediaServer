package filesys

import (
	"hash/fnv"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/GAumala/MediaServer/data"
)

func generateVideoKey(videoPath string) uint32 {
	hash := fnv.New32()
	hash.Write([]byte(videoPath))
	return hash.Sum32()
}

func findVideosInPath(verbose bool, root string,
	videoDict data.VideoDict) []data.VideoDir {
	if verbose {
		log.Println("looking for videos in: " + root)
	}

	dirDict := make(map[string][]data.VideoInfo)

	filepath.Walk(root,
		func(pathStr string, info os.FileInfo, err error) error {
			if data.IsStreamableVideoFormat(path.Ext(pathStr)) {
				name := path.Base(pathStr)
				dir := filepath.Dir(pathStr)
				key := generateVideoKey(pathStr)
				newVideo := data.VideoInfo{
					FilePath: pathStr,
					FileName: name,
					Key:      key,
				}

				dirDict[dir] = append(dirDict[dir], newVideo)
				videoDict[key] = newVideo
			}
			return err
		})

	vids := []data.VideoDir{}
	for dirPath, videos := range dirDict {
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
func FindAllVideos(config data.Config) (data.VideoDirectories, data.VideoDict) {
	dirs := make([]data.VideoDir, 0, 0)
	videoDict := make(map[uint32]data.VideoInfo)
	for _, pathStr := range config.VideoDirs {
		dirs = append(dirs, findVideosInPath(config.Verbose, pathStr, videoDict)...)
	}
	return dirs, videoDict
}
