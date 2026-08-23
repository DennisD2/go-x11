package main

import (
	"go-x11/pkg/x11"
	"os"
)

func main() {
	x11.WrapperInfo()

	if len(os.Args) < 2 {
		println("usage: go run main.go message-string")
		os.Exit(1)
	}
	argv := os.Args[1:]

	var options []x11.OptionDescRec

	x11.XtToolkitInitialize()

	app := x11.XtCreateApplicationContext()

	nOptions := 0
	argc := 0
	dpy := x11.XtOpenDisplay(app, "", argv[0], "Memo", options, &nOptions,
		&argc, argv)

	toplevel := x11.XtAppCreateShell("memo", "Memo",
		x11.ApplicationShellWidgetClass(),
		dpy, []x11.Arg{})

	/* Convert string to the form expected by Motif */
	xmStr := x11.XmStringCreateLtoR(argv[0], x11.XmFONTLIST_DEFAULT_TAG)

	/* define some args */
	args := x11.AppendArgList(nil, x11.XmNlabelString, x11.XtArgValFromXmString(xmStr))
	/* Create a Motif widget to display the string */
	x11.XtCreateManagedWidget("message", x11.LabelWidgetClass(), toplevel, args)

	x11.XmStringFree(xmStr) /* Free the compound string */

	x11.XtRealizeWidget(toplevel)
	x11.XtAppMainLoop(app)
}
