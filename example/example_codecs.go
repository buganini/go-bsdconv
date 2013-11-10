package main

import (
	"bsdconv"
	"fmt"
)

func main() {
	r1 := bsdconv.Codec_check(bsdconv.FROM, "_utf-8")
	fmt.Println(r1)

	r2 := bsdconv.Codec_check(bsdconv.INTER, "_utf-8")
	fmt.Println(r2)

	fmt.Println("From:")
	list_from := bsdconv.Codecs_list(bsdconv.FROM)
	fmt.Println(list_from)

	fmt.Println("Inter:")
	list_inter := bsdconv.Codecs_list(bsdconv.INTER)
	fmt.Println(list_inter)

	fmt.Println("To:")
	list_to := bsdconv.Codecs_list(bsdconv.TO)
	fmt.Println(list_to)
}
