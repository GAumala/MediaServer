package data
// VideoInfo type that describes a video to stream
type VideoInfo struct {
	FilePath string
    FileName string
}

var videoFormats = [2]string{".mp4", ".webm"}

func IsStreamableVideoFormat(videoExtension string) bool {
    for _, format := range videoFormats {
       if format == videoExtension {
           return true
       }
   }
   return false
}
