package x11

import "testing"

func TestXmStringCreateLtoR(t *testing.T) {
	text := "hello world"
	xs := XmStringCreateLtoR(text, XmSTRING_DEFAULT_CHARSET)
	if xs.XmString == nil {
		t.Fatalf("XmStringCreateLtoR returned nil for %q", text)
	}

	// Ensure conversion to XtArgVal yields a non-zero value
	v := XtArgValFromXmString(xs)
	if v == 0 {
		t.Fatalf("XtArgValFromXmString returned zero ArgVal")
	}

	XmStringFree(xs)
}
