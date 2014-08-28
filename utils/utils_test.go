package utils

import (
	"testing"
)

func TestHome(t *testing.T) {
	if len(Home()) == 0 {
		t.Error("func Home failed")
	}
}

