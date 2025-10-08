package http

type HTTPRequest struct {
	Username string
	Password string
}

func NewHTTPRequest(username, password string) *HTTPRequest {
	return &HTTPRequest{
		Username: username,
		Password: password,
	}
}
