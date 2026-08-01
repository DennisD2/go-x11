package xt

import "C"

// #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
// #cgo LDFLAGS: -lXt -lXm -lX11
// #include <stdlib.h>
// #include <X11/Intrinsic.h>
// #include <Xm/Xm.h>
// #include <Xm/Label.h>
import "C"
import "unsafe"

type Widget struct {
	Widget C.Widget
}

type AppContext struct {
	AppContext C.XtAppContext
}

type OptionDescList struct {
	OptionDescList C.XrmOptionDescList
}

type ArgList struct {
	ArgList C.ArgList
}

type _XtString struct {
	_XtString *C._XtString
}

func convertIt_ArgList(in *ArgList) C.ArgList {
	return in.ArgList
}

func convertIt_XtString(in []string) **C.char {
	if len(in) == 0 {
		return nil
	}
	ret := make([]*C.char, len(in))
	for i, arg := range in {
		ret[i] = C.CString(arg)
	}
	return &ret[0]
}

func AppInitialize(appContext *AppContext, appClass string, options OptionDescList, numOptions int,
	argc *int, argv []string, fallbackResources string, wargs *ArgList, numWargs int) Widget {

	c_appClass := C.CString(appClass)
	defer C.free(unsafe.Pointer(c_appClass))

	var cnumOptions = C.uint(numOptions)
	var cargc = C.int(*argc)
	var cnumWargs = C.uint(numWargs)

	var cargv = convertIt_XtString(argv)
	c_fallbackResources := C.XtNewString(C.CString(""))
	shell := C.XtAppInitialize(&appContext.AppContext, c_appClass, options.OptionDescList, cnumOptions,
		&cargc, cargv, &c_fallbackResources, (*C.struct___0)(unsafe.Pointer(wargs.ArgList)), cnumWargs)

	r := new(Widget)
	r.Widget = shell
	return *r
}

func CreateManagedWidget(name string, widgetClass unsafe.Pointer, parent Widget, args unsafe.Pointer, num_args int) Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cnum_args = C.uint(num_args)

	widget := C.XtCreateManagedWidget(c_name, (*C.struct__WidgetClassRec)(widgetClass),
		parent.Widget, (*C.struct___0)(args), cnum_args)

	r := new(Widget)
	r.Widget = widget
	return *r
}

func RealizeWidget(w Widget) {
	C.XtRealizeWidget(w.Widget)
}

func AppMainLoop(ctx *AppContext) {
	C.XtAppMainLoop(ctx.AppContext)
}
