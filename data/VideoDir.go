package data
// VideoDir type that represents a directory with videos
type VideoDir struct {
	DirPath string
    Videos []VideoInfo
}

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
