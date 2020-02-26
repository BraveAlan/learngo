package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	bytes, e := Fetch("https://album.zhenai.com/u/1244648936")
	if e != nil {
		t.Errorf(e.Error())
	} else {
		fmt.Printf("%s", bytes)
	}
}
