package x11

/*
#include <stdint.h>
#include <stdlib.h>
#include <sys/types.h>

// Xt includes
#include <X11/Intrinsic.h>

*/
import "C"
import "unsafe"

// ============================================================================
// Xt definitions
// ============================================================================
// Wrapper types
type XtAppContext unsafe.Pointer
type Widget unsafe.Pointer
type WidgetList []Widget
type WidgetClass unsafe.Pointer
type XtArgVal unsafe.Pointer //  Xt generic argument value
type CAddr unsafe.Pointer
type XtPointer unsafe.Pointer
type XtTranslations unsafe.Pointer
type Pixel uint32

// OptionDescRec Go structure
type OptionDescRec struct {
	Option    string
	Specifier string
	Kind      int
	Value     string
}

// Arg Go structure
type Arg struct {
	Name  string
	Value uintptr
}

// ArgList Go structure
type ArgList struct {
	Slice []Arg
	Size  int
}

type XtActionsRec struct {
	ActionString string
	Action       func(w Widget, event *XEvent, params []string)
}

type XtEventHandler func(w Widget, clientData CAddr, event *XEvent)

type GoXtResource struct {
	Name  string
	Class string
	Rtype string
	//Size  uint32 not used in Go structure
	//Offset       uint32 not used in Go structure
	Default_type string
	Default_addr unsafe.Pointer
}

type XrmOptionDescRec struct {
	Option    string
	Specifier string
	Kind      int
	Value     string
}
