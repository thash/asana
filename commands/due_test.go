package commands

import (
	"testing"
	"time"
)

func TestToDate(t *testing.T) {
	if toDate("2014-08-30") != "2014-08-30" {
		t.Error("Failed to parse date")
	}
	if toDate("today") != time.Now().Format("2006-01-02") {
		t.Error("Failed to parse date")
	}
	d, _ := time.ParseDuration("24h")
	if toDate("tomorrow") != time.Now().Add(d).Format("2006-01-02") {
		t.Error("Failed to parse date")
	}
	if toDate("hoge") != "hoge" {
		t.Error("Unparsable string should be returned as it is")
	}
}
