package main

import (
	"os"
	"syscall"
	"unsafe"
	"bsdconv"
	"fmt"
)

func main() {
	fp, fn := bsdconv.Mktemp("score.XXXXXX")
	syscall.Unlink(fn)
	clist := bsdconv.Fopen("characters_list.txt","w+")

	p:=bsdconv.Create("utf-8:score-train:null")
	p.Ctl(bsdconv.CTL_ATTACH_SCORE, unsafe.Pointer(fp), 0)
	p.Ctl(bsdconv.CTL_ATTACH_OUTPUT_FILE, unsafe.Pointer(clist), 0)

	p.Init()
	buf := make([]byte, 100)
	f, _ := os.Open(os.Args[1])
	count, _ := f.Read(buf)
	for count > 0 {
		p.Conv_chunk(buf[0:count])
		count, _ = f.Read(buf)
	}
	p.Conv_chunk_last(nil)

	bsdconv.Fclose(fp)
	p.Destroy()

	fmt.Println("bsdconv utf-32be:utf-8 characters_list.txt")
}
