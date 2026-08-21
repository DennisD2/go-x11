package main

import "C"
import (
	"fmt"
	"go-x11/x11"
	"os"
	"unsafe"
)

func yesCallback() {
	println("YES Button\n")

	//cmd := C.GoString(clientData)
	cmd := "FAKE ls -l"
	if len(cmd) > 0 {
		//syscall.Exec()
		fmt.Printf("system(%v)\n", cmd)
	} else {
		fmt.Printf("cmd is empty string")
	}
	os.Exit(0)
}

func noCallback() {
	println("NO Button\n")
	os.Exit(0)
}

func main() {
	x11.WrapperInfo()

	argv := os.Args[1:]
	if len(argv) != 2 {
		println("usage: xcecute message-string command\n")
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

	/* Create a simple manager widget to hold the other widgets */
	bb := x11.XtCreateManagedWidget("bboard", x11.BulletinBoardWidgetClass(), shell, nil)

	/* Convert string to the form expected by Motif */
	xmStr := x11.XmStringCreateLtoR(argv[0], x11.XmFONTLIST_DEFAULT_TAG)

	/* define some args */
	args := x11.AppendArgList(nil, x11.XmNlabelString, x11.XtArgValFromXmString(xmStr))
	args = x11.AppendArgList(args, x11.XmNx, 0)
	args = x11.AppendArgList(args, x11.XmNy, 0)
	/* Create a Motif widget to display the string */
	msg := x11.XtCreateManagedWidget("message", x11.LabelWidgetClass(), bb, args)
	_ = msg

	x11.XmStringFree(xmStr) /* Free the compound string */

	var height uint16
	// 2. Ermitteln Sie die Adresse der Variable als unsafe.Pointer
	ptrToHeight := unsafe.Pointer(&height)
	// 3. Konvertieren Sie diesen Pointer in ein uintptr für Ihre Struktur
	valueAsUintptr := uintptr(ptrToHeight)
	// 4. Argumentenliste befüllen
	inargs := x11.AppendArgList(nil, x11.XmNheight, valueAsUintptr)
	x11.XtGetValues(msg, inargs)

	fmt.Printf("height: %d\n", height)

	args = x11.AppendArgList(nil, x11.XmNx, 0)
	args = x11.AppendArgList(args, x11.XmNy, uintptr(height+20))
	yes := x11.XtCreateManagedWidget("yes", x11.PushButtonWidgetClass(), bb, args)

	args = x11.AppendArgList(nil, x11.XmNx, 200)
	args = x11.AppendArgList(args, x11.XmNy, uintptr(height+20))
	no := x11.XtCreateManagedWidget("no", x11.PushButtonWidgetClass(), bb, args)

	/* add callbacks */
	x11.XtAddCallback(yes, x11.XmNactivateCallback, yesCallback)
	x11.XtAddCallback(no, x11.XmNactivateCallback, noCallback)

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(&appContext)
}
