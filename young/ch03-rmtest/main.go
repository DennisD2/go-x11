package main

import "C"
import (
	"go-x11/pkg/x11"
	"os"
	"unsafe"
)

type ApplicationData struct {
	delay   int
	verbose bool
}

var data = ApplicationData{}

var defaultDelay = 2
var defaultVerbose = false

var resources = []x11.GoXtResource{
	{"delay", "Delay", x11.XtRInt, x11.XtRImmediate, unsafe.Pointer(&defaultDelay)},
	{"verbose", "Verbose", x11.XtRBoolean, x11.XtRString, unsafe.Pointer(&defaultVerbose)},
	/*{x11.XtNforeground, "class", "rtype", 1, 100, "deftype", (*byte)(unsafe.Pointer(&data))},
	{x11.XtNbackground, "class", "rtype", 1, 100, "deftype", (*byte)(unsafe.Pointer(&data))},*/
}

func main() {
	x11.WrapperInfo()

	argv := os.Args[1:]

	var appContext x11.XtAppContext
	var options []x11.OptionDescRec
	//initArgs := []x11.Arg{{Name: "width", Value: 200}}

	fallbacks := []string{""}
	initArgs := []x11.Arg{}

	toplevel := x11.XtAppInitialize(
		&appContext,
		"Rmtest",
		options,
		os.Args,
		fallbacks,
		initArgs,
	)

	argList := x11.ArgList{[]x11.Arg{}, 0}
	x11.XtGetApplicationResources(toplevel, (x11.XtPointer)(unsafe.Pointer(&data)), resources, len(resources), argList)

	/* Convert string to the form expected by Motif */
	xmStr := x11.XmStringCreateLtoR(argv[0], x11.XmFONTLIST_DEFAULT_TAG)

	/* define some args */
	args := x11.AppendArgList(nil, x11.XmNlabelString, x11.XtArgValFromXmString(xmStr))
	/* Create a Motif widget to display the string */
	widgetClass := x11.LabelWidgetClass() // or x11.LabelWidgetClass()
	x11.XtCreateManagedWidget("message", widgetClass, toplevel, args)

	x11.XmStringFree(xmStr) /* Free the compound string */

	x11.XtRealizeWidget(toplevel)
	x11.XtAppMainLoop(appContext)
}
