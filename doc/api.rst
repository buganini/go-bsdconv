Go binding API
==================

.. go:package:: bsdconv

.. go:type:: Bsdconv

.. go:function:: func Create(s string)(*Bsdconv)

	Create converter instance with given conversion string

.. go:function:: func (this Bsdconv) Init()

	Initialize the converter instance

.. go:function:: func (this Bsdconv) Conv_chunk(b []byte)([]byte)

	Perform chunked conversion

.. go:function:: func (this Bsdconv) Conv_chunk_last(b []byte)([]byte)

	Perform chunked conversion with flush

.. go:function:: func (this Bsdconv) Conv(b []byte)([]byte)

	Perform conversion with initialization and flush

.. go:function:: func (this Bsdconv) Conv_file(ifile string, ofile string)

	Perform conversion from ifile to ofile

.. go:function:: func (this Bsdconv) Destroy()

	Destroy the converter instance

.. go:function:: func (this Bsdconv) Counter(ct interface{})(interface{})

	Get counter or counters if ct is nil

.. go:function:: func (this Bsdconv) Ctl(ctl int, p unsafe.Pointer, v int)

	Manipulate the underlying codec parameters

.. go:function:: func Insert_phase(conversion string, codec string, phase_type int, phasen int)(string)

	Manipulate conversion string

.. go:function:: func Insert_codec(conversion string, codec string, phasen int, codecn int)(string)

	Manipulate conversion string

.. go:function:: func Replace_phase(conversion string, codec string, phase_type int, phasen int)(string)

	Manipulate conversion string

.. go:function:: func Replace_codec(conversion string, codec string, phasen int, codecn int)(string)

	Manipulate conversion string

.. go:function:: func Codec_check(t int, c string)(bool)

	Check codec availability with given phase type and codec name

.. go:function:: func Codecs_list(t int)([]string)

	Get codecs list of given phase type

.. go:function:: func Mktemp(template string)(*C.FILE, string)

	mkstemp()

.. go:function:: func Fopen(p string, m string)(*C.FILE)

	fopen()

.. go:function:: func Fclose(fp *C.FILE)

	fclose()