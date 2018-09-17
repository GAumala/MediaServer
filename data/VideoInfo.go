package data

/*VideoInfo type that describes a video to stream
* FilePath is the absolute path to the video file.
* FileName is the name of the video file.
* Key is a string that uniquely identifies a video file with URL friendly
* characters.
 */
type VideoInfo struct {
	FilePath string
	FileName string
	Key      uint32
}

/* VideoDict is a map whose keys are unique 32 bit integers that uniquely
 * identify videos in the filesystem.
 */
type VideoDict map[uint32]VideoInfo

var videoFormats = [2]string{".mp4", ".webm"}

/*IsStreamableVideoFormat returns true if the extension provided in videoExtension
* is an extension of video format that can be streamed over HTML5.
 */
func IsStreamableVideoFormat(videoExtension string) bool {
	for _, format := range videoFormats {
		if format == videoExtension {
			return true
		}
	}
	return false
}
