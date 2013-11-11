package main

import (
	"bsdconv"
	"fmt"
)

func main() {
	sin := "utf-8:utf-8,ascii"
	sout := bsdconv.Insert_phase(sin, "upper", bsdconv.INTER, 1)
	fmt.Println(sout)

	sin2 := sout
	sout2 := bsdconv.Replace_phase(sin2, "full", bsdconv.INTER, 1)
	fmt.Println(sout2)

	sin3 := sout2
	sout3 := bsdconv.Replace_codec(sin3, "big5", 2, 1)
	fmt.Println(sout3)

	sin4 := sout3
	sout4 := bsdconv.Insert_codec(sin4, "ascii", 0, 1)
	fmt.Println(sout4)
}
