package utils

import (
	"bytes"
	"net/http"
)

func CreateHttpRequest(method string, url string, body []byte) (response *http.Response, err error) {
	client := http.Client{}
	bodyRequest := &bytes.Buffer{}
	if method == "POST" {
		bodyRequest = bytes.NewBuffer(body)
	}
	req, _ := http.NewRequest(method, url, bodyRequest)
	req.Header.Set("Content-type", "application/json")
	// req.Header.Add("Authorization", CreateAuth())
	res, err := client.Do(req)
	return res, err
}
