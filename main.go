package main

import (
	"fmt"

	"github.com/kompere/kompere-api/requests"
)

func main() {
	request := &requests.Request{
		BaseUrl: "https://webhook.site/5146b497-397b-4675-81b3-8b100345418a",
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	res, err := request.Delete("", map[string]string{"success": "true"}, &requests.Config{Http: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Body)
}
