package x11

/* Below is single place where we have CFLAGS and LDFLAGS located. */

/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
#include <stdint.h>
#include <stdlib.h>

// Xt includes
#include <X11/Xlib.h>
*/
import "C"
import "unsafe"

// ============================================================================
// X definitions
// ============================================================================
// Wrapper types
type Display unsafe.Pointer
type Screen unsafe.Pointer
type Window unsafe.Pointer
type XEvent struct{ e C.XEvent }

type EventMask int64

const (
	NoEventMask         EventMask = EventMask(C.NoEventMask)
	KeyPressMask        EventMask = EventMask(C.KeyPressMask)
	KeyReleaseMask      EventMask = EventMask(C.KeyReleaseMask)
	ButtonPressMask     EventMask = EventMask(C.ButtonPressMask)
	ButtonReleaseMask   EventMask = EventMask(C.ButtonReleaseMask)
	PointerMotionMask   EventMask = EventMask(C.PointerMotionMask)
	StructureNotifyMask EventMask = EventMask(C.StructureNotifyMask)
)

var KeyPress int32 = C.KeyPress
var KeyRelease int32 = C.KeyRelease
var ButtonPress int32 = C.ButtonPress
var ButtonRelease int32 = C.ButtonRelease
var MotionNotify int32 = C.MotionNotify
var GenericEvent int32 = C.GenericEvent
