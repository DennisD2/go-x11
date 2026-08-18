package x11

// Below is simngle place where we have CFLAGS and LDFLAGS located.
/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
#include <stdint.h>
#include <stdlib.h>

// Xt includes
#include <X11/Xlib.h>
*/
import "C"

// ============================================================================
// X definitions
// ============================================================================
// Wrapper types
type Display struct{ d *C.Display }
type Screen struct{ s *C.Screen }
type Window struct{ w C.Window }
type XEvent struct{ e C.XEvent }
