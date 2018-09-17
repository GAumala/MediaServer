package data

/*VideoDir type that represents a directory with Videos.
* DirPath is an absolute path of the directory.
* Videos is an array of VideoInfo structs. Each entry represents a video that
* has the directory specified in DirPath as its direct parent directory.
 */
type VideoDir struct {
	DirPath string
	Videos  []VideoInfo
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

// NewVideoDir is a public constructor for VideoDir
func NewVideoDir(path string, Videos []VideoInfo) VideoDir {
	return VideoDir{path, Videos}
}
