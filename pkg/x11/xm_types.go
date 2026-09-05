// cgo -godefs defs.go

package x11

/*
#include <stdint.h>
#include <stdlib.h>

#include <X11/Xlib.h>
#include <Xm/Xm.h>
*/
import "C"
import "unsafe"

// ============================================================================
// Xm definitions
// ============================================================================

type XmString unsafe.Pointer
type XmStringCharSet unsafe.Pointer
type XmFontContext unsafe.Pointer
type XmFontList unsafe.Pointer

type XmAnyCallbackStruct struct {
	reason int
	event  *XEvent
}

type Orientation int

const (
	XmNO_ORIENTATION Orientation = iota // 0
	XmVERTICAL                          // 1
	XmHORIZONTAL                        // 2
)

type Attachment int

const (
	XmATTACH_NONE            Attachment = iota // 0
	XmATTACH_FORM                              // 1
	XmATTACH_OPPOSITE_FORM                     // 2
	XmATTACH_WIDGET                            // 3
	XmATTACH_OPPOSITE_WIDGET                   // 4
	XmATTACH_POSITION                          // 5
	XmATTACH_SELF                              // 6
)

type XmScaleCallbackStruct struct {
	Reason    int32
	Event     *C.XEvent
	Value     int32
	Pad_cgo_0 [4]byte
}
