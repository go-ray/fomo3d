package gateway

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-ray/fomo3d/conf"
)

func HttpGet(url string) ([]byte, error) {
	return Request("get", url, nil)
}

func HttpPost(url string, data []byte) ([]byte, error) {
	return Request("post", url, data)
}

func Request(method, url string, data []byte) ([]byte, error) {
	var resp *http.Response
	var err error
	switch strings.ToUpper(method) {
	case "POST":
		resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	default:
		resp, err = http.Get(url)
	}
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

func fomoHost() string {
	return conf.Cfg.FomoApi.Host
}
