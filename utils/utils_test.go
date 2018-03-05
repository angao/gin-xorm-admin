package utils

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	str := RandomString(32)
	if len(str) == 32 {
		t.Log("pass")
	} else {
		t.Error("no pass")
	}
}
