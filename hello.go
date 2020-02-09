package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/john-black-3k/hello-go/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
