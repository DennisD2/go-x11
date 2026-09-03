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

func XCreateGC(display *Display, d Drawable, valuemask int, values *XGCValues) GC {
	ret := C.XCreateGC(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.ulong(valuemask),
		(*C.XGCValues)(unsafe.Pointer(values)),
	)
	return GC(unsafe.Pointer(ret))
}

func XDrawRectangle(display *Display, d Drawable, gc GC, x int, y int, width uint, height uint) int {
	ret := C.XDrawRectangle(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
	)
	return int(ret)
}

func XFillRectangle(display *Display, d Drawable, gc GC, x int, y int, width uint, height uint) int {
	ret := C.XFillRectangle(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
	)
	return int(ret)
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

func XDrawLine(display *Display, d Drawable, gc GC, x1 int, y1 int, x2 uint, y2 uint) int {
	ret := C.XDrawLine(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x1),
		C.int(y1),
		C.int(x2),
		C.int(y2),
	)
	return int(ret)
}

func XDrawArc(display *Display, d Drawable, gc GC, x int, y int,
	width uint, height uint,
	angle1 int, angle2 int) int {
	ret := C.XDrawArc(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
		C.int(angle1),
		C.int(angle2),
	)
	return int(ret)
}

func XFillArc(display *Display, d Drawable, gc GC, x int, y int,
	width uint, height uint,
	angle1 int, angle2 int) int {
	ret := C.XFillArc(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
		C.int(angle1),
		C.int(angle2),
	)
	return int(ret)
}

func XDrawArcs(display *Display, d Drawable, gc GC, arcs []XArc) int {
	ret := C.XDrawArcs(
		(*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		(*C.XArc)(unsafe.Pointer(&arcs[0])),
		C.int(len(arcs)),
	)
	return int(ret)
}

func XDrawString(display *Display, d Drawable, gc GC, x int, y int, str string) int {
	ret := C.XDrawString((*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y),
		C.CString(str),
		C.int(len(str)))
	return int(ret)
}

func XDrawPoint(display *Display, d Drawable, gc GC, x int, y int) int {
	ret := C.XDrawPoint((*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		C.int(x),
		C.int(y))
	return int(ret)
}

func XDrawPoints(display *Display, d Drawable, gc GC, points []XPoint, mode int) int {
	ret := C.XDrawPoints((*C.Display)(unsafe.Pointer(display)),
		C.Drawable(d),
		C.GC(unsafe.Pointer(gc)),
		(*C.XPoint)(unsafe.Pointer(&points[0])),
		C.int(len(points)),
		C.int(mode))
	return int(ret)
}

func XQueryFont(display *Display, xid XID) *XFontStruct {
	fontStruct := C.XQueryFont((*C.Display)(unsafe.Pointer(display)), C.XID(xid))
	return (*XFontStruct)(unsafe.Pointer(fontStruct))
}

func XGContextFromGC(gc GC) GContext {
	ret := C.XGContextFromGC((C.GC)(unsafe.Pointer(gc)))
	goCtx := (*GContext)(unsafe.Pointer(&ret))
	return *goCtx
}

func XTextExtents(font *XFontStruct, text string, nchar int,
	dir *int, ascent *int, descent *int, overall *XCharStruct) {
	cdir := C.int(*dir)
	casc := C.int(*ascent)
	cdesc := C.int(*descent)

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	cFontStruct := (*C.XFontStruct)(unsafe.Pointer(font))
	cCharStruct := (*C.XCharStruct)(unsafe.Pointer(overall))

	C.XTextExtents(cFontStruct, ctext, C.int(nchar),
		&cdir, &casc, &cdesc, cCharStruct)

	*dir = int(cdir)
	*ascent = int(casc)
	*descent = int(cdesc)
}
