package x11

/*
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
type Window uint64
type Drawable uint64
type GC unsafe.Pointer
type XEvent unsafe.Pointer
type XPointer unsafe.Pointer
type XFontStruct struct{}
type GContext unsafe.Pointer

type EventMask int64
type XID uint64

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

var CoordModeOrigin = C.CoordModeOrigin
var CoordModePrevious = C.CoordModePrevious
