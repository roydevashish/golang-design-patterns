package videolist

import "github.com/roydevashish/golang-design-patterns/proxy/package_yt/youtube"

type VideoList struct {
	videoList map[string]youtube.Video
}

func NewVideoList() *VideoList {
	return &VideoList{
		videoList: make(map[string]youtube.Video),
	}
}

func (v *VideoList) Add(video youtube.Video) {
	v.videoList[video.GetVideoId()] = video
}

func (v *VideoList) Watch(videoId string) {
	video := v.videoList[videoId]
	video.Render()
}
