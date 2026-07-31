package main

// #cgo LDFLAGS: -lXm -lXt -lX11
// #include <X11/Intrinsic.h>
// #include <Xm/Xm.h>          /* Required by all Motif applications */
// #include <Xm/Label.h>       /* Required by XmLabel widget */
// #include <hello.h>
import "C"
import (
	"unsafe"
)

func convertIt(in C.XmString) C.XtArgVal {
	return *(*C.XtArgVal)(unsafe.Pointer(&in))
}

func main() {
	C.hello1()

	appClass := "AppClass"
	var ctx = new(AppContext)
	var options = new(OptionDescList)
	var noptions = 0
	var argc = 0
	var argv []string
	fallbackResources := ""
	var argList = new(ArgList)
	var narglist = 0

	shell := AppInitialize(ctx, appClass, *options, noptions,
		&argc, argv, fallbackResources, argList, narglist)

	/* Convert the first argument to the form expected by Motif */
	xmstr := C.XmStringCreateLtoR(C.CString("hehe"), C.XmFONTLIST_DEFAULT_TAG)

	/* Create a Motif XmLabel widget to display the string */
	/* C.XmNlabelString, xmstr, */
	var wargc C.uint = 1
	var wargs C.Arg
	//C.XtSetArg(wargs[0], C.XmNlabelString, "hehe test")
	wargs.name = C.XmNlabelString
	wargs.value = convertIt(xmstr)

	/*msg := */
	C.XtCreateManagedWidget(C.CString("message"),
		C.xmLabelWidgetClass, shell.widget,
		&wargs,
		wargc)
	//fmt.Printf("%x\n", msg)

	C.XmStringFree(xmstr) /* Free the compound string */

	/*
	 * Realize the shell and enter an event loop.
	 */
	C.hello1()
	C.XtRealizeWidget(shell.widget)
	C.XtAppMainLoop(ctx.appContext)

	C.hello1()
}
