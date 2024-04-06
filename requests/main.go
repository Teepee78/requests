package requests

import (
	"net/http"
)

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

/*
Create a custom request with set config
*/
type Request struct {
	BaseUrl string
	Headers map[string]string
}
