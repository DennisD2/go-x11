package x11

// #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
// #include <stdlib.h>
// #include <X11/Intrinsic.h>
// #include <Xm/Xm.h>
// #include <Xm/Label.h>
// #include "wrapperInfo.h"
import "C"

import (
	"unsafe"
)

// ============================================================================
// Type Wrappers - Exported Go wrappers for all C types
// ============================================================================

// Widget wraps C.Widget
type Widget struct {
	Widget C.Widget
}

// AppContext wraps C.XtAppContext
type AppContext struct {
	AppContext C.XtAppContext
}

// OptionDescList wraps C.XrmOptionDescList
type OptionDescList struct {
	OptionDescList C.XrmOptionDescList
}

// ArgList wraps C.ArgList (kept for AppInitialize compatibility)
type ArgList struct {
	ArgList C.ArgList
}

// XmString wraps C.XmString (Motif compound string)
type XmString struct {
	XmString C.XmString
}

// XtArgVal wraps C.XtArgVal (generic argument value)
type XtArgVal struct {
	ArgVal C.XtArgVal
}

// WidgetClass is an opaque handle to a widget class
type WidgetClass unsafe.Pointer

// ============================================================================
// Xt Functions - X Toolkit
// ============================================================================

// AppInitialize initializes an Xt application
func AppInitialize(appContext *AppContext, appClass string, options OptionDescList, numOptions int,
	argc *int, argv []string, fallbackResources string, wargs *ArgList, numWargs int) Widget {

	c_appClass := C.CString(appClass)
	defer C.free(unsafe.Pointer(c_appClass))

	var cnumOptions = C.uint(numOptions)
	var cargc = C.int(*argc)
	var cnumWargs = C.uint(numWargs)

	var cargv **C.char
	if len(argv) > 0 {
		cstrs := make([]*C.char, len(argv))
		for i, s := range argv {
			cstrs[i] = C.CString(s)
		}
		defer func() {
			for _, s := range cstrs {
				C.free(unsafe.Pointer(s))
			}
		}()
		cargv = &cstrs[0]
	} else {
		cargv = nil
	}

	var c_fallbackResources = C.XtNewString(C.CString(""))
	shell := C.XtAppInitialize(&appContext.AppContext, c_appClass, options.OptionDescList, cnumOptions,
		&cargc, cargv, &c_fallbackResources, (*C.struct___0)(unsafe.Pointer(wargs.ArgList)), cnumWargs)

	return Widget{Widget: shell}
}

// CreateManagedWidget creates a managed Xt widget
func CreateManagedWidget(name string, widgetClass WidgetClass, parent Widget, args ArgList, num_args int) Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cnum_args = C.uint(num_args)

	var cargs *C.struct___0
	if args.ArgList != nil {
		cargs = (*C.struct___0)(unsafe.Pointer(args.ArgList))
	} else {
		cargs = nil
	}

	widget := C.XtCreateManagedWidget(c_name, (*C.struct__WidgetClassRec)(unsafe.Pointer(widgetClass)),
		parent.Widget, cargs, cnum_args)

	return Widget{Widget: widget}
}

// RealizeWidget makes a widget visible
func RealizeWidget(w Widget) {
	C.XtRealizeWidget(w.Widget)
}

// AppMainLoop enters the Xt event loop
func AppMainLoop(ctx *AppContext) {
	C.XtAppMainLoop(ctx.AppContext)
}

// ============================================================================
// Xm Functions - Motif Toolkit
// ============================================================================

// XmStringCreateLtoR creates a Motif compound string from left-to-right text
func XmStringCreateLtoR(text string) XmString {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreateLtoR(c_text, C.XmFONTLIST_DEFAULT_TAG)
	return XmString{XmString: xmstr}
}

// XmStringFree frees an XmString
func XmStringFree(s XmString) {
	C.XmStringFree(s.XmString)
}

// XtArgValFromXmString converts an XmString into an XtArgVal wrapper
func XtArgValFromXmString(x XmString) XtArgVal {
	return XtArgVal{ArgVal: *(*C.XtArgVal)(unsafe.Pointer(&x.XmString))}
}

// XtArgValFromInt converts a Go int into an XtArgVal wrapper.
func XtArgValFromInt(i int) XtArgVal {
	ci := C.long(i)
	return XtArgVal{ArgVal: *(*C.XtArgVal)(unsafe.Pointer(&ci))}
}

// XtArgValFromString converts a Go string into an XtArgVal that holds a C string pointer.
// The returned value owns the C string memory; free it with XtArgValFreeString when done.
func XtArgValFromString(s string) XtArgVal {
	cs := C.CString(s)
	return XtArgVal{ArgVal: *(*C.XtArgVal)(unsafe.Pointer(&cs))}
}

// XtArgValFreeString frees a C string previously created by XtArgValFromString.
func XtArgValFreeString(v XtArgVal) {
	if v.ArgVal == 0 {
		return
	}
	C.free(unsafe.Pointer(uintptr(v.ArgVal)))
}

// LabelWidgetClass returns the Xm Label widget class
func LabelWidgetClass() WidgetClass {
	return WidgetClass(unsafe.Pointer(C.xmLabelWidgetClass))
}

type XmN_StringID int

const (
	XmNlabelString = 1
)

var XmN_Strings = map[XmN_StringID]*C.char{
	XmNlabelString: C.XmNlabelString,
	// many more strings will come here in future
}

// AddArgListLabelString allocates an Arg array with a single (name, XmString) pair.
// Returns pointer to the Arg array and the count (1). Caller must call FreeArgList.
// AddArgListLabelString allocates an Arg array with a single (name, XmString) pair.
// Returns an ArgList object containing the allocated C array and the count (1).
// Caller must call FreeArgList when done.
// AddArgListLabelString allocates an Arg array with a single (name, XmString) pair.
// Returns an ArgList object containing the allocated C array and the count (1).
// Caller must call FreeArgList when done.
func AddArgListLabelString(nameId XmN_StringID, value XtArgVal) (ArgList, int) {
	return AppendArgList(ArgList{ArgList: nil}, 0, nameId, value)
}

// AppendArgList appends a (name, value) pair to an existing ArgList.
// If argList.ArgList is nil and count==0 this behaves like AddArgListLabelString.
// Returns the updated ArgList and the new count. Caller must call FreeArgList when done.
func AppendArgList(argList ArgList, count int, nameId XmN_StringID, value XtArgVal) (ArgList, int) {
	newCount := count + 1
	argSize := unsafe.Sizeof(C.Arg{})
	total := C.size_t(uintptr(argSize) * uintptr(newCount))
	var p unsafe.Pointer
	if argList.ArgList == nil {
		p = C.malloc(total)
	} else {
		p = C.realloc(unsafe.Pointer(argList.ArgList), total)
	}
	if p == nil {
		// allocation failed; return original
		return argList, count
	}
	// compute pointer to the new element
	elemPtr := unsafe.Pointer(uintptr(p) + uintptr(count)*argSize)
	carg := (*C.Arg)(elemPtr)
	carg.name = XmN_Strings[nameId]
	carg.value = value.ArgVal
	return ArgList{ArgList: C.ArgList(p)}, newCount
}

// ============================================================================
// Arg List Helpers
// ============================================================================

// FreeArgList frees memory allocated by AddArgListLabelString
func FreeArgList(argList ArgList, count int) {
	if argList.ArgList == nil {
		return
	}
	C.free(unsafe.Pointer(argList.ArgList))
}

// ============================================================================
// Utility Functions
// ============================================================================

// Hello calls the C hello1() function (for testing)
func WrapperInfo() {
	C.wrapperInfo()
}
