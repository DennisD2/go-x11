// cgo -godefs defs.go

package x11

/*
#include <stdint.h>
#include <stdlib.h>

#include <X11/Xlib.h>
*/
import "C"

type XmScaleCallbackStruct struct {
	Reason    int32
	Event     *C.XEvent
	Value     int32
	Pad_cgo_0 [4]byte
}
