package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// GetData from url,use method get
func GetData(url string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Close = true

	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &transport,
		Timeout:   time.Duration(10) * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// SubmitData from url,use method submit
func SubmitData(url string, data []byte, method string) ([]byte, error) {
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Close = true

	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: &transport,
		Timeout:   time.Duration(10) * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
