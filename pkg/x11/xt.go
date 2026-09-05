package x11

import (
	"fmt"
	"reflect"
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
	options []XrmOptionDescRec,
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
	options []XrmOptionDescRec,
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
		C.WidgetClass(widgetClass),
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
	return Window(cw)
}

func XtName(w Widget) string {
	cn := C.XtName(C.Widget(w))
	return C.GoString(cn)
}

func XtSuperclass(w Widget) WidgetClass {
	cc := C.XtSuperclass(C.Widget(w))
	return WidgetClass(cc)
}

func XtClass(w Widget) WidgetClass {
	cc := C.XtClass(C.Widget(w))
	return WidgetClass(cc)
}

func XtParent(w Widget) Widget {
	cw := C.XtParent(C.Widget(w))
	return Widget(cw)
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
	// get info struct for this callbackID
	callbackInfo, exists := callbackRegistry[callbackID]
	callbackMutex.Unlock()

	// Create Go call data object from C call_data
	currentCallData := XtPointer(unsafe.Pointer(call_data))

	if exists {
		// Call callback with clientData and CallData
		callbackInfo.Func(callbackInfo.Widget, callbackInfo.ClientData, currentCallData)
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
	return XtTranslations(C.XtParseTranslationTable(cTableString))
}

func XtAugmentTranslations(w Widget, translations XtTranslations) {
	C.XtAugmentTranslations(C.Widget(w), C.XtTranslations(translations))
}

// ============================================================================
// XtEvent handler code
// ============================================================================

//export goEventHandlerBridge
func goEventHandlerBridge(w C.Widget, client_data C.XtPointer, event *C.XEvent, continue_to_dispatch *C.Boolean) {
	if client_data == nil {
		return
	}

	// convert client_data to Go cgo.Handle
	handlePtr := (*cgo.Handle)(unsafe.Pointer(client_data))

	// get go function event handler to call
	goFunc, ok := handlePtr.Value().(XtEventHandler)
	if !ok || goFunc == nil {
		return
	}

	// pack C types in Go types
	goWidget := Widget(w)
	goEvent := (*XEvent)(unsafe.Pointer(event))
	goCAddr := CAddr(client_data)

	// execute event handler with correct arguments
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

func XtCloseDisplay(display *Display) {
	C.XtCloseDisplay((*C.Display)(unsafe.Pointer(display)))
}

func XtDestroyApplicationContext(appContext XtAppContext) {
	C.XtDestroyApplicationContext(C.XtAppContext(appContext))
}

func XtCreateApplicationContext() XtAppContext {
	return (XtAppContext(C.XtCreateApplicationContext()))
}

func XtOpenDisplay(appContext XtAppContext, display_string string, application_name string,
	application_class string, options []OptionDescRec, num_options, argc *int, argv []string) *Display {
	var cDisplay = C.CString(display_string)
	defer C.free(unsafe.Pointer(cDisplay))
	cName := C.CString(application_name)
	defer C.free(unsafe.Pointer(cName))
	cClass := C.CString(application_class)
	defer C.free(unsafe.Pointer(cClass))
	var cArgc = C.int(*argc)

	// 3. options-Slice zu C-Array-Pointer konvertieren
	var cOptions *C.XrmOptionDescRec
	if len(options) > 0 {
		// Zeiger auf das erste Element des Slices holen und zu C-Typ casten
		cOptions = (*C.XrmOptionDescRec)(unsafe.Pointer(&options[0]))
	}
	cOptionLength := C.Cardinal(len(options))

	// 4. argv-Slice ([]string) zu C-Style char** konvertieren
	var cArgv **C.char
	if len(argv) > 0 {
		// Allokiere ein C-Array von "char*" Pointern
		cArgvArr := make([]*C.char, len(argv))
		for i, s := range argv {
			cStr := C.CString(s)
			defer C.free(unsafe.Pointer(cStr)) // Wird nach Funktionsende freigegeben
			cArgvArr[i] = cStr
		}
		// Zeiger auf das erste Element des Pointer-Arrays holen
		cArgv = (**C.char)(unsafe.Pointer(&cArgvArr[0]))
	}

	d := C.XtOpenDisplay(C.XtAppContext(appContext), cDisplay, cName, cClass,
		cOptions, cOptionLength, &cArgc, cArgv)
	return (*Display)(unsafe.Pointer(d))
}

func XtToolkitInitialize() {
	C.XtToolkitInitialize()
}

func XtAppCreateShell(application_name string, application_class string, widgetClass WidgetClass,
	display *Display, args []Arg) Widget {

	cName := C.CString(application_name)
	defer C.free(unsafe.Pointer(cName))
	cClass := C.CString(application_class)
	defer C.free(unsafe.Pointer(cClass))

	cWidgetClass := C.WidgetClass(widgetClass)
	cDisplay := (*C.Display)(unsafe.Pointer(display))

	// handle ArgList
	var cArgsList unsafe.Pointer
	cNumArgs := C.Cardinal(len(args))
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

	w := C.call_XtAppCreateShell(cName, cClass, cWidgetClass, cDisplay, cArgsList, cNumArgs)

	return (Widget)(unsafe.Pointer(w))
}

func XtWarning(message string) {
	C.XtWarning(C.CString(message))
}

func XtStringConversionWarning(from_value string, to_type string) {
	C.XtStringConversionWarning(C.CString(from_value), C.CString(to_type))
}

func calculateByteSize(class string, rtype string) int {
	switch rtype {
	case "Integer", "Int": // entspricht XtRInt
		//return int(C.sizeof_int) // liefert verlässlich 4
		return int(unsafe.Sizeof(int(0)))

	case "Boolean": // entspricht XtRBoolean
		return int(C.sizeof_Boolean) // liefert meist 1 (oder Compiler-abhängig gepadded)

	case "String": // entspricht XtRString (C-Typ: char*)
		// Ein C-Zeiger hat exakt dieselbe Bit-Breite wie ein Go-Pointer/uintptr.
		// Das vermeidet Cgo-Übersetzungsfehler vollständig!
		return int(unsafe.Sizeof(uintptr(0))) // liefert verlässlich 8 auf 64-Bit

	case "Pixel": // entspricht XtRPixel (C-Typ: unsigned long)
		return int(C.sizeof_ulong) // liefert verlässlich 8 auf 64-Bit

	case "Widget": // entspricht XtRWidget (C-Typ: Widget, ein Pointer)
		return int(unsafe.Sizeof(uintptr(0))) // liefert verlässlich 8

	default:
		// Sicherer Fallback für unmanaged Typen
		return int(C.sizeof_int)
	}
}

// ParseBufferToStruct uses data from resultBuffer and writes it to the
// related fields in "base" struct. Uses reflection
func ParseBufferToStruct(base any, resultBuffer []byte) error {
	val := reflect.ValueOf(base)

	// make sure this is a pointer to a struct
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("base muss ein gültiger, nicht-nil Zeiger auf eine Struktur sein")
	}

	// get structure information
	structVal := val.Elem()
	if structVal.Kind() != reflect.Struct {
		return fmt.Errorf("base zeigt nicht auf eine Struktur (Kind: %s)", structVal.Kind())
	}
	structType := structVal.Type()

	// iterate through all fields in struct
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		fieldType := structType.Field(i)

		// byte offset and size of this field inside go struct
		// Go takes care for alignment
		offset := fieldType.Offset
		size := field.Type().Size()

		// Sicherheitscheck, damit wir nicht über die Buffer-Grenzen lesen
		if int(offset+size) > len(resultBuffer) {
			return fmt.Errorf("Field %s lies outside of resultBuffer", fieldType.Name)
		}

		// read value from buffer into field, using correct type
		switch field.Kind() {
		case reflect.Int32:
			// copy 4 bytes and use as int32 value
			val := *(*int32)(unsafe.Pointer(&resultBuffer[offset]))
			// convert to int64 for Go and set field value
			field.SetInt(int64(val))

		case reflect.Int:
			// Auf 64-Bit-Systemen ist das Go-eigene 'int' meist 64-Bit groß
			// WICHTIG: Xt schreibt als 'XtRInt' IMMER nur 4 Bytes (int32) in den Buffer!
			// Wir lesen daher strikt nur die 4 Bytes von Xt...
			val32 := *(*int32)(unsafe.Pointer(&resultBuffer[offset]))
			/// convert to int64 for Go and set field value
			field.SetInt(int64(val32))

		case reflect.Bool:
			// Xt liefert Booleans oft als 4-Byte-Wert (wie gesehen: 1 0 0 0)
			// Wir prüfen, ob im 4-Byte Segment ein Wert ungleich Null steht
			isTrue := false
			if size == 4 {
				val := *(*int32)(unsafe.Pointer(&resultBuffer[offset]))
				isTrue = val != 0
			} else {
				// Falls es doch ein echtes 1-Byte C-Boolean ist
				val := resultBuffer[offset]
				isTrue = val != 0
			}
			field.SetBool(isTrue)

		case reflect.String:
			// Extrem wichtig: Xt schreibt bei Strings die 8-Byte Speicheradresse (char*)
			// in den Buffer, nicht den Text selbst!
			cStrPtr := *(*unsafe.Pointer)(unsafe.Pointer(&resultBuffer[offset]))

			if cStrPtr != nil {
				// Konvertiere die C-Adresse zurück in einen echten Go-String
				// (Hier nutzen wir C.GoString, stelle sicher, dass "C" importiert ist)
				goStr := C.GoString((*C.char)(cStrPtr))
				field.SetString(goStr)
			} else {
				field.SetString("")
			}

		default:
			return fmt.Errorf("Yet unsupported field type, reflection fails for: %s (type %s)", fieldType.Name, field.Kind())
		}
	}

	return nil
}

func XtGetApplicationResources(w Widget, base any, resources []GoXtResource, num_resources int,
	args ArgList) {

	/*
		for _, resource := range resources {
			fmt.Printf("%v\n", resource)
		}
	*/

	var resultBuffer [64]byte

	// The C implementation uses two parameters to write back requested resources
	// How C implementation works:
	// 1. "base" - this a generic pointer to a C struct where results are written to
	//             the struct offers "place" for result data (primitives, pointers to char* etc
	// 2. "resources: fields "size" and "offset" are information regarding result data size
	//             and offset into/inside "base" space. This allows C to do a generic copy
	//             like "copy( from, to, size) where to and size are given by "base", offset and size

	// Approach to do this with a Go implementation
	// 1. Set up internal C array from Go data which can be used and filled by C
	// 2. Copy back retrieved values from C array into Go struct pointed to by "base"

	// First step is just set up C array with all required data feasible for call of C function
	// While string parts are trivial, the complicated fields are: resource_offset, resource_offset
	// A. resource_size: this must be the value which sizeof(<type>) calculated in C.
	// B. resource_offset: this must point to some memory location aligned correctly and with enough space
	//    for datatype to be retrieved

	// Solution approach that was taken:
	// issue A: is solved by function calculateByteSize()
	// issue B: we use a byte buffer and let write C the data there using size+offset values
	//          From the flat byte buffer, the return value struct is filled by function
	//          ParseBufferToStruct()

	// handle resources
	var cResList unsafe.Pointer
	cNumRes := C.Cardinal(len(resources))
	var cResArray []C.XtResource

	resultBufferOffset := 0
	if len(resources) > 0 {
		cResSlice := C.malloc(C.size_t(len(resources)) * C.sizeof_XtResource)
		defer C.free(cResSlice)

		//cResArray := (*[1 << 20]C.XtResource)(cResSlice)[:len(resources):len(resources)]
		cResArray = unsafe.Slice((*C.XtResource)(cResSlice), len(resources))
		for i, res := range resources {
			cResArray[i].resource_name = C.CString(res.Name)
			cResArray[i].resource_class = C.CString(res.Class)
			cResArray[i].resource_type = C.CString(res.Rtype)

			size := calculateByteSize(res.Class, res.Rtype)
			if size > 1 {
				remainder := resultBufferOffset % size
				if remainder != 0 {
					resultBufferOffset += (size - remainder) // Füge Padding-Bytes ein
				}
			}
			cResArray[i].resource_size = C.Cardinal(size)
			cResArray[i].resource_offset = C.Cardinal(resultBufferOffset)
			resultBufferOffset += size

			cResArray[i].default_type = C.CString(res.Default_type)
			cResArray[i].default_addr = (C.XtPointer)(res.Default_addr)
			defer C.free(unsafe.Pointer(cResArray[i].resource_name))
			defer C.free(unsafe.Pointer(cResArray[i].resource_class))
			defer C.free(unsafe.Pointer(cResArray[i].resource_type))
			defer C.free(unsafe.Pointer(cResArray[i].default_type))
		}
		cResList = cResSlice
	}

	// handle ArgList
	var cArgsList unsafe.Pointer
	cNumArgs := C.Cardinal(args.Size)
	if args.Size > 0 {
		cArgsSlice := C.malloc(C.size_t(args.Size) * C.sizeof_Arg)
		defer C.free(cArgsSlice)

		cArgsArray := (*[1 << 20]C.Arg)(cArgsSlice)[:args.Size:args.Size]
		for i, arg := range args.Slice {
			cArgsArray[i].name = C.CString(arg.Name)
			cArgsArray[i].value = C.XtArgVal(arg.Value)
			defer C.free(unsafe.Pointer(cArgsArray[i].name))
		}
		cArgsList = cArgsSlice
	}

	val := reflect.ValueOf(base)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		fmt.Println("base pointer is not valid")
		return
	}

	// result bytes will be found in resultBuffer
	C.call_XtGetApplicationResources(
		C.Widget(w),
		(C.XtPointer)(unsafe.Pointer(&resultBuffer)), // Xt schreibt in den temporären byteBuffer
		(*C.XtResource)(cResList),
		cNumRes,
		cArgsList,
		cNumArgs,
	)

	/*
		for i, b := range resultBuffer {
			fmt.Printf("[%v]=%v ", i, b)
			if i > 0 && i%8 == 0 {
				fmt.Println()
			}
		}
		fmt.Println()
		fmt.Println()
	*/
	// fill in result struct
	err := ParseBufferToStruct(base, resultBuffer[:])
	if err != nil {
		fmt.Printf("Error during parsing/reflection: %v\n", err)
	}
}
