package x11

import (
	"unsafe"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include <sys/types.h>

// Xt includes
#include <X11/Intrinsic.h>
#include <X11/Composite.h>

#include "cheader.h"
#include "wrapperInfo.h"

*/
import "C"

type XrmValue struct {
	size int
	addr XPointer
}

type GoXtConverter func(args *XrmValue, num_args int, from *XrmValue, to *XrmValue)

// C.XtAddressMode enum values see
var XtAddress = int(C.XtAddress)

type XtConvertArgList struct {
	address_mode int // XtAddressMode
	address_id   XtPointer
	size         int
}

func XtAddConverter(from_type string, to_type string, converter GoXtConverter,
	convert_args XtConvertArgList) {

	// We need to set up a C converter function, which points back into our Go converter function
	// - Set a general cConverterBridge() receiving all conversion calls
	// - depending on incoming types of from and to values (e.g. char* and int) we lookup in a map[from,to] -> goFunction
	//   to retrieve the C function to call
	// - call function
	// - copy result back to C "to" parameter

	// set up conversion C function from: converter
	var cConv unsafe.Pointer
	var cNumArgs = C.Cardinal(0)

	// set up conversion ArgList from: convert_args
	var cConvertArgs unsafe.Pointer

	XtWarning("XtAddConverter to be implemented")
	C.call_XtAddConverter(C.CString(from_type), C.CString(to_type), cConv, cConvertArgs, cNumArgs)
}

func XtAppAddConverter(ctx XtAppContext, from_type string, to_type string, converter GoXtConverter,
	convert_args XtConvertArgList) {
	//C.XtAppAddConverter(C.XtAppContext(ctx), C.CString(from_type), C.CString(to_type), converter, convert_args, C.Cardinal(num_args))
}
