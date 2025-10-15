package youtube

import "fmt"

type YouTubeVideo struct {
	videoId string
}

func NewYouTubeVideo(videoId string) *YouTubeVideo {
	yt := &YouTubeVideo{videoId: videoId}
	yt.Download()
	return yt
}

func (y *YouTubeVideo) Download() {
	fmt.Println("downloading video with id", y.videoId, "from youtube api")
}

func (y *YouTubeVideo) Render() {
	fmt.Println("rendering video", y.videoId)
}

func (y *YouTubeVideo) GetVideoId() string {
	return y.videoId
}
