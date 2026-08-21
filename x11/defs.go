//go:build ignore

//go:generate sh -c "go tool cgo -godefs defs.go > x11_types.go"

package x11

/*
#include <X11/Xlib.h>
*/
import "C"

type XAnyEvent C.XAnyEvent
type XKeyEvent C.XKeyEvent
type XButtonEvent C.XButtonEvent
type XMotionEvent C.XMotionEvent
