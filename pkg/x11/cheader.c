
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

Widget call_XtAppCreateShell(
    char *application_name,
    char *application_class,
    WidgetClass widget_class,
    Display *display,
    void *args, Cardinal num_args
) {
    return XtAppCreateShell(application_name, application_class, widget_class, display,
        (ArgList)args, num_args);
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

// action handling

// Declaration of global C->Go bridge for actions
extern void goActionBridgeWithId(Widget w, XEvent* event, String* params, Cardinal* num_params, int actionId);

// Macro to generate a function, usable by C C-side Action handler. Function name is also generated and contains ID
#define ACTION_BRIDGE(id) \
    static void c_action_bridge_##id(Widget w, XEvent* ev, String* p, Cardinal* n) { \
        goActionBridgeWithId(w, ev, p, n, id); \
    }

// A pool of bridge functions callable by C. These will call the go entry function with their ID
ACTION_BRIDGE(0) ACTION_BRIDGE(1) ACTION_BRIDGE(2) ACTION_BRIDGE(3) ACTION_BRIDGE(4)
ACTION_BRIDGE(5) ACTION_BRIDGE(6) ACTION_BRIDGE(7) ACTION_BRIDGE(8) ACTION_BRIDGE(9)

// returns Go function by ID
XtActionProc get_bridge_ptr(int id) {
    switch(id) {
        case 0: return c_action_bridge_0;
        case 1: return c_action_bridge_1;
        case 2: return c_action_bridge_2;
        case 3: return c_action_bridge_3;
        case 4: return c_action_bridge_4;
        case 5: return c_action_bridge_5;
        case 6: return c_action_bridge_6;
        case 7: return c_action_bridge_7;
        case 8: return c_action_bridge_8;
        case 9: return c_action_bridge_9;
        default: return NULL;
    }
}

// Code for XtActionProc handling
void set_c_action_entry(XtActionsRec *table, int index, const char *name, XtActionProc proc) {
    table[index].string = (String)name;
    table[index].proc = proc;
}


// Code for Xt XEventHandler

void call_XtAddEventHandler(Widget w, EventMask mask, Boolean non_maskable, XtEventHandler proc,
    XtPointer client_data) {
        XtAddEventHandler(w, mask, non_maskable, proc, client_data);
}

// void* args
void call_XtGetApplicationResources(Widget w, XtPointer base, XtResourceList resources, Cardinal num_resources,
    void *args, Cardinal num_args) {
    XtGetApplicationResources(w,
        base,
        resources, num_resources,
        (ArgList)args, num_args
    );
}

// void* converter, void *convert_args
void call_XtAddConverter(_XtString from_type, _XtString to_type, void *converter,
    void *convert_args, Cardinal num_args) {
    XtAddConverter(from_type, to_type, converter, convert_args, num_args);
}
