package main

import (
	"os"
	"fmt"
	"bsdconv"
)

func main() {
	c:=bsdconv.Create(os.Args[1])
	c.Init()
	buf := make([]byte, 100)
	inf := os.Stdin
	count, _ := inf.Read(buf)
	for count > 0 {
		fmt.Print(string(c.Conv_chunk(buf[0:count])))
		count, _ = inf.Read(buf)
	}
	fmt.Print(string(c.Conv_chunk_last(nil)))
	fmt.Println("====================================")
	fmt.Println(c.Counter(nil))
	fmt.Println(c)
	c.Destroy()
}
