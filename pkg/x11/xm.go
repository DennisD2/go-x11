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

// XmString wraps C.XmString (Motif compound string)
type XmString struct {
	s C.XmString
}

type XmStringCharset struct {
	charSet string
}

// ============================================================================
// Xm Functions - Motif Toolkit
// ============================================================================

// XmStringCreate creates a Motif compound string
func XmStringCreate(text string, charset XmStringCharset) XmString {
	c_text := C.CString(text)
	c_charSet := C.CString(charset.charSet)
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreate(c_text, c_charSet)
	return XmString{s: xmstr}
}

// XmStringCreateLtoR creates a Motif compound string from left-to-right text
func XmStringCreateLtoR(text string, charset XmStringCharset) XmString {
	c_text := C.CString(text)
	c_charSet := C.CString(charset.charSet)
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreateLtoR(c_text, c_charSet)
	return XmString{s: xmstr}
}

// XmStringFree frees an XmString
func XmStringFree(s XmString) {
	C.XmStringFree(s.s)
}

func XmStringConcat(a XmString, b XmString) XmString {
	return XmString{C.XmStringConcat(a.s, b.s)}
}
