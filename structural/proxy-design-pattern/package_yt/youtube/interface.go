package youtube

type Video interface {
	Render()
	GetVideoId() string
}
