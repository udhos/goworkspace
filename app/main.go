package main

import (
	"fmt"

	"github.com/udhos/goworkspace-lib/lib"
)

func main() {
	modify("fooBAR")
}

func modify(s string) {
	result := lib.Modify(s)
	fmt.Printf(`lib.Modify("%s") => "%s"`+"\n", s, result)
}
