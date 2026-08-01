package x11

import "C"

// #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
// #cgo LDFLAGS: -lXt -lXm -lX11
// #include <stdlib.h>
// #include <X11/Intrinsic.h>
// #include <Xm/Xm.h>
// #include <Xm/Label.h>
// #include "hello.h"
import "C"
import (
	"unsafe"
)

// ============================================================================
// Type Wrappers - Exported Go wrappers for all C types
// ============================================================================

// Widget wraps C.Widget
type Widget struct {
	w C.Widget
}

// AppContext wraps C.XtAppContext
type AppContext struct {
	ctx C.XtAppContext
}

// OptionDescList wraps C.XrmOptionDescList
type OptionDescList struct {
	opts C.XrmOptionDescList
}

// ArgList wraps C.ArgList (kept for AppInitialize compatibility)
type ArgList struct {
	args C.ArgList
}

// XmString wraps C.XmString (Motif compound string)
type XmString struct {
	s C.XmString
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

	c_fallbackResources := C.XtNewString(C.CString(""))
	shell := C.XtAppInitialize(&appContext.ctx, c_appClass, options.opts, cnumOptions,
		&cargc, cargv, &c_fallbackResources, (*C.struct___0)(unsafe.Pointer(wargs.args)), cnumWargs)

	return Widget{w: shell}
}

// CreateManagedWidget creates a managed Xt widget
func CreateManagedWidget(name string, widgetClass WidgetClass, parent Widget, args unsafe.Pointer, num_args int) Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cnum_args = C.uint(num_args)

	widget := C.XtCreateManagedWidget(c_name, (*C.struct__WidgetClassRec)(unsafe.Pointer(widgetClass)),
		parent.w, (*C.struct___0)(args), cnum_args)

	return Widget{w: widget}
}

// RealizeWidget makes a widget visible
func RealizeWidget(w Widget) {
	C.XtRealizeWidget(w.w)
}

// AppMainLoop enters the Xt event loop
func AppMainLoop(ctx *AppContext) {
	C.XtAppMainLoop(ctx.ctx)
}

// ============================================================================
// Xm Functions - Motif Toolkit
// ============================================================================

// XmStringCreateLtoR creates a Motif compound string from left-to-right text
func XmStringCreateLtoR(text string) XmString {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	xmstr := C.XmStringCreateLtoR(c_text, C.XmFONTLIST_DEFAULT_TAG)
	return XmString{s: xmstr}
}

// XmStringFree frees an XmString
func XmStringFree(s XmString) {
	C.XmStringFree(s.s)
}

// ============================================================================
// Widget Class Accessors
// ============================================================================

// LabelWidgetClass returns the Xm Label widget class
func LabelWidgetClass() WidgetClass {
	return WidgetClass(unsafe.Pointer(C.xmLabelWidgetClass))
}

// ============================================================================
// Arg List Helpers
// ============================================================================

// NewArgListLabelString allocates an Arg array for XmNlabelString with an XmString value
func NewArgListLabelString(x XmString) (unsafe.Pointer, int) {
	var n C.uint = 1
	size := C.size_t(unsafe.Sizeof(C.Arg{})) * C.size_t(n)
	p := C.malloc(size)
	arg0 := (*C.Arg)(p)
	arg0.name = C.XmNlabelString
	arg0.value = *(*C.XtArgVal)(unsafe.Pointer(&x.s))
	return p, int(n)
}

// FreeArgList frees memory allocated by NewArgListLabelString
func FreeArgList(p unsafe.Pointer, count int) {
	if p == nil {
		return
	}
	C.free(p)
}

// ============================================================================
// Utility Functions
// ============================================================================

// Hello calls the C hello1() function (for testing)
func Hello() {
	C.hello1()
}
