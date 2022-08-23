package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HTTP_GET(url string) (string, error) {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	} else {
		if resp.StatusCode == 200 {
			var bodyReader io.ReadCloser = resp.Body
			body, err := ioutil.ReadAll(bodyReader)
			if err != nil {
				return "", err
			} else {
				return string(body), nil
			}
		} else {
			return "", fmt.Errorf("服务器异常")
		}
	}
}

func main() {
	url := "http://localhost:8888/test/1"
	ret, err := HTTP_GET(url)
	fmt.Println(ret, err)
}
