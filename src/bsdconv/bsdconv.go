package bsdconv

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbsdconv
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <bsdconv.h>
*/
import "C"
import "unsafe"
import "strings"

const (
	FROM = C.FROM
	INTER = C.INTER
	TO = C.TO
	CTL_ATTACH_SCORE = C.BSDCONV_ATTACH_SCORE
	CTL_SET_WIDE_AMBI = C.BSDCONV_SET_WIDE_AMBI
	CTL_SET_TRIM_WIDTH = C.BSDCONV_SET_TRIM_WIDTH
	CTL_ATTACH_OUTPUT_FILE = C.BSDCONV_ATTACH_OUTPUT_FILE
	CTL_AMBIGUOUS_PAD = C.BSDCONV_AMBIGUOUS_PAD
)

type Bsdconv struct {
	ins *_Ctype_struct_bsdconv_instance
}

func Create(s string)(*Bsdconv) {
	conv := C.CString(s)
	ins := C.bsdconv_create(conv)
	C.free(unsafe.Pointer(conv))
	if ins == nil {
		return nil
	}
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

func (this Bsdconv) Ctl(ctl int, p unsafe.Pointer, v int) {
	ins := this.ins
	C.bsdconv_ctl(ins, C.int(ctl), p, C.int(v))
}

func Codec_check(t int, c string)(bool) {
	r := C.bsdconv_codec_check(C.int(t), C.CString(c))
	return uint(r) != 0
}

func Codecs_list(t int)([]string) {
	p := C.bsdconv_codecs_list(C.int(t))
	defer C.bsdconv_free(unsafe.Pointer(p));
	ret := []string{}
	for *p != nil {
		ret = append(ret, C.GoString(*p))
		C.bsdconv_free(unsafe.Pointer(*p))
		p = (**_Ctype_char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(*p)))
	}
	return ret
}

func Mktemp(template string)(*C.FILE, string) {
	t := C.strdup(C.CString(template))
	fd := C.mkstemp(t)
	fp := C.fdopen(fd, C.CString("wb+"))
	fn := C.GoString(t)
	C.free(unsafe.Pointer(t))
	return fp, fn
}

func Fopen(p string, m string)(*C.FILE) {
	f, _ := C.fopen(C.CString(p), C.CString(m))
	return f
}

func Fclose(fp *C.FILE) {
	C.fclose(fp)
}
