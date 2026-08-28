//go:build ignore

//go:generate sh -c "go tool cgo -godefs defs.go > x11_types.go"

package x11

/*
#include <X11/Xlib.h>
#include <X11/Intrinsic.h>
*/
import "C"

type XAnyEvent C.XAnyEvent
type XKeyEvent C.XKeyEvent
type XButtonEvent C.XButtonEvent
type XMotionEvent C.XMotionEvent
type XtResource C.XtResource
type XtResourceList C.XtResourceList
type XGCValues C.XGCValues
type XArc C.XArc

// We have these as manually defined Go types, for better usability
// type OptionDescRec C.XrmOptionDescRec
// type Arg C.Arg
// type ArgList C.ArgList
// type XtActionsRec C.XtActionsRec
