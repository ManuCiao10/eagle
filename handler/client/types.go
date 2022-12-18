package client

import (
	"io"
	"net/http"

	sessionjar "github.com/juju/persistent-cookiejar"
)

type Client struct {
	client         *http.Client
	jar            *sessionjar.Jar
	LatestResponse *Response
}

type Request struct {
	client *Client

	method, url, host string

	header http.Header

	body io.Reader

	cookies []*http.Cookie
}

type Response struct {
	headers http.Header

	body []byte

	status     string
	statusCode int
	cookies    []*http.Cookie
}
