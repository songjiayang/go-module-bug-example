package main

import (
	"fmt"
	"go.mod/b"
)

func main() {
	b.Init()
	b.Add("foo", "foo")
	fmt.Println(b.Get("foo"))
}
