package data

/*VideoInfo type that describes a video to stream
* FilePath contains the absolute path to the video file
* FileName contains the name of the file.
 */
type VideoInfo struct {
	FilePath string
	FileName string
}

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
