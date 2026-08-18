
#include <stdlib.h>
#include <X11/Intrinsic.h>
#include <Xm/Xm.h>
#include <Xm/Label.h>

#include "cheader.h"
#include "_cgo_export.h"

extern Widget call_XtInitialize(
    char *shell_name,
    char *application_class,
    void *options, Cardinal num_options,
    int *argc_in_out, void *argv_in_out
) {
    return XtInitialize(
            (String)shell_name,
            (String)application_class,
            (XrmOptionDescList)options, num_options,
            argc_in_out, (String*)argv_in_out
        );
    }

// options: void* -> XrmOptionDescList
// argv_in_out: void* -> String*
// fallback_resources: void* -> String*
// args: void* -> ArgList
Widget call_XtAppInitialize(
    XtAppContext *app_context_return,
    char *application_class,
    void *options, Cardinal num_options,
    int *argc_in_out, void *argv_in_out,
    void *fallback_resources,
    void *args, Cardinal num_args
) {
    return XtAppInitialize(
        app_context_return,
        (String)application_class,
        (XrmOptionDescList)options, num_options,
        argc_in_out, (String*)argv_in_out,
        (String*)fallback_resources,
        (ArgList)args, num_args
    );
}

// args: void* -> ArgList
Widget call_XtCreateManagedWidget(
    char *name,
    WidgetClass widget_class,
    Widget parent,
    void *args, Cardinal num_args
) {
    return XtCreateManagedWidget(name, widget_class, parent, (ArgList)args, num_args);
}

void set_option_rec(void *options_base, int index, char *opt, char *spec, int kind, char *val) {
    XrmOptionDescList list = (XrmOptionDescList)options_base;
    list[index].option = opt;
    list[index].specifier = spec;
    list[index].argKind = (XrmOptionKind)kind;
    list[index].value = (XPointer)val;
}

// callback handling
extern void goCallbackDispatcher(Widget w, XtPointer client_data, XtPointer call_data);

void call_XtAddCallback(Widget w, char *name, uintptr_t go_func_id) {
    XtAddCallback(w, name, (XtCallbackProc)goCallbackDispatcher, (XtPointer)go_func_id);
}

void call_XtManageChildren(Widget **widgets, int num_widgets) {
    XtManageChildren((WidgetList)widgets, (Cardinal)num_widgets);
}

void call_XtSetValues( Widget w, void *args, Cardinal n) {
    XtSetValues(w, args, (Cardinal)n );
}
