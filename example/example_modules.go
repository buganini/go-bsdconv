package main

import (
	"bsdconv"
	"fmt"
)

func main() {
	r1 := bsdconv.Module_check(bsdconv.FROM, "_utf-8")
	fmt.Println(r1)

	r2 := bsdconv.Module_check(bsdconv.INTER, "_utf-8")
	fmt.Println(r2)

	fmt.Println("Filter:")
	list_filter := bsdconv.Modules_list(bsdconv.FILTER)
	fmt.Println(list_filter)

	fmt.Println("From:")
	list_from := bsdconv.Modules_list(bsdconv.FROM)
	fmt.Println(list_from)

	fmt.Println("Inter:")
	list_inter := bsdconv.Modules_list(bsdconv.INTER)
	fmt.Println(list_inter)

	fmt.Println("To:")
	list_to := bsdconv.Modules_list(bsdconv.TO)
	fmt.Println(list_to)
}
