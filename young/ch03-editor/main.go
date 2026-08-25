package main

import (
	"go-x11/pkg/x11"
	"os"
)

func main() {
	x11.WrapperInfo()

	var app x11.XtAppContext
	var xrmOptions []x11.XrmOptionDescRec
	fallbacks := []string{""}
	initArgs := []x11.Arg{}

	shell := x11.XtAppInitialize(
		&app,
		"Editor",
		xrmOptions,
		os.Args,
		fallbacks,
		initArgs,
	)

	// The overall window layout is handled by an XmForm widget
	panel := x11.XtCreateManagedWidget("message", x11.FormWidgetClass(), shell, &x11.ArgList{})

	//  An XmRowColumn widget holds the buttons along the top of the window
	args := x11.AppendArgList(nil, x11.XmNnumColumns, 3)
	args = x11.AppendArgList(args, x11.XmNorientation, uintptr(x11.XmHORIZONTAL))
	args = x11.AppendArgList(args, x11.XmNtopAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNrightAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNleftAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNbottomAttachment, uintptr(x11.XmATTACH_NONE))
	commands := x11.XtCreateManagedWidget("commands", x11.RowColumnWidgetClass(), panel, args)

	// Another XmRowColumn widget contains a column of buttons along the left side of the window
	args = x11.AppendArgList(nil, x11.XmNnumColumns, 1)
	args = x11.AppendArgList(args, x11.XmNorientation, uintptr(x11.XmVERTICAL))
	args = x11.AppendArgList(args, x11.XmNtopAttachment, uintptr(x11.XmATTACH_WIDGET))
	args = x11.AppendArgList(args, x11.XmNtopWidget, uintptr(commands))
	args = x11.AppendArgList(args, x11.XmNrightAttachment, uintptr(x11.XmATTACH_NONE))
	args = x11.AppendArgList(args, x11.XmNleftAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNbottomAttachment, uintptr(x11.XmATTACH_FORM))
	options := x11.XtCreateManagedWidget("options", x11.RowColumnWidgetClass(), panel, args)

	// The middle window, in which the application can display text or graphics is an XmDrawingArea widget.
	args = x11.AppendArgList(nil, x11.XmNtopAttachment, uintptr(x11.XmATTACH_WIDGET))
	args = x11.AppendArgList(args, x11.XmNtopWidget, uintptr(commands))
	args = x11.AppendArgList(args, x11.XmNrightAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNleftWidget, uintptr(options))
	args = x11.AppendArgList(args, x11.XmNleftAttachment, uintptr(x11.XmATTACH_WIDGET))
	args = x11.AppendArgList(args, x11.XmNbottomAttachment, uintptr(x11.XmATTACH_FORM))
	canvas := x11.XtCreateManagedWidget("canvas", x11.DrawingAreaWidgetClass(), panel, args)
	_ = canvas

	x11.XtCreateManagedWidget("button1", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("button2", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("button3", x11.PushButtonWidgetClass(), commands, nil)

	x11.XtCreateManagedWidget("button1", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("button2", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("button3", x11.PushButtonWidgetClass(), options, nil)

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(app)
}
