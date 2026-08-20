package main

import "C"
import (
	"go-x11/x11"
	"os"
)

func quit(w x11.Widget, clientData x11.CAddr, event *x11.XEvent) {
	println("Mouse Button was selected! Pure Go code + X11 did this!")
	//var e C.XEvent = event.e
	//println("XEvent:", e)
	//XtCloseDisplay(XtDisplay(w));
	//exit(0);
}

func main() {
	x11.WrapperInfo()

	if len(os.Args) < 2 {
		println("usage: go run main.go message-string")
		os.Exit(1)
	}

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

	/* Convert some string to the form expected by Motif */
	var xmStr x11.XmString
	if len(os.Args) > 1 {
		xmStr = x11.Xs_concat_words(os.Args[1:])
	} else {
		xmStr = x11.XmStringCreateLtoR("please click me!", x11.XmFONTLIST_DEFAULT_TAG)
	}

	/* define some args */
	args := x11.AppendArgList(nil, x11.XmNlabelString, x11.XtArgValFromXmString(xmStr))
	args = x11.AppendArgList(args, x11.XtNwidth, 400)
	args = x11.AppendArgList(args, x11.XtNheight, 200)
	/* Create a Motif widget to display the string */
	widgetClass := x11.LabelWidgetClass() // or x11.PushButtonWidgetClass()
	msg := x11.XtCreateManagedWidget("message", widgetClass, shell, args)
	_ = msg

	x11.XmStringFree(xmStr) /* Free the compound string */

	/* add a event handler */
	x11.XtAddEventHandler(msg, x11.ButtonPressMask, false,
		quit, nil)

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(&appContext)
}
