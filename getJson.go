package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
