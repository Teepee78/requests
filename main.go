package main

import (
	"fmt"

	"github.com/kompere/kompere-api/requests"
)

func main() {
	request := &requests.Request{
		BaseUrl: "https://vendstash.mitochronhub.com",
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	res, err := request.GetAsync("/orders", &requests.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	done := make(chan bool)
	go func() {
		response := <-res
		fmt.Println(response.Body)
		done <- true
	}()

	sync, err := request.Get("", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sync.Body)

	<-done
}
