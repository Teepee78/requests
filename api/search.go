package api

import (
	"fmt"
	"net/http"

	"github.com/kompere/kompere-api/cmd"
)

func search(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{} = r.Context().Value(cmd.Body).(map[string]interface{})
	search, searchExists := body["search"]
	if !searchExists {
		http.Error(w, "Search string not defined", http.StatusBadRequest)
		return
	}

	fmt.Println(search)
}
