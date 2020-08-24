package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func mainy() {

	url := "localhost:8080/videos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic c3Jpbmk6cGFzc3dvcmQ=")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
