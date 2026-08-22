package x11

import (
	"runtime/cgo"
	"sync"
	"unsafe"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include <sys/types.h>

// Xt includes
#include <X11/Intrinsic.h>
#include <X11/Composite.h>

#include "cheader.h"
#include "wrapperInfo.h"

*/
import "C"

// ============================================================================
// Xt definitions
// ============================================================================
// Wrapper types
type XtAppContext unsafe.Pointer
type Widget unsafe.Pointer
type WidgetList []Widget
type WidgetClass struct{ c C.WidgetClass }
type XtArgVal unsafe.Pointer //  Xt generic argument value
type CAddr unsafe.Pointer

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

type XtTranslations struct{ t C.XtTranslations }

type XtPointer unsafe.Pointer

type XtEventHandler func(w Widget, clientData CAddr, event *XEvent)

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
	C.XtDispatchEvent((*C.XEvent)(unsafe.Pointer(e)))
}

func XtNextEvent(e *XEvent) {
	C.XtNextEvent((*C.XEvent)(unsafe.Pointer(e)))
}

func XtAppNextEvent(appContext XtAppContext, e *XEvent) {
	C.XtAppNextEvent((C.XtAppContext)(unsafe.Pointer(appContext)), (*C.XEvent)(unsafe.Pointer(e)))
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

	return Widget(cWidget)
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
		(*C.XtAppContext)(unsafe.Pointer(appContext)),
		cAppClass,
		cOptions, numOptions,
		&argc,
		cArgvPtr, /* XtString* argv_in_out */
		cFallbacks,
		cArgsList, /* ArgList args */
		numArgs,
	)

	return Widget(cWidget)
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
		C.Widget(parent),
		cArgsList,
		numArgs,
	)

	return Widget(widget)
}

func XtManageChild(w Widget) {
	C.XtManageChild(C.Widget(w))
}

// CreateManagedWidget creates a managed Xt widget
func XtCreateManagedWidget(name string, widgetClass WidgetClass, parent Widget, args *ArgList) Widget {
	w := XtCreateWidget(name, widgetClass, parent, args)
	C.XtManageChild(C.Widget(w))
	return w
}

// RealizeWidget makes a widget visible
func XtRealizeWidget(w Widget) {
	C.XtRealizeWidget(C.Widget(w))
}

func XtIsRealized(w Widget) bool {
	return C.XtIsRealized(C.Widget(w)) == C.TRUE
}

func XtIsManaged(w Widget) bool {
	return C.XtIsManaged(C.Widget(w)) == C.TRUE
}

func XtDestroyWidget(w Widget) {
	C.XtDestroyWidget(C.Widget(w))
}

// not sure that this works as expected - test
func XtDisplay(w Widget) *Display {
	cd := C.XtDisplay(C.Widget(w))
	return (*Display)(unsafe.Pointer(cd))
}

func XtScreen(w Widget) *Screen {
	cs := C.XtScreen(C.Widget(w))
	return (*Screen)(unsafe.Pointer(cs))
}

// returns object, not reference
func XtWindow(w Widget) Window {
	cw := C.XtWindow(C.Widget(w))
	return Window(unsafe.Pointer(&cw))
}

// AppMainLoop enters the Xt event loop
func XtAppMainLoop(ctx XtAppContext) {
	C.XtAppMainLoop((C.XtAppContext)(ctx))
}

func XtArgValFromXmString(xmstr XmString) uintptr {
	return uintptr(unsafe.Pointer(xmstr))
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
	if v == nil {
		return
	}
	C.free(unsafe.Pointer(uintptr(v)))
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

	C.call_XtSetValues(C.Widget(w), cArgsList, numArgs)
}

func XtGetValues(w Widget, args *ArgList) {
	if args == nil || args.Size <= 0 || len(args.Slice) == 0 {
		return
	}

	// allocate C array where X will put requested data
	cArgs := make([]C.Arg, args.Size)

	// free later allocated Go strings
	allocatedStrings := make([]*C.char, args.Size)
	defer func() {
		for _, cStr := range allocatedStrings {
			if cStr != nil {
				C.free(unsafe.Pointer(cStr))
			}
		}
	}()

	// 2convert go -> c
	for i := 0; i < args.Size; i++ {
		cStr := C.CString(args.Slice[i].Name)
		allocatedStrings[i] = cStr

		cArgs[i].name = cStr
		// Das Feld .value im C-Arg erwartet ein XtArgVal (oft synonym zu long/uintptr)
		cArgs[i].value = C.XtArgVal(args.Slice[i].Value)
	}

	// call c function
	C.XtGetValues(
		C.Widget(w),
		(*C.Arg)(unsafe.Pointer(&cArgs[0])),
		C.Cardinal(args.Size),
	)

	// copy back values into Go arry
	for i := 0; i < args.Size; i++ {
		args.Slice[i].Value = uintptr(cArgs[i].value)
	}
}

// ============================================================================
// Xt callback code
// ============================================================================
type CallbackInfo struct {
	Func       func(w Widget, clientData XtPointer, callData XtPointer)
	Widget     Widget
	ClientData XtPointer
}

var (
	callbackRegistry = make(map[uintptr]CallbackInfo)
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
	info, exists := callbackRegistry[callbackID]
	callbackMutex.Unlock()

	// Das aktuelle, dynamische call_data von C in den Go-Typ verpacken
	currentCallData := XtPointer(unsafe.Pointer(call_data))

	if exists {
		info.Func(info.Widget, info.ClientData, currentCallData)
	}
}

// XtAddCallback registers a Go-Funktion for a Widget-Event
func XtAddCallback(widget Widget, callbackName string,
	goFunction func(w Widget, clientData1 XtPointer, callData XtPointer), clientData XtPointer) {
	cName := C.CString(callbackName)
	defer C.free(unsafe.Pointer(cName))

	// register Go-Funktion and create ID
	callbackMutex.Lock()
	id := nextCallbackID
	callbackRegistry[id] = CallbackInfo{
		Func:       goFunction,
		Widget:     widget,
		ClientData: clientData,
	}
	nextCallbackID++
	callbackMutex.Unlock()

	// call Xt function via wrapper
	C.call_XtAddCallback(C.Widget(widget), cName, C.uintptr_t(id))
}

// ============================================================================
// Xt actions code
// ============================================================================

// registry maps an ID to a go function
var (
	actionPool      = make(map[int]func(w Widget, event *XEvent, params []string))
	nextActionID    = 0
	actionPoolMutex sync.RWMutex
)

//export goActionBridgeWithId
func goActionBridgeWithId(w C.Widget, event *C.XEvent, params *C.String, num_params *C.Cardinal, actionId C.int) {
	// 1Try to get go function for actionId
	actionPoolMutex.RLock()
	goFunc, exists := actionPool[int(actionId)]
	actionPoolMutex.RUnlock()

	if !exists || goFunc == nil {
		return
	}

	// Convert C parameters to Go types
	goParams := []string{}
	if num_params != nil && *num_params > 0 {
		slice := (*[1 << 20]C.String)(unsafe.Pointer(params))[:*num_params:*num_params]
		for _, cStr := range slice {
			goParams = append(goParams, C.GoString(cStr))
		}
	}
	goWidget := Widget(w)
	goEvent := (*XEvent)(unsafe.Pointer(event))

	// call the go function
	goFunc(goWidget, goEvent, goParams)
}

func XtAppAddActions(appContext XtAppContext, actionsTable []XtActionsRec) {
	numActions := len(actionsTable)
	if numActions == 0 {
		return
	}

	cActionsTable := make([]C.XtActionsRec, numActions)

	actionPoolMutex.Lock()
	for i, goAction := range actionsTable {
		// generate new ID for mapping
		id := nextActionID
		nextActionID++

		// save the action functio (a go function) with its ID
		actionPool[id] = goAction.Action

		cName := C.CString(goAction.ActionString)

		// get one of pool-contained C functions, mapped to this ID
		cBridgeProc := C.get_bridge_ptr(C.int(id))

		// set all values in C ActionsTable structure
		C.set_c_action_entry(&cActionsTable[0], C.int(i), cName, cBridgeProc)
	}
	actionPoolMutex.Unlock()
	// with valid C values, call C function
	C.XtAppAddActions((C.XtAppContext)(unsafe.Pointer(appContext)), &cActionsTable[0], C.Cardinal(numActions))
}

func XtParseTranslationTable(table []string) XtTranslations {
	goStr := ""
	for t := range table {
		goStr += table[t]
		goStr += "\n"
	}
	cTableString := C.CString(goStr)
	return XtTranslations{t: C.XtParseTranslationTable(cTableString)}
}

func XtAugmentTranslations(w Widget, translations XtTranslations) {
	C.XtAugmentTranslations(C.Widget(w), translations.t)
}

// ============================================================================
// XtEvent handler code
// ============================================================================

//export goEventHandlerBridge
func goEventHandlerBridge(w C.Widget, client_data C.XtPointer, event *C.XEvent, continue_to_dispatch *C.Boolean) {
	if client_data == nil {
		return
	}

	// 1. Wandle das client_data zurück in das Go cgo.Handle um
	handlePtr := (*cgo.Handle)(unsafe.Pointer(client_data))

	// 2. Extrahiere die hinterlegte Go-Funktion
	goFunc, ok := handlePtr.Value().(XtEventHandler)
	if !ok || goFunc == nil {
		return
	}

	// 3. Verpacke die C-Typen in deine Framework-Typen
	goWidget := Widget(w)
	goEvent := (*XEvent)(unsafe.Pointer(event))
	goCAddr := CAddr(client_data)

	// 4. Rufe den puren Go-Code auf!
	goFunc(goWidget, goCAddr, goEvent)
}

func XtAddEventHandler(w Widget, eventMask EventMask, nonMaskable bool, proc XtEventHandler, clientData CAddr) {
	// 1. Create handle
	handle := cgo.NewHandle(proc)
	cClientData := C.XtPointer(unsafe.Pointer(&handle))

	// convert params
	var cNonMaskable C.Boolean = C.False
	if nonMaskable {
		cNonMaskable = C.True
	}

	C.call_XtAddEventHandler(
		C.Widget(w),
		C.EventMask(eventMask),
		cNonMaskable,
		C.XtEventHandler(C.goEventHandlerBridge),
		cClientData,
	)
}

func ConvertAnyEvent(ev *XEvent) XAnyEvent {
	return *(*XAnyEvent)(unsafe.Pointer(ev))
}

func ConvertButtonEvent(ev *XEvent) XButtonEvent {
	return *(*XButtonEvent)(unsafe.Pointer(ev))
}

func ConvertKeyEvent(ev *XEvent) XKeyEvent {
	return *(*XKeyEvent)(unsafe.Pointer(ev))
}

func ConvertMotionNotifyEvent(ev *XEvent) XMotionEvent {
	return *(*XMotionEvent)(unsafe.Pointer(ev))
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
		cArray[i] = C.Widget(widget)
	}

	cPtr := (**C.Widget)(unsafe.Pointer(&cArray[0]))

	// call Xt function via wrapper
	C.call_XtManageChildren(cPtr, C.int(len(widgetList)))
}
