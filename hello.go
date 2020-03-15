package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/wuilliam321/gocode/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH yraM"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
