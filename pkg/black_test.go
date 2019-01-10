package pkg_test

import (
	"fmt"
	"testing"

	. "github.com/komod/play-go/pkg"
)

func Test(t *testing.T) {
	const c = "c"
	s := &MyString{}
	s.SetContent(c)
	if c != fmt.Sprint(s) {
		t.Error("wrong")
	}
}
