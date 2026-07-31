package main

// #cgo LDFLAGS: -lXm -lXt -lX11
// #include <X11/Intrinsic.h>
// #include <Xm/Xm.h>          /* Required by all Motif applications */
// #include <Xm/Label.h>       /* Required by XmLabel widget */
// #include <hello.h>
import "C"
import (
	"unsafe"

	"go-xt/xt"
	"go-xt/xm"
)

func convertIt(in xm.XmString) C.XtArgVal {
	return *(*C.XtArgVal)(unsafe.Pointer(&in.XmString))
}

func main() {
	C.hello1()

	appClass := "AppClass"
	var ctx = new(xt.AppContext)
	var options = new(xt.OptionDescList)
	var noptions = 0
	var argc = 0
	var argv []string
	fallbackResources := ""
	var argList = new(xt.ArgList)
	var narglist = 0

	shell := xt.AppInitialize(ctx, appClass, *options, noptions,
		&argc, argv, fallbackResources, argList, narglist)

	/* Convert the first argument to the form expected by Motif */
	xmstr := xm.StringCreateLtoR("hehe", unsafe.Pointer(C.XmFONTLIST_DEFAULT_TAG))

	/* Create a Motif XmLabel widget to display the string */
	/* C.XmNlabelString, xmstr, */
	var wargs C.Arg
	//C.XtSetArg(wargs[0], C.XmNlabelString, "hehe test")
	wargs.name = C.XmNlabelString
	wargs.value = convertIt(xmstr)

	msg := xt.CreateManagedWidget("message",
		unsafe.Pointer(C.xmLabelWidgetClass), shell, unsafe.Pointer(&wargs), 1)
	_ = msg

	xm.StringFree(xmstr) /* Free the compound string */

	/*
	 * Realize the shell and enter an event loop.
	 */
	C.hello1()
	xt.RealizeWidget(shell)
	xt.AppMainLoop(ctx)

	C.hello1()
}
