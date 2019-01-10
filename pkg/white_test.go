package pkg

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	const c = "c"
	if c != fmt.Sprint(&MyString{c}) {
		t.Error("wrong")
	}
}
