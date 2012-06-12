package main

import (
	"fmt"
)

var funcs = []func() string{
	func() string { return "main" },
}

func main() {
	for _, f := range funcs {
		fmt.Println(f())
	}
}
