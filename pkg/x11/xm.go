package x11

/*
#include <stdint.h>
#include <stdlib.h>

#include <Xm/Xm.h>
*/
import "C"
import "unsafe"

// ============================================================================
// Xm definitions
// ============================================================================

type XmString unsafe.Pointer
type XmStringCharset unsafe.Pointer

type XmAnyCallbackStruct struct {
	reason int
	event  *XEvent
}

// ============================================================================
// Xm Functions - Motif Toolkit
// ============================================================================

// XmStringCreate creates a Motif compound string
func XmStringCreate(text string, charset XmStringCharset) XmString {
	c_text := C.CString(text)
	c_charSet := (*C.char)(unsafe.Pointer(charset))
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreate(c_text, c_charSet)
	return XmString(unsafe.Pointer(xmstr))
}

// XmStringCreateLtoR creates a Motif compound string from left-to-right text
func XmStringCreateLtoR(text string, charset XmStringCharset) XmString {
	c_text := C.CString(text)
	c_charSet := (*C.char)(unsafe.Pointer(charset))
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreateLtoR(c_text, c_charSet)
	return XmString(unsafe.Pointer(xmstr))
}

// XmStringFree frees an XmString
func XmStringFree(s XmString) {
	C.XmStringFree(C.XmString(unsafe.Pointer(s)))
}

func XmStringConcat(a XmString, b XmString) XmString {
	goa := C.XmString(unsafe.Pointer(a))
	gob := C.XmString(unsafe.Pointer(b))
	return XmString(C.XmStringConcat(goa, gob))
}
