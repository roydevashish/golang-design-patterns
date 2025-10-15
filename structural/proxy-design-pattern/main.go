package main

import (
	videolist "github.com/roydevashish/golang-design-patterns/proxy/video_list"
	youtubevideoproxy "github.com/roydevashish/golang-design-patterns/proxy/youtube_video_proxy"
)

func main() {
	vl := videolist.NewVideoList()
	videoIds := []string{"12345", "abcde", "golang"}
	for _, v := range videoIds {
		vl.Add(youtubevideoproxy.NewYouTubeVideoProxy(v))
	}

	vl.Watch("abcde")
}
