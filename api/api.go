package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/thash/asana/config"
	"github.com/thash/asana/utils"
)

const (
	GetBase   = "https://app.asana.com"
	PostBase  = "https://app.asana.com/api/1.0"
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/36.0.1985.125 Safari/537.36"
)

type Base struct {
	Gid  string `json:"gid"`
	Name string `json:"name"`
}

func Get(path string, params url.Values) []byte {
	req, err := http.NewRequest("GET", getURL(path, params), nil)
	utils.Check(err)
	return fire(req)
}

func getURL(path string, params url.Values) string {
	if params == nil || params.Encode() == "" {
		return GetBase + path
	} else {
		return GetBase + path + "?" + params.Encode()
	}
}

func Post(path string, data string) []byte {
	req, err := http.NewRequest("POST", PostBase+path, strings.NewReader(data))
	utils.Check(err)
	return fire(req)
}

func Put(path string, data string) []byte {
	req, err := http.NewRequest("PUT", PostBase+path, strings.NewReader(data))
	utils.Check(err)
	return fire(req)
}

func fire(req *http.Request) []byte {
	client := &http.Client{}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Authorization", "Bearer " + config.Load().Personal_access_token)

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)

	utils.Check(err)

	if resp.StatusCode >= 300 {
		println(resp.Status)
	}

	return body
}
