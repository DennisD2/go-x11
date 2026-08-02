package main

// #cgo LDFLAGS: -lXm -lXt -lX11
import "C"

import (
	"go-x11/x11"
)

func main() {
	x11.WrapperInfo()

	var ctx = new(x11.AppContext)
	var options = new(x11.OptionDescList)
	var noptions = 0
	var argc = 0
	var argv []string
	fallbackResources := ""
	var argList = new(x11.ArgList)
	var narglist = 0

	shell := x11.AppInitialize(ctx, "AppClass", *options, noptions,
		&argc, argv, fallbackResources, argList, narglist)

	/* Convert the first argument to the form expected by Motif */
	xmstr := x11.XmStringCreateLtoR("hehe")

	/* Create a Motif XmLabel widget to display the string */
	value := x11.XtArgValFromXmString(xmstr)
	args := x11.AppendArgList(nil, x11.XmNlabelString, value)

	widgetClass := x11.LabelWidgetClass()

	msg := x11.CreateManagedWidget("message", widgetClass, shell, args, args.Size)
	_ = msg

	x11.XmStringFree(xmstr) /* Free the compound string */
	x11.FreeArgList(args)

	/*
	 * Realize the shell and enter an event loop.
	 */
	x11.RealizeWidget(shell)
	x11.AppMainLoop(ctx)

}
