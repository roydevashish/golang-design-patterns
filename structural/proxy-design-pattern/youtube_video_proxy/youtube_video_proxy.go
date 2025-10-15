package youtubevideoproxy

import "github.com/roydevashish/golang-design-patterns/proxy/package_yt/youtube"

type YouTubeVideoProxy struct {
	videoId      string
	youtubeVideo *youtube.YouTubeVideo
}

func NewYouTubeVideoProxy(videoId string) *YouTubeVideoProxy {
	return &YouTubeVideoProxy{
		videoId: videoId,
	}
}

func (y *YouTubeVideoProxy) GetVideoId() string {
	return y.videoId
}

func (y *YouTubeVideoProxy) Render() {
	if(y.youtubeVideo == nil) {
		y.youtubeVideo = youtube.NewYouTubeVideo(y.videoId)
	}

	y.youtubeVideo.Render()
}
