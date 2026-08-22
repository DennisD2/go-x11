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

	var appContext x11.XtAppContext
	var options []x11.OptionDescRec
	fallbacks := []string{""}
	//initArgs := []x11.Arg{{Name: "width", Value: 200}}
	initArgs := []x11.Arg{}

	shell := x11.XtAppInitialize(
		&appContext,
		"GoXtApp",
		options,
		os.Args,
		fallbacks,
		initArgs,
	)

	/* Convert string to the form expected by Motif */
	xmStr := x11.XmStringCreateLtoR(argv[0], x11.XmFONTLIST_DEFAULT_TAG)

	/* define some args */
	args := x11.AppendArgList(nil, x11.XmNlabelString, x11.XtArgValFromXmString(xmStr))
	/* Create a Motif widget to display the string */
	widgetClass := x11.LabelWidgetClass() // or x11.LabelWidgetClass()
	x11.XtCreateManagedWidget("message", widgetClass, shell, args)

	x11.XmStringFree(xmStr) /* Free the compound string */

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(appContext)
}
