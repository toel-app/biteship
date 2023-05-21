package biteship

type ClientOption interface {
	Apply(client *Client)
}

// WithSecret example: biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....
func WithSecret(secret string) ClientOption {
	return withSecret{secret: secret}
}

func WithUrl(url string) ClientOption {
	return withBaseUrl{url: url}
}

func WithHttpRequest(httpRequest IHttpRequest) ClientOption {
	return withHttpRequest{HttpRequest: httpRequest}
}

type withHttpRequest struct {
	HttpRequest IHttpRequest
}

func (w withHttpRequest) Apply(client *Client) {
	client.HttpRequest = w.HttpRequest
}

type withBaseUrl struct {
	url string
}

func (w withBaseUrl) Apply(client *Client) {
	client.BiteshipUrl = w.url
}

type withSecret struct {
	secret string
}

func (w withSecret) Apply(client *Client) {
	client.SecretKey = w.secret
}
