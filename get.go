package requests

import (
	"fmt"
	"net/http"
)

func (req *Request) Get(url string, config *Config) (*Response, error) {
	// create request
	request, err := http.NewRequest(http.MethodGet, req.BaseUrl+url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	setHeaders(&request.Header, req.Headers, config.Headers)

	return sendRequest(request, config.Http)
}

func (req *Request) GetAsync(url string, config *Config) (chan *Response, error) {
	response := make(chan *Response)
	var err error

	go func() {
		res, getError := req.Get(url, config)
		response <- res
		err = getError
		close(response)
	}()

	return response, err
}
