package api

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"../config"
	"../utils"
)

type Base struct {
	Id   int
	Name string
}

func Get(path string, params url.Values) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", buildGetUrl("https://app.asana.com", path, params), nil)
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36"
	req.Header.Set("User-Agent", ua)
	req.SetBasicAuth(config.Load().Api_key, "")
	resp, err := client.Do(req)
	utils.Check(err)

	contents, err2 := ioutil.ReadAll(resp.Body)
	utils.Check(err2)

	return contents
}

func buildGetUrl(host string, path string, params url.Values) string {
	if params == nil || params.Encode() == "" {
		return host + path
	} else {
		return host + path + "?" + params.Encode()
	}
}
