package main

import "C"

// #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
// #cgo LDFLAGS: -lXt -lX11
// #include <stdlib.h>
// #include <X11/Intrinsic.h>
import "C"
import "unsafe"

type Widget struct {
	widget C.Widget
}

type AppContext struct {
	appContext C.XtAppContext
}

type OptionDescList struct {
	optionDescList C.XrmOptionDescList
}

type ArgList struct {
	argList C.ArgList
}

type _XtString struct {
	_XtString *C._XtString
}

func convertIt_ArgList(in *ArgList) C.ArgList {
	return in.argList
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

func AppInitialize(ctx *AppContext, appClass string, options OptionDescList, num_options int,
	argc *int, argv []string, fallbackResources string, wargs *ArgList, num_wargs int) Widget {

	c_appClass := C.CString(appClass)
	defer C.free(unsafe.Pointer(c_appClass))

	var cnum_options = C.uint(num_options)
	var cargc = C.int(*argc)
	var cnum_wargs = C.uint(num_wargs)

	var cargv = convertIt_XtString(argv)
	c_fallbackResources := C.XtNewString(C.CString(""))
	shell := C.XtAppInitialize(&ctx.appContext, c_appClass, options.optionDescList, cnum_options,
		&cargc, cargv, &c_fallbackResources, (*C.struct___1)(unsafe.Pointer(wargs.argList)), cnum_wargs)

	r := new(Widget)
	r.widget = shell
	return *r
}

func CreateManagedWidget(name string, widgetClass C.WidgetClass, parent Widget, args C.ArgList, num_args int) Widget {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cnum_args = C.uint(num_args)
	
	widget := C.XtCreateManagedWidget(c_name, widgetClass, parent.widget,
		(*C.struct___1)(unsafe.Pointer(args)), cnum_args)

	r := new(Widget)
	r.widget = widget
	return *r
}

func RealizeWidget(w Widget) {
	C.XtRealizeWidget(w.widget)
}

func AppMainLoop(ctx *AppContext) {
	C.XtAppMainLoop(ctx.appContext)
}
