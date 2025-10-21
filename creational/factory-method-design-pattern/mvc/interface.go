package mvc

type ViewEngine interface {
	Render(filename string, data map[string]string) string
}
