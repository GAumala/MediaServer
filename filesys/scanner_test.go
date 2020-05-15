package filesys

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GAumala/MediaServer/data"
)

func TestFindAllVideosDefault(t *testing.T) {
	c := data.DefaultConfig()
	dirs, dict := FindAllVideos(*c)
	assert.Equal(t, sort.IsSorted(dirs), true, "FindAllVideios() should return sorted directories")
	n := 0
	for _, dir := range dirs {
		n += len(dir.Videos)
		for _, v := range dir.Videos {
			assert.Equal(t, dict[v.Key], v, "VideoInfo is not equal to value in dictionay")
		}
	}
	assert.Equal(t, n, len(dict), "FindAllVideos() should return same files num")
}

func TestFindAllVideosNullDir(t *testing.T) {
	c := &data.Config{
		Port:      8080,
		Verbose:   true,
		VideoDirs: []string{},
	}
	dirs, dict := FindAllVideos(*c)
	assert.Equal(t, len(dirs), 0, "FindAllVideios(NullDir) should return null result")
	assert.Equal(t, len(dict), 0, "FindAllVideios(NullDir) should return null result")
}

func TestFindAllVideosWrongDir(t *testing.T) {
	c := &data.Config{
		Port:      8080,
		Verbose:   true,
		VideoDirs: []string{"/Nonexist/aaa/"},
	}
	dirs, dict := FindAllVideos(*c)
	assert.Equal(t, len(dirs), 0, "FindAllVideios(WrongDir) should return null result")
	assert.Equal(t, len(dict), 0, "FindAllVideios(WrongDir) should return null result")
}
