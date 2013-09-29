package bsdconv

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbsdconv
#include <stdio.h>
#include <bsdconv.h>
*/
import "C"
import "unsafe"

type Bsdconv struct {
	ins *_Ctype_struct_bsdconv_instance
}

func Create(s string)(*Bsdconv) {
	conv := C.CString(s)
	ins := C.bsdconv_create(conv)
	C.free(unsafe.Pointer(conv))
	ret := new(Bsdconv)
	ret.ins = ins
	return ret
}


func (this Bsdconv) Conv(b []byte)([]byte) {
	ins := this.ins
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

func (this Bsdconv) Destroy() {
	C.bsdconv_destroy(this.ins)
}
