package api

import (
	"testing"
	"net/url"
)

func TestGetURL(t *testing.T) {
	if getURL("/", url.Values{}) != "https://app.asana.com/" {
		t.Error("built URL is Invalid")
    }

	params := url.Values{}
	params.Add("hoge", "1")
	if getURL("/wei", params) != "https://app.asana.com/wei?hoge=1" {
		t.Error("built URL is Invalid")
    }
}
