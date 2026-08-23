## go-x11 - Go Wrapper for X11 libraries
Work in progress.

Idea is to offer all libX* functions and associated defines+structures+whatever in a single package called "x11".
Parameter names, meanings and return values should be as close to the original C API to make it possible to
use existing documents (e.g. *The definitive guides to the XWindow system*) and knowhow to develop.
The Go code (example here: main.c) does not need to use "C" go module, unsafe pointers and such. From the user side, everything
is clean Go code. Inside the x11 module, all these nasty conversion and casting things are encapsulated.
This is the approach for the code.

I've looked at the xgb package, which may be better and more modern. But it looked to me at first glance, that all
concepts of XWindow were thrown away and xgb brings its own new ideas for the API. This is not what I wanted, 
because I am very familiar with the old C API and wanted to bring this into the Go world.

The implementation uses CGO, wrappers for many C types, some Go-types where their internals need to be accessed 
and some "glue" C functions inside x11 package where nasty casts are required, which would not be possible in Go.

## How does the code look like
To get a feeling, below is main.go listing. This will initialize X, sets up a button with some defined resources,
displays the button and responds to a callback if button is pressed after entering application event mainloop in
last line.

```cgo
package main

import (
	"fmt"
	"go-x11/pkg/x11"
	"os"
)

func main() {
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
	widgetClass := x11.PushButtonWidgetClass() // or x11.LabelWidgetClass()
	msg := x11.XtCreateManagedWidget("message", widgetClass, shell, args)
	_ = msg

	x11.XmStringFree(xmStr) /* Free the compound string */

	/* add a callback */
	x11.XtAddCallback(msg, x11.XmNactivateCallback, func(w x11.Widget, clientData x11.XtPointer, callData x11.XtPointer) {
		println("Button was selected! Pure Go code + X11 did this!")
	}, x11.XtPointer(nil))

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(appContext)
}
```

## Build it
```shell
go build .
# or build with all demo apps
go build ./...
```

## Run it
in main.go, a single Motif widget is created, and a callback is appended to it.
```shell
go run . 
# or with some args
go run . abc def ghi
```

## Tests
Execute tests:
```shell
go test -v ./...
```

## What do we have so far
X11 together with Motif offers are large amount of objects and functions. I have started moving to go-x11 API
with a few, needed in almost every application. Other will follow.

### From Xlib
Wrappers for:
* Display
* Screen
* Window
* XEvent

### From Xt (X Toolkit)
Wrapper for:
* XtAppContext
* Widget
* WidgetClass
* XtArgVal
* XtTranslations
* XtPointer
* EventMask

Go structs for:
* OptionDescRec
* Arg
* ArgList
* XtActionsRec
* XtEventHandler
* WidgetList

Xt Functions:
* XtInitialize
* XtAppInitialize
* XtOpenDisplay
* XtCloseDisplay
* XtCreateApplicationContext
* XtDestroyApplicationContext
* XtToolkitInitialize
* XtAppCreateShell
* XtWarning
* XtCreateWidget
* XtCreateManagedWidget
* XtRealizeWidget
* XtAppMainLoop
* XtAddCallback
* XtDispatchEvent
* XtNextEvent
* XtAppNextEvent
* XtIsRealized
* XtIsManaged
* XtDestroyWidget
* XtDisplay
* XtScreen
* XtWindow
* XtSetValues
* XtGetValues
* XtManageChildren
* XtAppAddActions
* XtParseTranslationTable
* XtAugmentTranslations
* XtAddEventHandler

## From Xm (Motif)
Wrappers for:
* XmString
* XmStringCharset

Xm Functions:
* XmStringCreate
* XmStringCreateLtoR
* XmStringFree
* XmStringConcat

## Misc. and helpers
Wrapper for:
* CAddr

Xt Args handling helpers:
* AppendArgList
* XtArgValFromInt
* XtArgValFromString
* XtArgValFromXmString
* XtArgValFreeString
* ConvertAnyEvent
* ConvertButtonEvent
* ConvertKeyEvent
* ConvertMotionNotifyEvent

## From libXS (See Young's Book)
* Xs_concat_words

## Open issues
* API deviates slightly from X11. Check this, use tests. Not done yet.
* XtAppAddActions and Action handling is not complete. Action functions
  cannot take parameters, this should be added.

## Fixed issues
* XtAddEventHandler(w Widget, eventMask EventMask, nonMaskable bool, proc GoXtEventHandler, clientData CAddr) is
  not using or handling clientData. Needs to be fixed.


## Internals

Uses CGO to access C code.

Analysis of created CGO wrapper code:
```shell
go tool cgo main.go
```
This creates a directory _obj. And adds all generated glue code from CGO. 
Very interesting is e.g. _obj/_cgo_gotypes.go for all in-betrween types/structs/casts.

## More information
* CGo - https://blog.marlin.org/cgo-referencing-c-library-in-go
* "unsafe" et al - https://leapcell.medium.com/gos-unsafe-unlocking-performance-hacks-with-a-risk-16d1d8dd9afb
* unsafe package - https://pkg.go.dev/unsafe@go1.26.5
* Good spicker, german - https://opensource.archium.org/index.php/Der_Golang-Spicker
* unsafe package - https://leapcell.medium.com/gos-unsafe-unlocking-performance-hacks-with-a-risk-16d1d8dd9afb