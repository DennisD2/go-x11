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

// Xm includes
#include <Xm/Xm.h>
#include "Xm/ArrowB.h"
#include "Xm/ArrowBG.h"
#include "Xm/BulletinB.h"
#include "Xm/CascadeB.h"
#include "Xm/CascadeBG.h"
#include "Xm/ClipWindowP.h"
#include "Xm/ComboBox.h"
#include "Xm/Command.h"
#include "Xm/DesktopP.h"
#include "Xm/DialogS.h"
#include "Xm/DialogSEP.h"
#include "Xm/DragC.h"
#include "Xm/DragIcon.h"
#include "Xm/DrawingA.h"
#include "Xm/DrawnB.h"
#include "Xm/DropDown.h"
#include "Xm/DropSMgr.h"
#include "Xm/DropTrans.h"
#include "Xm/Ext18List.h"
#include "Xm/ExtObjectP.h"
#include "Xm/FileSB.h"
#include "Xm/Form.h"
#include "Xm/Frame.h"
#include "Xm/Gadget.h"
#include "Xm/GrabShell.h"
#include "Xm/Label.h"
#include "Xm/LabelG.h"
#include "Xm/List.h"
#include "Xm/MainW.h"
#include "Xm/Manager.h"
#include "Xm/MenuShell.h"
#include "Xm/MessageB.h"
#include "Xm/MultiList.h"
#include "Xm/Notebook.h"
#include "Xm/PanedW.h"
#include "Xm/Primitive.h"
//#include "Xm/Print.h"
#include "Xm/ProtocolsP.h"
#include "Xm/PushB.h"
#include "Xm/PushBG.h"
#include "Xm/RowColumn.h"
#include "Xm/SSpinB.h"
#include "Xm/SashP.h"
#include "Xm/Scale.h"
#include "Xm/ScrollBar.h"
#include "Xm/ScrolledW.h"
#include "Xm/SelectioB.h"
#include "Xm/SeparatoG.h"
#include "Xm/Separator.h"
#include "Xm/ShellEP.h"
#include "Xm/SpinB.h"
#include "Xm/TearOffBP.h"
#include "Xm/ToggleB.h"
#include "Xm/ToggleBG.h"
#include "Xm/VendorS.h"
#include "Xm/VendorSEP.h"


#include "cheader.h"
#include "wrapperInfo.h"

*/
import "C"

// ============================================================================
// X definitions
// ============================================================================
// Wrapper types
type XEvent struct{ e C.XEvent }

// ============================================================================
// Xt definitions
// ============================================================================

// String defines
var XtNaccelerators = C.GoString(C.XtNaccelerators)
var XtNallowHoriz = C.GoString(C.XtNallowHoriz)
var XtNallowVert = C.GoString(C.XtNallowVert)
var XtNancestorSensitive = C.GoString(C.XtNancestorSensitive)
var XtNbackground = C.GoString(C.XtNbackground)
var XtNbackgroundPixmap = C.GoString(C.XtNbackgroundPixmap)
var XtNbitmap = C.GoString(C.XtNbitmap)
var XtNborderColor = C.GoString(C.XtNborderColor)
var XtNborder = C.GoString(C.XtNborder)
var XtNborderPixmap = C.GoString(C.XtNborderPixmap)
var XtNborderWidth = C.GoString(C.XtNborderWidth)
var XtNcallback = C.GoString(C.XtNcallback)
var XtNchildren = C.GoString(C.XtNchildren)
var XtNcolormap = C.GoString(C.XtNcolormap)
var XtNdepth = C.GoString(C.XtNdepth)
var XtNdestroyCallback = C.GoString(C.XtNdestroyCallback)
var XtNeditType = C.GoString(C.XtNeditType)
var XtNfile = C.GoString(C.XtNfile)
var XtNfont = C.GoString(C.XtNfont)
var XtNforceBars = C.GoString(C.XtNforceBars)
var XtNforeground = C.GoString(C.XtNforeground)
var XtNfunction = C.GoString(C.XtNfunction)
var XtNheight = C.GoString(C.XtNheight)
var XtNhighlight = C.GoString(C.XtNhighlight)
var XtNhSpace = C.GoString(C.XtNhSpace)
var XtNindex = C.GoString(C.XtNindex)
var XtNinitialResourcesPersistent = C.GoString(C.XtNinitialResourcesPersistent)
var XtNinnerHeight = C.GoString(C.XtNinnerHeight)
var XtNinnerWidth = C.GoString(C.XtNinnerWidth)
var XtNinnerWindow = C.GoString(C.XtNinnerWindow)
var XtNinsertPosition = C.GoString(C.XtNinsertPosition)
var XtNinternalHeight = C.GoString(C.XtNinternalHeight)
var XtNinternalWidth = C.GoString(C.XtNinternalWidth)
var XtNjumpProc = C.GoString(C.XtNjumpProc)
var XtNjustify = C.GoString(C.XtNjustify)
var XtNknobHeight = C.GoString(C.XtNknobHeight)
var XtNknobIndent = C.GoString(C.XtNknobIndent)
var XtNknobPixel = C.GoString(C.XtNknobPixel)
var XtNknobWidth = C.GoString(C.XtNknobWidth)
var XtNlabel = C.GoString(C.XtNlabel)
var XtNlength = C.GoString(C.XtNlength)
var XtNlowerRight = C.GoString(C.XtNlowerRight)
var XtNmappedWhenManaged = C.GoString(C.XtNmappedWhenManaged)
var XtNmenuEntry = C.GoString(C.XtNmenuEntry)
var XtNname = C.GoString(C.XtNname)
var XtNnotify = C.GoString(C.XtNnotify)
var XtNnumChildren = C.GoString(C.XtNnumChildren)
var XtNorientation = C.GoString(C.XtNorientation)
var XtNparameter = C.GoString(C.XtNparameter)
var XtNpixmap = C.GoString(C.XtNpixmap)
var XtNpopupCallback = C.GoString(C.XtNpopupCallback)
var XtNpopdownCallback = C.GoString(C.XtNpopdownCallback)
var XtNresize = C.GoString(C.XtNresize)
var XtNreverseVideo = C.GoString(C.XtNreverseVideo)
var XtNscreen = C.GoString(C.XtNscreen)
var XtNscrollProc = C.GoString(C.XtNscrollProc)
var XtNscrollDCursor = C.GoString(C.XtNscrollDCursor)
var XtNscrollHCursor = C.GoString(C.XtNscrollHCursor)
var XtNscrollLCursor = C.GoString(C.XtNscrollLCursor)
var XtNscrollRCursor = C.GoString(C.XtNscrollRCursor)
var XtNscrollUCursor = C.GoString(C.XtNscrollUCursor)
var XtNscrollVCursor = C.GoString(C.XtNscrollVCursor)
var XtNselection = C.GoString(C.XtNselection)
var XtNselectionArray = C.GoString(C.XtNselectionArray)
var XtNsensitive = C.GoString(C.XtNsensitive)
var XtNshown = C.GoString(C.XtNshown)
var XtNspace = C.GoString(C.XtNspace)
var XtNstring = C.GoString(C.XtNstring)
var XtNtextOptions = C.GoString(C.XtNtextOptions)
var XtNtextSink = C.GoString(C.XtNtextSink)
var XtNtextSource = C.GoString(C.XtNtextSource)
var XtNthickness = C.GoString(C.XtNthickness)
var XtNthumb = C.GoString(C.XtNthumb)
var XtNthumbProc = C.GoString(C.XtNthumbProc)
var XtNtop = C.GoString(C.XtNtop)
var XtNtranslations = C.GoString(C.XtNtranslations)
var XtNunrealizeCallback = C.GoString(C.XtNunrealizeCallback)
var XtNupdate = C.GoString(C.XtNupdate)
var XtNuseBottom = C.GoString(C.XtNuseBottom)
var XtNuseRight = C.GoString(C.XtNuseRight)
var XtNvalue = C.GoString(C.XtNvalue)
var XtNvSpace = C.GoString(C.XtNvSpace)
var XtNwidth = C.GoString(C.XtNwidth)
var XtNwindow = C.GoString(C.XtNwindow)
var XtNx = C.GoString(C.XtNx)
var XtNy = C.GoString(C.XtNy)
var XtNfontSet = C.GoString(C.XtNfontSet)
var XtNcreateHook = C.GoString(C.XtNcreateHook)
var XtNchangeHook = C.GoString(C.XtNchangeHook)
var XtNconfigureHook = C.GoString(C.XtNconfigureHook)
var XtNgeometryHook = C.GoString(C.XtNgeometryHook)
var XtNdestroyHook = C.GoString(C.XtNdestroyHook)
var XtNshells = C.GoString(C.XtNshells)
var XtNnumShells = C.GoString(C.XtNnumShells)

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
// Xm definitions
// ============================================================================
// string defines
var XmNlabelString = C.GoString(C.XmNlabelString)

func ArrowButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmArrowButtonWidgetClass}
}
func ArrowButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmArrowButtonGadgetClass}
}
func BulletinBoardWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmBulletinBoardWidgetClass}
}
func CascadeButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmCascadeButtonWidgetClass}
}
func CascadeButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmCascadeButtonGadgetClass}
}
func ClipWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmClipWindowWidgetClass}
}
func ComboBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmComboBoxWidgetClass}
}
func CommandWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmCommandWidgetClass}
}
func DesktopClass() WidgetClass {
	return WidgetClass{c: C.xmDesktopClass}
}
func DialogShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDialogShellWidgetClass}
}
func DialogShellExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDialogShellExtObjectClass}
}
func DragContextClass() WidgetClass {
	return WidgetClass{c: C.xmDragContextClass}
}
func DragIconObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDragIconObjectClass}
}
func DrawingAreaWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDrawingAreaWidgetClass}
}
func DrawnButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDrawnButtonWidgetClass}
}
func DropDownWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDropDownWidgetClass}
}
func DropSiteManagerObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDropSiteManagerObjectClass}
}
func DropTransferObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDropTransferObjectClass}
}
func Ext18ListWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmExt18ListWidgetClass}
}
func ExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmExtObjectClass}
}
func FileSelectionBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmFileSelectionBoxWidgetClass}
}
func FormWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmFormWidgetClass}
}
func FrameWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmFrameWidgetClass}
}
func GadgetClass() WidgetClass {
	return WidgetClass{c: C.xmGadgetClass}
}
func GrabShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmGrabShellWidgetClass}
}
func LabelWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmLabelWidgetClass}
}
func LabelGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmLabelGadgetClass}
}
func ListWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmListWidgetClass}
}
func MainWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMainWindowWidgetClass}
}
func ManagerWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmManagerWidgetClass}
}
func MenuShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMenuShellWidgetClass}
}
func MessageBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMessageBoxWidgetClass}
}
func MultiListWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMultiListWidgetClass}
}
func NotebookWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmNotebookWidgetClass}
}
func PanedWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPanedWindowWidgetClass}
}
func PrimitiveWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPrimitiveWidgetClass}
}

//	func PrintShellWidgetClass() WidgetClass {
//		return WidgetClass{c: C.xmPrintShellWidgetClass}
//	}
func ProtocolObjectClass() WidgetClass {
	return WidgetClass{c: C.xmProtocolObjectClass}
}
func PushButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPushButtonWidgetClass}
}
func PushButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmPushButtonGadgetClass}
}
func RowColumnWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmRowColumnWidgetClass}
}
func SimpleSpinBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSimpleSpinBoxWidgetClass}
}
func SashWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSashWidgetClass}
}
func ScaleWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmScaleWidgetClass}
}
func ScrollBarWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmScrollBarWidgetClass}
}
func ScrolledWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmScrolledWindowWidgetClass}
}
func SelectionBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSelectionBoxWidgetClass}
}
func SeparatorGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmSeparatorGadgetClass}
}
func SeparatorWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSeparatorWidgetClass}
}
func ShellExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmShellExtObjectClass}
}
func SpinBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSpinBoxWidgetClass}
}
func TearOffButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmTearOffButtonWidgetClass}
}
func ToggleButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmToggleButtonWidgetClass}
}
func ToggleButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmToggleButtonGadgetClass}
}
func vendorShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.vendorShellWidgetClass}
}
func VendorShellExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmVendorShellExtObjectClass}
}

// XmString wraps C.XmString (Motif compound string)
type XmString struct {
	XmString C.XmString
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

func XtAppInitialize(
	appContext *XtAppContext,
	appClass string,
	options []OptionDescRec,
	goArgs []string,
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
		cArgsPtr,
		cFallbacks,
		cArgsList, numArgs,
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

func ManageChild(w Widget) {
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
// string defines
var XmNactivateCallback = C.GoString(C.XmNactivateCallback)

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
