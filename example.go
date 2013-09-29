package main

import (
	"fmt"
	"bsdconv"
)

func main() {
	c:=bsdconv.Create("utf-8:casefold:utf-8")
	fmt.Println(string(bsdconv.Conv(c, []byte("AaЯя"))))
	bsdconv.Destroy(c)
}
