package bsdconv

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbsdconv
#include <stdio.h>
#include <bsdconv.h>
*/
import "C"
import "unsafe"
import "strings"

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

func (this Bsdconv) String()(string) {
	ins := this.ins
	str := C.bsdconv_pack(ins)
	s := []string{"bsdconv.Create(\"", C.GoString(str), "\")"};
	C.bsdconv_free(unsafe.Pointer(str))
	return strings.Join(s, "")
}

func (this Bsdconv) Conv_chunk(b []byte)([]byte) {
	ins := this.ins
	ins.output_mode=C.BSDCONV_AUTOMALLOC;
	ins.input.data=unsafe.Pointer(&b[0])
	ins.input.len=C.size_t(len(b))
	ins.input.flags=0
	ins.input.next=nil
	C.bsdconv(ins)
	ret:=C.GoBytes(unsafe.Pointer(ins.output.data), C.int(ins.output.len))
	C.bsdconv_free(unsafe.Pointer(ins.output.data))
	return ret
}

func (this Bsdconv) Init() {
	C.bsdconv_init(this.ins);
}

func (this Bsdconv) Conv_chunk_last(b []byte)([]byte) {
	ins := this.ins
	ins.output_mode=C.BSDCONV_AUTOMALLOC;
	if len(b) > 0 {
		ins.input.data=unsafe.Pointer(&b[0])
	}
	ins.input.len=C.size_t(len(b))
	ins.input.flags=0
	ins.input.next=nil
	ins.flush=1
	C.bsdconv(ins)
	ret:=C.GoBytes(unsafe.Pointer(ins.output.data), C.int(ins.output.len))
	C.bsdconv_free(unsafe.Pointer(ins.output.data))
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

func (this Bsdconv) Counter(ct interface{})(interface{}) {
	ins := this.ins
	if(ct==nil){
		ret := map[string] uint {}
		p := ins.counter
		for p != nil {
			ret[C.GoString(p.key)] = uint(p.val)
			p = p.next
		}
		return ret
	}
	v := C.bsdconv_counter(ins, C.CString(ct.(string)))
	return uint(*v)
}
