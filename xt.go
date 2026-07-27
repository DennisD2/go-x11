package main

import "C"

// #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
// #cgo LDFLAGS: -lXt -lX11
// #include <X11/Intrinsic.h>
import "C"

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

func AppInitialize(ctx *AppContext, appClass string, options OptionDescList, argList ArgList) Widget {
	c_appClass := C.CString(appClass)
	var cargc C.int = 1
	var cargv **C.char
	var cnumOptions C.uint = 0
	c_fallbackResources := C.XtNewString(C.CString(""))
	cnumArgs := C.uint(0)
	shell := C.XtAppInitialize(&ctx.appContext, c_appClass, options.optionDescList, cnumOptions,
		&cargc, cargv, &c_fallbackResources, argList.argList, cnumArgs)

	r := new(Widget)
	r.widget = shell

	C.XtRealizeWidget(shell)
	C.XtAppMainLoop(ctx.appContext)
	return *r
}
