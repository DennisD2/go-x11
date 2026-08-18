package x11

/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
#include <stdint.h>
#include <stdlib.h>

// Xt includes
#include <X11/Intrinsic.h>
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
