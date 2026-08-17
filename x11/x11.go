package x11

import "C"
import (
	"sync"
	"unsafe"
)

/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
#include <stdlib.h>
#include <X11/Intrinsic.h>
#include <Xm/Xm.h>
#include <Xm/Label.h>
#include <Xm/PushB.h>

#include "cheader.h"
#include "wrapperInfo.h"

*/
import "C"

var XmNlabelString = C.GoString(C.XmNlabelString)
var XtNWidth = C.GoString(C.XtNwidth)
var XtNHeight = C.GoString(C.XtNheight)

// LabelWidgetClass returns the Xm Label widget class
func LabelWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmLabelWidgetClass}
}

func PushButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPushButtonWidgetClass}
}

// Wrapper-Typen
type AppContext struct{ ctx C.XtAppContext }
type Widget struct{ w C.Widget }

type WidgetClass struct{ c C.WidgetClass }

type OptionDescRec struct {
	Option    string
	Specifier string
	Kind      int
	Value     string
}

type Arg struct {
	Name  string
	Value uintptr
}

// ArgList wraps C.ArgList
type ArgList struct {
	Slice []Arg
	Size  int // Ermöglicht den Aufruf 'args.Size' in main
}

// XtArgVal wraps C.XtArgVal (generic argument value)
type XtArgVal struct {
	ArgVal C.XtArgVal
}

// XmString wraps C.XmString (Motif compound string)
type XmString struct {
	XmString C.XmString
}

func WrapperInfo() {
	C.wrapperInfo()
}

func XtAppInitialize(
	appContext *AppContext,
	appClass string,
	options []OptionDescRec,
	goArgs []string,
	fallbackResources []string,
	args []Arg,
) Widget {

	// 1. Anwendungsklasse
	cAppClass := C.CString(appClass)
	defer C.free(unsafe.Pointer(cAppClass))

	// 2. Options-Liste (C-Speicher)
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

	// 3. Argument-Vektor (argc / argv) -> Geändert auf C.malloc wegen cgo-Pointer-Rules
	argc := C.int(len(goArgs))
	// Wir allozieren das char** Array direkt im C-Speicher
	cArgsPtr := C.malloc(C.size_t(len(goArgs)+1) * C.size_t(unsafe.Sizeof(uintptr(0))))
	defer C.free(cArgsPtr)

	// Cast auf ein bearbeitbares Go-Slice aus C-Pointern
	cArgsArray := (*[1 << 20]*C.char)(cArgsPtr)[: len(goArgs)+1 : len(goArgs)+1]
	for i, arg := range goArgs {
		cArgsArray[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgsArray[i]))
	}
	cArgsArray[len(goArgs)] = nil // NULL-Terminierung am Ende

	// 4. Fallback Resources -> Geändert auf C.malloc wegen cgo-Pointer-Rules
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

	// 5. ArgList
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

	// 6. Aufruf über unsere C-Hilfsfunktion (Jetzt absolut sicher vor GC-Verschiebungen)
	cWidget := C.call_XtAppInitialize(
		&appContext.ctx,
		cAppClass,
		cOptions, numOptions,
		&argc,
		cArgsPtr,   // Übergabe des reinen C-Speicher-Pointers
		cFallbacks, // Übergabe des reinen C-Speicher-Pointers
		cArgsList, numArgs,
	)

	return Widget{w: cWidget}
}

// CreateManagedWidget creates a managed Xt widget
func CreateManagedWidget(name string, widgetClass WidgetClass, parent Widget, args *ArgList) Widget {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cArgsList unsafe.Pointer
	var numArgs C.Cardinal = 0

	// Wenn args nicht nil ist und Elemente enthält, konvertieren wir es für C
	if args != nil && len(args.Slice) > 0 {
		numArgs = C.Cardinal(len(args.Slice))
		cArgsSlice := C.malloc(C.size_t(len(args.Slice)) * C.sizeof_Arg)
		defer C.free(cArgsSlice)

		// Cast auf das temporäre C-Array
		cArgsArray := (*[1 << 20]C.Arg)(cArgsSlice)[:len(args.Slice):len(args.Slice)]
		for i, arg := range args.Slice {
			cArgsArray[i].name = C.CString(arg.Name)
			cArgsArray[i].value = C.XtArgVal(arg.Value)

			// Die C-Strings für die Namen müssen nach dem Xt-Aufruf wieder freigegeben werden
			defer C.free(unsafe.Pointer(cArgsArray[i].name))
		}
		cArgsList = cArgsSlice
	}
	// Aufruf über Ihren funktionierenden call_XtCreateManagedWidget Helper
	widget := C.call_XtCreateManagedWidget(
		cName,
		widgetClass.c,
		parent.w,
		cArgsList,
		numArgs,
	)

	return Widget{w: widget}
}

// RealizeWidget makes a widget visible
func XtRealizeWidget(w Widget) {
	C.XtRealizeWidget(w.w)
}

// AppMainLoop enters the Xt event loop
func XtAppMainLoop(ctx *AppContext) {
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

// AppendArgList fügt ein neues Argument hinzu und gibt die aktualisierte Liste zurück.
// Wenn die übergebene Liste nil ist, wird eine neue erstellt.
func AppendArgList(list *ArgList, name string, value uintptr) *ArgList {
	if list == nil {
		list = &ArgList{
			Slice: make([]Arg, 0),
		}
	}

	// Neues Argument an das interne Slice anhängen
	list.Slice = append(list.Slice, Arg{
		Name:  name,
		Value: value,
	})

	// Size-Zähler für main aktualisieren
	list.Size = len(list.Slice)

	return list
}

// ============================================================================
// callback code
// ============================================================================

var XmNactivateCallback = C.GoString(C.XmNactivateCallback)

var (
	callbackRegistry = make(map[uintptr]func())
	callbackMutex    sync.Mutex
	nextCallbackID   uintptr = 1
)

//export goCallbackDispatcher
func goCallbackDispatcher(w C.Widget, client_data C.XtPointer, call_data C.XtPointer) {
	// Wir holen uns die ID zurück, die wir als client_data übergeben haben
	callbackID := uintptr(client_data)

	callbackMutex.Lock()
	goFunc, exists := callbackRegistry[callbackID]
	callbackMutex.Unlock()

	if exists && goFunc != nil {
		// Hier wird die originale Go-Funktion aus main ausgeführt!
		goFunc()
	}
}

// XtAddCallback registriert eine Go-Funktion für ein Widget-Event
func XtAddCallback(widget Widget, callbackName string, goFunction func()) {
	cName := C.CString(callbackName)
	defer C.free(unsafe.Pointer(cName))

	// Go-Funktion registrieren und ID generieren
	callbackMutex.Lock()
	id := nextCallbackID
	callbackRegistry[id] = goFunction
	nextCallbackID++
	callbackMutex.Unlock()

	// Aufruf unseres C-Helpers mit der ID
	C.call_XtAddCallback(widget.w, cName, C.uintptr_t(id))
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
