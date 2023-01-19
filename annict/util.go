package annict

import "net/http"

type authedTransport struct {
	wrapped http.RoundTripper
	token   string
}

func NewAuthedTransport(token string, rt http.RoundTripper) *authedTransport {
	return &authedTransport{
		wrapped: rt,
		token:   token,
	}
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.wrapped.RoundTrip(req)
}
