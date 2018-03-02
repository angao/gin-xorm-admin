package utils

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	str := RandomString(8)
	if len(str) == 8 {
		t.Log("pass")
	} else {
		t.Error("no pass")
	}
}
