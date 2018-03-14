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

func TestContains(t *testing.T) {
	slice := make([]string, 5)
	slice = append(slice, "hello")
	slice = append(slice, "world")
	slice = append(slice, "golang")
	slice = append(slice, "gin")
	slice = append(slice, "xorm")

	if Contains(slice, "hello") {
		t.Log("pass")
	} else {
		t.Error("no pass")
	}

	if !Contains(slice, "hhh") {
		t.Log("pass")
	} else {
		t.Error("no pass")
	}
}
