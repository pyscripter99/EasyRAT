package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetUrl() string {
	return fmt.Sprintf("http%s://%s:%d", func() string {
		if Https {
			return "s"
		} else {
			return ""
		}
	}(), Ip, Port)
}

func Get(path string) (*http.Response, error) {
	resp, err := http.Get(GetUrl() + path)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Post(path string, contentType string, data io.Reader) (*http.Response, error) {
	resp, err := http.Post(GetUrl()+path, contentType, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PostForm(path string, data url.Values) (*http.Response, error) {
	resp, err := http.PostForm(GetUrl()+path, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PostJSON(path string, data interface{}) (*http.Response, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(GetUrl()+path, "application/json", &buf)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func WakeServer() {
	Get("/")
}
