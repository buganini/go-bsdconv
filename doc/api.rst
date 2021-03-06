Go binding API
==================

.. go:package:: bsdconv

.. go:const:: FILTER
.. go:const:: FROM
.. go:const:: INTER
.. go:const:: TO

	Phase type

.. go:const:: CTL_ATTACH_SCORE
.. go:const:: CTL_ATTACH_OUTPUT_FILE
.. go:const:: CTL_AMBIGUOUS_PAD

	Request for :go:func:`(bsdconv.Bsdconv) Ctl`

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

.. go:function:: func (this Bsdconv) Ctl(request int, ptr unsafe.Pointer, val int)

	Manipulate the underlying codec parameters

.. go:function:: func Insert_phase(conversion string, codec string, phase_type int, phasen int)(string)

	Insert conversion phase into bsdconv conversion string

.. go:function:: func Insert_codec(conversion string, codec string, phasen int, codecn int)(string)

	Insert conversion codec into bsdconv conversion string

.. go:function:: func Replace_phase(conversion string, codec string, phase_type int, phasen int)(string)

	Replace conversion phase in the bsdconv conversion string

.. go:function:: func Replace_codec(conversion string, codec string, phasen int, codecn int)(string)

	Replace conversion codec in the bsdconv conversion string

.. go:function:: func Codec_check(type int, module string)(bool)

	DEPRECATED: Use :go:func:`Module_check` instead

.. go:function:: func Module_check(type int, module string)(bool)

	Check availability with given type and module name

.. go:function:: func Codecs_list(type int)([]string)

	DEPRECATED: Use :go:func:`Modules_list` instead

.. go:function:: func Modules_list(type int)([]string)

	Get modules list of specified type

.. go:function:: func Mktemp(template string)(*C.FILE, string)

	mkstemp()

.. go:function:: func Fopen(p string, m string)(*C.FILE)

	fopen()

.. go:function:: func Fclose(fp *C.FILE)

	fclose()
