package xm

import "C"

// #cgo LDFLAGS: -lXm
// #include <stdlib.h>
// #include <Xm/Xm.h>
import "C"
import "unsafe"

type XmString struct {
	XmString C.XmString
}

func StringCreateLtoR(text string, tag unsafe.Pointer) XmString {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	xmstr := C.XmStringCreateLtoR(c_text, (C.XmStringTag)(tag))

	r := new(XmString)
	r.XmString = xmstr
	return *r
}

func StringFree(xmstr XmString) {
	C.XmStringFree(xmstr.XmString)
}
