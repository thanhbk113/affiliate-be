package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"affiliate/internal/constants"

	"github.com/labstack/echo/v4"
)

var client = &http.Client{}

// HeaderJSON ...
var HeaderJSON = map[string]string{
	echo.HeaderContentType: "application/json",
}

type RequestHeader struct {
	Authorization string `json:"authorization"`
	ApiKey        string `json:"apiKey"`
	AppName       string `json:"appName"`
}

// AssignHeader ...
func (h RequestHeader) AssignHeader(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	if h.Authorization != "" {
		r.Header.Set(constants.HeaderAuthorization, h.Authorization)
	}
	if h.ApiKey != "" {
		r.Header.Set(constants.HeaderApiKey, h.ApiKey)
	}
	if h.AppName != "" {
		r.Header.Set(constants.HeaderAppName, h.AppName)
	}
}

// RequestPostAnyHttp ...
func RequestPostAnyHttp(uri string, body []byte, header RequestHeader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header.AssignHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	by, err := ioutil.ReadAll(resp.Body)
	return by, err
}
