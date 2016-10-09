package data

import (
    "path"
    "path/filepath"
    "sort"
)
/*VideoDir type that represents a directory with videos.
* DirPath is an absolute path of the directory.
* Videos is an array of VideoInfo structs. Each entry represents a video that
* has the directory specified in DirPath as its direct parent directory.
*/
type VideoDir struct {
	DirPath string
    Videos []VideoInfo
}

//VideoDirectories is a type of an array of VideoDir. it implements the sort.Interface
type VideoDirectories []VideoDir

func (vd VideoDirectories) Len() int {
    return len(vd)
}
func (vd VideoDirectories) Swap(i, j int) {
    vd[i], vd[j] = vd[j], vd[i]
}
func (vd VideoDirectories) Less(i, j int) bool {
    return vd[i].DirPath < vd[j].DirPath
}

/*FindVideo looks for a video specified with an absolute path pathStr inside
* the VideoDirectories struct.
* First it checks if the VideoDirectories struct contains the directory of the
* video. if it succeeds the it gets the VideoDir struct representing that
* directory and checks if there is a VideoInfo struct with that absolute path.
* If the VideoInfo struct is found, it is returned, otherwise an empty struct
* is returned.
* All searches in slices are performed with the binary search algorithm of
* the Search function of the package sort
*/
func (vd VideoDirectories) FindVideo(pathStr string) VideoInfo {
	//First check if directory exists
	target := filepath.Dir(pathStr)
	length := vd.Len()
	pos := sort.Search(length, func (i int) bool { return vd[i].DirPath >= target})

	if(pos < length && vd[pos].DirPath == target) {//FOUND DIR
		//Now check if a file with that name exists
		videoDir := vd[pos]
		length = len(videoDir.Videos)
		target = path.Base(pathStr)
		pos = sort.Search(length, func(i int) bool {
			return videoDir.Videos[i].FileName >= target
		})

		if(pos < length && videoDir.Videos[pos].FileName == target) {//FOUND FILE
			return videoDir.Videos[pos]
		}
	}

	return VideoInfo{}
}
