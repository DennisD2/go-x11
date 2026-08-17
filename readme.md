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

What do we have so far:

Wrapper structs for:
* XtAppContext
* Widget
* WidgetClass
* XtArgVal

Go structs for:
* OptionDescRec
* Arg
* ArgList
* XmString

Xt Functions:
* XtAppInitialize
* XtCreateWidget
* XtCreateManagedWidget
* XtRealizeWidget
* XtAppMainLoop
* XtAddCallback

Xm Functions:
* XmStringCreateLtoR
* XmStringFree

Xt Args handling helpers:
* AppendArgList
* XtArgValFromInt
* XtArgValFromString
* XtArgValFreeString

## Test it
in main.go, a single Motif widget is created, and a callback is appended to it.

## Internals

Uses CGO to access C code.

Analysis of created CGO wrapper code:
```shell
go tool cgo main.go
```
This creates a directory _obj. And adds all generated glue code from CGO. 
Very interesting is e.g. _obj/_cgo_gotypes.go for all in-betrween types/structs/casts.

## More information
* Good spicker, german - https://opensource.archium.org/index.php/Der_Golang-Spicker
* unsafe package - https://leapcell.medium.com/gos-unsafe-unlocking-performance-hacks-with-a-risk-16d1d8dd9afb