package x11

/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
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
	XmString C.XmString
}

// ============================================================================
// Xm Functions - Motif Toolkit
// ============================================================================

// XmStringCreateLtoR creates a Motif compound string from left-to-right text
func XmStringCreateLtoR(text string) XmString {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreateLtoR(c_text, C.XmFONTLIST_DEFAULT_TAG)
	return XmString{XmString: xmstr}
}

// XmStringFree frees an XmString
func XmStringFree(s XmString) {
	C.XmStringFree(s.XmString)
}
