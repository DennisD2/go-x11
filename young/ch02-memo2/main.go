package main

import (
	"go-x11/x11"
	"os"
)

func quitAction(w x11.Widget, event x11.XEvent, params []string) {
	println("quitAcion, q was pressed")
	os.Exit(0)
}

func writeAction(w x11.Widget, event x11.XEvent, params []string) {
	println("writeAction, w was pressed")
}

var actionsTable = []x11.XtActionsRec{
	{"bye", quitAction},
	{"write", writeAction},
}

/* Bind the action "bye()" to typing the key "Q" */
var defaultTranslations []string = []string{"<Key>Q:  bye()", "<Key>W:  write()"}

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

	/* Convert some string to the form expected by Motif */
	xmStr := x11.XmStringCreateLtoR(argv[0], x11.XmFONTLIST_DEFAULT_TAG)

	/* define some args */
	args := x11.AppendArgList(nil, x11.XmNlabelString, x11.XtArgValFromXmString(xmStr))
	args = x11.AppendArgList(args, x11.XtNwidth, 400)
	args = x11.AppendArgList(args, x11.XtNheight, 200)
	/* Create a Motif widget to display the string */
	widgetClass := x11.PushButtonWidgetClass() // or x11.LabelWidgetClass()
	msg := x11.XtCreateManagedWidget("message", widgetClass, shell, args)

	x11.XmStringFree(xmStr) /* Free the compound string */

	/* Register the action functions */
	x11.XtAppAddActions(appContext, actionsTable)

	/* Compile the translation table */
	transTable := x11.XtParseTranslationTable(defaultTranslations)

	/*
	 * Merge the new translations with any existing
	 * translations for the label widget.
	 */
	x11.XtAugmentTranslations(msg, transTable)

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(&appContext)
}
