package commands

import (
	"testing"
)

func TestToDate(t *testing.T) {
	if toDate("2014-08-30") != "2014-08-30" {
		t.Error("Failed to parse date")
	}
}
