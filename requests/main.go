package requests

import "net/http"

var client = &http.Client{}

type Config struct {
	Http    bool
	Headers map[string]string
}

type Response struct {
	Status  int
	Body    interface{}
	Headers http.Header
	Cookies []*http.Cookie
	Request *http.Request
}
