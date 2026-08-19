package main

import (
	"go-x11/x11"
	"os"
)

func main() {
	x11.WrapperInfo()

	if len(argv) == 0 {
		println("usage: ch02-memo message-string\n")
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
	//args = x11.AppendArgList(args, x11.XtNwidth, 400)
	//args = x11.AppendArgList(args, x11.XtNheight, 200)
	/* Create a Motif widget to display the string */
	widgetClass := x11.LabelWidgetClass() // or x11.LabelWidgetClass()
	msg := x11.XtCreateManagedWidget("message", widgetClass, shell, args)
	_ = msg

	x11.XmStringFree(xmStr) /* Free the compound string */

	/* add a callback */
	//x11.XtAddCallback(msg, x11.XmNactivateCallback, func() {
	//	println("Button was selected! Pure Go code + X11 did this!")
	//})

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(&appContext)
}
