package x11

import (
	"sync"
	"unsafe"
)

/*
#include <stdint.h>
#include <stdlib.h>

// Xt includes
#include <X11/Intrinsic.h>
#include <X11/Composite.h>

#include "cheader.h"
#include "wrapperInfo.h"

// Diese Hilfsfunktion befüllt die C-Struktur direkt in C
static void set_c_action_entry(XtActionsRec *table, int index, const char *name, XtActionProc proc) {
    table[index].string = (String)name;
    table[index].proc = proc;
}

// Deklaration der globalen Go-Brücke
extern void goActionBridge(Widget w, XEvent* event, String* params, Cardinal* num_params);


*/
import "C"

// ============================================================================
// Xt definitions
// ============================================================================
// Wrapper types
type XtAppContext struct{ ctx C.XtAppContext }
type Widget struct{ w C.Widget }
type WidgetList []*Widget
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
	Size  int
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
	return uintptr(unsafe.Pointer(xmstr.s))
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

func XtSetValues(w Widget, args *ArgList) {
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

	C.call_XtSetValues(w.w, cArgsList, numArgs)
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

// ============================================================================
// Xt actions code
// ============================================================================

type XtActionsRec struct {
	ActionString string
	Action       func(w Widget, event XEvent, params []string)
}

type XtTranslations struct{ t C.XtTranslations }

// Globale Registry, damit die Brücke weiß, welche Go-Funktion zu welchem Namen gehört
var globalActionMap = make(map[string]func(w Widget, event XEvent, params []string))

//export goActionBridge
func goActionBridge(w C.Widget, event *C.XEvent, params *C.String, num_params *C.Cardinal) {
	// 1. Parameter aus C in ein Go-String-Slice umwandeln
	goParams := []string{}
	if num_params != nil && *num_params > 0 {
		slice := (*[1 << 20]C.String)(unsafe.Pointer(params))[:*num_params:*num_params]
		for _, cStr := range slice {
			goParams = append(goParams, C.GoString(cStr))
		}
	}

	goWidget := Widget{w: w}
	goEvent := XEvent{*event}

	// 3. Ausführen der registrierten Funktion mit allen Parametern
	// (Hier beispielhaft für "quit" – für dynamische Zuordnungen siehe vorherige Schritte)
	// TODO IS NOT GENERIC !!!
	if goFunc, exists := globalActionMap["bye"]; exists && goFunc != nil {
		goFunc(goWidget, goEvent, goParams)
	}
}

/*func XtAppAddActions(appContext XtAppContext, actionsTable []XtActionsRec) {
	var numActions = C.Cardinal(len(actionsTable))
	C.XtAppAddActions(appContext.ctx, cActionsTable, numActions)
}*/

func XtAppAddActions(appContext XtAppContext, actionsTable []XtActionsRec) {
	numActions := len(actionsTable)
	if numActions == 0 {
		return
	}

	// 1. Erzeuge ein echtes C-Array im Speicher, das groß genug ist
	cActionsTable := make([]C.XtActionsRec, numActions)

	// 2. Befülle das C-Array mit den Daten aus dem Go-Slice
	for i, goAction := range actionsTable {
		// Registriere die Go-Funktion in unserer Map unter ihrem Namen
		globalActionMap[goAction.ActionString] = goAction.Action

		// Erzeuge einen C-String für den Action-Namen
		cName := C.CString(goAction.ActionString)
		// Hinweis: Kein defer C.free(unsafe.Pointer(cName)), da Xt diesen String im Speicher behält!

		// Befülle den C-Eintrag über die C-Hilfsfunktion
		C.set_c_action_entry(
			&cActionsTable[0], // Zeiger auf den Anfang des C-Arrays
			C.int(i),
			cName,
			C.XtActionProc(unsafe.Pointer(C.goActionBridge)), // Alle zeigen auf dieselbe Brücke
		)
	}

	// 3. Übergabe an das originale X-Toolkit (Jetzt ist cActionsTable definiert!)
	C.XtAppAddActions(
		appContext.ctx,
		&cActionsTable[0], // Zeiger auf das erste Element für C-Kompatibilität
		C.Cardinal(numActions),
	)
}

func XtParseTranslationTable(table []string) XtTranslations {
	goStr := ""
	for t := range table {
		goStr += table[t]
	}
	cTableString := C.CString(goStr)
	return XtTranslations{t: C.XtParseTranslationTable(cTableString)}
}

func XtAugmentTranslations(w Widget, translations XtTranslations) {
	C.XtAugmentTranslations(w.w, translations.t)
}

// ============================================================================
// functions,functions,functions
// ============================================================================

func XtManageChildren(widgetList WidgetList) {
	if len(widgetList) == 0 {
		return
	}

	cArray := make([]C.Widget, len(widgetList))
	for i, widget := range widgetList {
		cArray[i] = widget.w
	}

	cPtr := (**C.Widget)(unsafe.Pointer(&cArray[0]))

	// call Xt function via wrapper
	C.call_XtManageChildren(cPtr, C.int(len(widgetList)))
}
