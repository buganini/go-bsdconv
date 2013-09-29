package bsdconv

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbsdconv
#include <stdio.h>
#include <bsdconv.h>
*/
import "C"
import "unsafe"

func Create(s string)(*_Ctype_struct_bsdconv_instance) {
	conv := C.CString(s)
	ins := C.bsdconv_create(conv)
	C.free(unsafe.Pointer(conv))
	return ins
}


func Conv(ins *_Ctype_struct_bsdconv_instance, b []byte)([]byte) {
	C.bsdconv_init(ins);
	ins.output_mode=C.BSDCONV_AUTOMALLOC;
	ins.input.data=unsafe.Pointer(&b[0])
	ins.input.len=C.size_t(len(b))
	ins.input.flags=0
	ins.input.next=nil
	ins.flush=1
	C.bsdconv(ins)
	ret:=C.GoBytes(unsafe.Pointer(ins.output.data), C.int(ins.output.len))
	C.bsdconv_free(unsafe.Pointer(ins.output.data))
	return ret
}

func Destroy(ins *_Ctype_struct_bsdconv_instance) {
	C.bsdconv_destroy(ins)
}
