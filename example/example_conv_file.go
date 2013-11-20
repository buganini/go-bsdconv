package main

//go run example_conv_file.go utf-8:full:utf-8 in.txt out.txt

import (
	"os"
	"fmt"
	"bsdconv"
)

func main() {
	c:=bsdconv.Create(os.Args[1])
	if c==nil {
		os.Exit(1)
	}

	c.Conv_file(os.Args[2], os.Args[3])

	fmt.Println("====================================")
	fmt.Println(c.Counter(nil))
	fmt.Println(c)
	c.Destroy()
}
