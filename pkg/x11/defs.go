//go:build ignore

//go:generate sh -c "go tool cgo -godefs defs.go > x11_types.go"

package x11

/*
#include <X11/Xlib.h>
#include <X11/Intrinsic.h>
#include <Xm/Xm.h>
*/
import "C"

// Xlib
type XAnyEvent C.XAnyEvent
type XKeyEvent C.XKeyEvent
type XButtonEvent C.XButtonEvent
type XMotionEvent C.XMotionEvent
type XGCValues C.XGCValues
type XArc C.XArc
type XPoint C.XPoint
type XCharStruct C.XCharStruct

// Xt
type XtResource C.XtResource
type XtResourceList C.XtResourceList

// Motif
type XmScaleCallbackStruct C.XmScaleCallbackStruct

// We have these as manually defined Go types, for better usability
// type OptionDescRec C.XrmOptionDescRec
// type Arg C.Arg
// type ArgList C.ArgList
// type XtActionsRec C.XtActionsRec
