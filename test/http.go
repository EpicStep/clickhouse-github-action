package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendQuery(query string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8123/?query=%s", query))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
