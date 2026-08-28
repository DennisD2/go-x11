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
type Window uint64
type Drawable uint64
type GC unsafe.Pointer
type XEvent unsafe.Pointer
type XPointer unsafe.Pointer

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

func XCreateGC(display *Display, d Drawable, valuemask int, values *XGCValues) GC {
	//ret := C.XCreateGC((*C.Display)(unsafe.Pointer(display)), C.Drawable(unsafe.Pointer(d)), valuemask, values)
	ret := C.XCreateGC(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.ulong(valuemask),
		(*C.XGCValues)(unsafe.Pointer(values)),
	)
	return GC(unsafe.Pointer(ret))
}

func XDrawRectangle(display *Display, d Drawable, gc GC, x int, y int, width uint, height uint) {
	C.XDrawRectangle(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
	)
}

func XFillRectangle(display *Display, d Drawable, gc GC, x int, y int, width uint, height uint) {
	C.XFillRectangle(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
	)
}

func XSetForeground(display *Display, gc GC, pixel Pixel) {
	C.XSetForeground((*C.Display)(unsafe.Pointer(display)),
		C.GC(unsafe.Pointer(gc)),
		C.ulong(pixel))
}

func BlackPixelOfScreen(screen *Screen) Pixel {
	//return C.BlackPixelOfScreen((*C.Screen)(unsafe.Pointer(screen)))
	var cs *C.Screen = (*C.Screen)(unsafe.Pointer(screen))
	return Pixel(cs.black_pixel)
}

func WhitePixelOfScreen(screen *Screen) Pixel {
	//return C.WhitePixelOfScreen((*C.Screen)(unsafe.Pointer(screen)))
	var cs *C.Screen = (*C.Screen)(unsafe.Pointer(screen))
	return Pixel(cs.white_pixel)
}
