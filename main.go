package main

import (
	"fmt"

	"github.com/kompere/kompere-api/requests"
)

func main() {
	res, err := requests.Get("https://shoprite.ng", &requests.Config{Http: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Body)
}
