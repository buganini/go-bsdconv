/*
 * Copyright (c) 2013-2014 Kuan-Chung Chiu <buganini@gmail.com>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */
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
import "syscall"

const (
	IBUFLEN = 8192
	FILTER = C.FILTER
	FROM = C.FROM
	INTER = C.INTER
	TO = C.TO
	CTL_ATTACH_SCORE = C.BSDCONV_CTL_ATTACH_SCORE
	CTL_ATTACH_OUTPUT_FILE = C.BSDCONV_CTL_ATTACH_OUTPUT_FILE
	CTL_AMBIGUOUS_PAD = C.BSDCONV_CTL_AMBIGUOUS_PAD
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

func (this Bsdconv) Init() {
	C.bsdconv_init(this.ins);
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

func (this Bsdconv) Conv_file(ifile string, ofile string) {
	ins := this.ins

	inf := C.fopen(C.CString(ifile), C.CString("r"))
	if(inf==nil) {
		return
	}
	t := C.strdup(C.CString(ofile+".XXXXXX"))
	fd := C.mkstemp(t)
	if(fd == -1) {
		C.fclose(inf)
		C.free(unsafe.Pointer(t))
		return
	}
	otf := C.fdopen(fd, C.CString("wb+"))
	tempfile := C.GoString(t)
	C.free(unsafe.Pointer(t))

	var stat syscall.Stat_t
	syscall.Fstat(int(C.fileno(inf)), &stat)
	syscall.Fchown(int(C.fileno(otf)), int(stat.Uid), int(stat.Gid))
	syscall.Fchmod(int(C.fileno(otf)), stat.Mode)

	C.bsdconv_init(ins)
	for ins.flush==0 {
		in := C.bsdconv_malloc(IBUFLEN)
		ins.input.data = in
		ins.input.len = C.fread(in, 1, IBUFLEN, inf)
		ins.input.flags |= C.F_FREE
		ins.input.next = nil
		if(ins.input.len == 0){
			ins.flush = 1
		}
		ins.output_mode = C.BSDCONV_FILE
		ins.output.data = unsafe.Pointer(otf)
		C.bsdconv(ins)
	}

	C.fclose(inf)
	C.fclose(otf)
	syscall.Unlink(ofile)
	syscall.Rename(tempfile, ofile)
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

func Insert_phase(conversion string, codec string, phase_type int, phasen int)(string) {
	s := C.bsdconv_insert_phase(C.CString(conversion), C.CString(codec), C.int(phase_type), C.int(phasen));
	defer C.bsdconv_free(unsafe.Pointer(s))
	return C.GoString(s);
}

func Insert_codec(conversion string, codec string, phasen int, codecn int)(string) {
	s := C.bsdconv_insert_codec(C.CString(conversion), C.CString(codec), C.int(phasen), C.int(codecn));
	defer C.bsdconv_free(unsafe.Pointer(s))
	return C.GoString(s);
}

func Replace_phase(conversion string, codec string, phase_type int, phasen int)(string) {
	s := C.bsdconv_replace_phase(C.CString(conversion), C.CString(codec), C.int(phase_type), C.int(phasen));
	defer C.bsdconv_free(unsafe.Pointer(s))
	return C.GoString(s);
}

func Replace_codec(conversion string, codec string, phasen int, codecn int)(string) {
	s := C.bsdconv_replace_codec(C.CString(conversion), C.CString(codec), C.int(phasen), C.int(codecn));
	defer C.bsdconv_free(unsafe.Pointer(s))
	return C.GoString(s);
}

func Module_check(t int, c string)(bool) {
	r := C.bsdconv_module_check(C.int(t), C.CString(c))
	return uint(r) != 0
}

func Codec_check(t int, c string)(bool) {
	return Module_check(t, c);
}

func Modules_list(t int)([]string) {
	p := C.bsdconv_modules_list(C.int(t))
	defer C.bsdconv_free(unsafe.Pointer(p));
	ret := []string{}
	for *p != nil {
		ret = append(ret, C.GoString(*p))
		C.bsdconv_free(unsafe.Pointer(*p))
		p = (**_Ctype_char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(*p)))
	}
	return ret
}

func Codecs_list(t int)([]string) {
	return Modules_list(t);
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
