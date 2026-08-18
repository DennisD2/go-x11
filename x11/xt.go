package x11

import "C"
import (
	"sync"
	"unsafe"
)

/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
#include <stdint.h>
#include <stdlib.h>

// Xt includes
#include <X11/Intrinsic.h>

#include "cheader.h"
#include "wrapperInfo.h"

*/
import "C"

// ============================================================================
// Xt definitions
// ============================================================================
// Wrapper types
type XtAppContext struct{ ctx C.XtAppContext }
type Widget struct{ w C.Widget }
type WidgetClass struct{ c C.WidgetClass }

// XtArgVal wraps C.XtArgVal (generic argument value)
type XtArgVal struct {
	ArgVal C.XtArgVal
}

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
	Size  int // Ermöglicht den Aufruf 'args.Size' in main
}

// ============================================================================
// Wrapper own definitions
// ============================================================================
func WrapperInfo() {
	C.wrapperInfo()
}

// ============================================================================
// Xt functions
// ============================================================================

func XtDispatchEvent(e *XEvent) {
	cep := &(e.e)
	C.XtDispatchEvent(cep)
}

func XtNextEvent(e *XEvent) {
	cep := &(e.e)
	C.XtNextEvent(cep)
}

func XtAppNextEvent(appContext XtAppContext, e *XEvent) {
	cep := &(e.e)
	C.XtAppNextEvent(appContext.ctx, cep)
}

func XtInitialize(
	shellName string,
	appClass string,
	options []OptionDescRec,
	argv []string,
) Widget {
	cAppClass := C.CString(appClass)
	defer C.free(unsafe.Pointer(cAppClass))
	cShellName := C.CString(shellName)
	defer C.free(unsafe.Pointer(cShellName))

	var cOptions unsafe.Pointer
	numOptions := C.Cardinal(len(options))
	if len(options) > 0 {
		cOptions = C.malloc(C.size_t(len(options)) * C.sizeof_XrmOptionDescRec)
		defer C.free(cOptions)

		for i, opt := range options {
			cOpt := C.CString(opt.Option)
			cSpec := C.CString(opt.Specifier)
			cVal := C.CString(opt.Value)

			C.set_option_rec(cOptions, C.int(i), cOpt, cSpec, C.int(opt.Kind), cVal)

			defer C.free(unsafe.Pointer(cOpt))
			defer C.free(unsafe.Pointer(cSpec))
			defer C.free(unsafe.Pointer(cVal))
		}
	}

	// handle argc+argv
	argc := C.int(len(argv))
	// Wir allozieren das char** Array direkt im C-Speicher
	cArgvPtr := C.malloc(C.size_t(len(argv)+1) * C.size_t(unsafe.Sizeof(uintptr(0))))
	defer C.free(cArgvPtr)

	// Cast auf ein bearbeitbares Go-Slice aus C-Pointern
	cArgsArray := (*[1 << 20]*C.char)(cArgvPtr)[: len(argv)+1 : len(argv)+1]
	for i, arg := range argv {
		cArgsArray[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgsArray[i]))
	}
	cArgsArray[len(argv)] = nil // NULL-Terminierung am Ende

	// call Xt function via wrapper
	cWidget := C.call_XtInitialize(
		cShellName,
		cAppClass,
		cOptions, numOptions,
		&argc,
		cArgvPtr,
	)

	return Widget{w: cWidget}
}

func XtAppInitialize(
	appContext *XtAppContext,
	appClass string,
	options []OptionDescRec,
	argv []string,
	fallbackResources []string,
	args []Arg,
) Widget {

	// handle appClass
	cAppClass := C.CString(appClass)
	defer C.free(unsafe.Pointer(cAppClass))

	// handle options
	var cOptions unsafe.Pointer
	numOptions := C.Cardinal(len(options))
	if len(options) > 0 {
		cOptions = C.malloc(C.size_t(len(options)) * C.sizeof_XrmOptionDescRec)
		defer C.free(cOptions)

		for i, opt := range options {
			cOpt := C.CString(opt.Option)
			cSpec := C.CString(opt.Specifier)
			cVal := C.CString(opt.Value)

			C.set_option_rec(cOptions, C.int(i), cOpt, cSpec, C.int(opt.Kind), cVal)

			defer C.free(unsafe.Pointer(cOpt))
			defer C.free(unsafe.Pointer(cSpec))
			defer C.free(unsafe.Pointer(cVal))
		}
	}

	// handle argc+argv
	argc := C.int(len(argv))
	// Wir allozieren das char** Array direkt im C-Speicher
	cArgvPtr := C.malloc(C.size_t(len(argv)+1) * C.size_t(unsafe.Sizeof(uintptr(0))))
	defer C.free(cArgvPtr)

	// Cast auf ein bearbeitbares Go-Slice aus C-Pointern
	cArgsArray := (*[1 << 20]*C.char)(cArgvPtr)[: len(argv)+1 : len(argv)+1]
	for i, arg := range argv {
		cArgsArray[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgsArray[i]))
	}
	cArgsArray[len(argv)] = nil // NULL-Terminierung am Ende

	// handle fallbackResources
	var cFallbacks unsafe.Pointer
	if len(fallbackResources) > 0 {
		cFallbacks = C.malloc(C.size_t(len(fallbackResources)+1) * C.size_t(unsafe.Sizeof(uintptr(0))))
		defer C.free(cFallbacks)

		cFBArray := (*[1 << 20]*C.char)(cFallbacks)[: len(fallbackResources)+1 : len(fallbackResources)+1]
		for i, fb := range fallbackResources {
			cFBArray[i] = C.CString(fb)
			defer C.free(unsafe.Pointer(cFBArray[i]))
		}
		cFBArray[len(fallbackResources)] = nil // NULL-Terminierung
	}

	// handle ArgList
	var cArgsList unsafe.Pointer
	numArgs := C.Cardinal(len(args))
	if len(args) > 0 {
		cArgsSlice := C.malloc(C.size_t(len(args)) * C.sizeof_Arg)
		defer C.free(cArgsSlice)

		cArgsArray := (*[1 << 20]C.Arg)(cArgsSlice)[:len(args):len(args)]
		for i, arg := range args {
			cArgsArray[i].name = C.CString(arg.Name)
			cArgsArray[i].value = C.XtArgVal(arg.Value)
			defer C.free(unsafe.Pointer(cArgsArray[i].name))
		}
		cArgsList = cArgsSlice
	}

	// call Xt function via wrapper
	cWidget := C.call_XtAppInitialize(
		&appContext.ctx,
		cAppClass,
		cOptions, numOptions,
		&argc,
		cArgvPtr, /* XtString* argv_in_out */
		cFallbacks,
		cArgsList, /* ArgList args */
		numArgs,
	)

	return Widget{w: cWidget}
}

func XtCreateWidget(name string, widgetClass WidgetClass, parent Widget, args *ArgList) Widget {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cArgsList unsafe.Pointer
	var numArgs C.Cardinal = 0

	// convert args
	if args != nil && len(args.Slice) > 0 {
		numArgs = C.Cardinal(len(args.Slice))
		cArgsSlice := C.malloc(C.size_t(len(args.Slice)) * C.sizeof_Arg)
		defer C.free(cArgsSlice)

		// casting to C array
		cArgsArray := (*[1 << 20]C.Arg)(cArgsSlice)[:len(args.Slice):len(args.Slice)]
		for i, arg := range args.Slice {
			cArgsArray[i].name = C.CString(arg.Name)
			cArgsArray[i].value = C.XtArgVal(arg.Value)

			//  free these strings
			defer C.free(unsafe.Pointer(cArgsArray[i].name))
		}
		cArgsList = cArgsSlice
	}
	// call Xt function via wrapper
	widget := C.call_XtCreateManagedWidget(
		cName,
		widgetClass.c,
		parent.w,
		cArgsList,
		numArgs,
	)

	return Widget{w: widget}
}

func XtManageChild(w Widget) {
	C.XtManageChild(w.w)
}

// CreateManagedWidget creates a managed Xt widget
func XtCreateManagedWidget(name string, widgetClass WidgetClass, parent Widget, args *ArgList) Widget {
	w := XtCreateWidget(name, widgetClass, parent, args)
	C.XtManageChild(w.w)
	return w
}

// RealizeWidget makes a widget visible
func XtRealizeWidget(w Widget) {
	C.XtRealizeWidget(w.w)
}

func XtIsRealized(w Widget) bool {
	return C.XtIsRealized(w.w) == C.TRUE
}

func XtIsManaged(w Widget) bool {
	return C.XtIsManaged(w.w) == C.TRUE
}

func XtDestroyWidget(w Widget) {
	C.XtDestroyWidget(w.w)
}

// not sure that this works as expected - test
func XtDisplay(w Widget) *Display {
	cd := C.XtDisplay(w.w)
	return &Display{d: cd}
}

func XtScreen(w Widget) *Screen {
	cs := C.XtScreen(w.w)
	return &Screen{s: cs}
}

// returns object, not reference
func XtWindow(w Widget) Window {
	cw := C.XtWindow(w.w)
	return Window{w: cw}
}

// AppMainLoop enters the Xt event loop
func XtAppMainLoop(ctx *XtAppContext) {
	C.XtAppMainLoop(ctx.ctx)
}

func XtArgValFromXmString(xmstr XmString) uintptr {
	// Konvertiert den C-Pointer (*C.char / XmString) in eine Go-Ganzzahl (uintptr)
	return uintptr(unsafe.Pointer(xmstr.XmString))
}

// XtArgValFromInt converts a Go int into an XtArgVal wrapper.
func XtArgValFromInt(i int) uintptr {
	ci := C.long(i)
	return uintptr(unsafe.Pointer(&ci))
}

// XtArgValFromString converts a Go string into an XtArgVal that holds a C string pointer.
// The returned value owns the C string memory; free it with XtArgValFreeString when done.
func XtArgValFromString(s string) uintptr {
	cs := C.CString(s)
	return uintptr(unsafe.Pointer(&cs))
}

// XtArgValFreeString frees a C string previously created by XtArgValFromString.
func XtArgValFreeString(v XtArgVal) {
	if v.ArgVal == 0 {
		return
	}
	C.free(unsafe.Pointer(uintptr(v.ArgVal)))
}

// AppendArgList appends a (name, value) pair to an existing ArgList.
// If argList.ArgList is nil a new ArgList is created on-the-fly.
// Returns the updated ArgList. Caller must call FreeArgList when done.
func AppendArgList(list *ArgList, name string, value uintptr) *ArgList {
	if list == nil {
		list = &ArgList{
			Slice: make([]Arg, 0),
		}
	}

	// append
	list.Slice = append(list.Slice, Arg{
		Name:  name,
		Value: value,
	})

	// update size counter
	list.Size = len(list.Slice)

	return list
}

// ============================================================================
// Xt callback code
// ============================================================================

var (
	callbackRegistry = make(map[uintptr]func())
	callbackMutex    sync.Mutex
	nextCallbackID   uintptr = 1
)

// We need to make this function callable by C, so "export" keyword
//
//export goCallbackDispatcher
func goCallbackDispatcher(w C.Widget, client_data C.XtPointer, call_data C.XtPointer) {
	// get ID
	callbackID := uintptr(client_data)

	callbackMutex.Lock()
	goFunc, exists := callbackRegistry[callbackID]
	callbackMutex.Unlock()

	if exists && goFunc != nil {
		// execute go function which serves as callback
		goFunc()
	}
}

// XtAddCallback registers a Go-Funktion for a Widget-Event
func XtAddCallback(widget Widget, callbackName string, goFunction func()) {
	cName := C.CString(callbackName)
	defer C.free(unsafe.Pointer(cName))

	// register Go-Funktion and create ID
	callbackMutex.Lock()
	id := nextCallbackID
	callbackRegistry[id] = goFunction
	nextCallbackID++
	callbackMutex.Unlock()

	// call Xt function via wrapper
	C.call_XtAddCallback(widget.w, cName, C.uintptr_t(id))
}
