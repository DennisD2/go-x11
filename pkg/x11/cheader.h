
#include <stdint.h>
#include <X11/Intrinsic.h>
#include <Xm/Xm.h>

// options: void* -> XrmOptionDescList
// argv_in_out: void* -> String*
extern Widget call_XtInitialize(
    char *shell_name,
    char *application_class,
    void *options, Cardinal num_options,
    int *argc_in_out, void *argv_in_out
);

// options: void* -> XrmOptionDescList
// argv_in_out: void* -> String*
// fallback_resources: void* -> String*
// args: void* -> ArgList
extern Widget call_XtAppInitialize(
    XtAppContext *app_context_return,
    char *application_class,
    void *options, Cardinal num_options,
    int *argc_in_out, void *argv_in_out,
    void *fallback_resources,
    void *args, Cardinal num_args
);

// args: void* -> ArgList
extern Widget call_XtCreateManagedWidget(
    char *name,
    WidgetClass widget_class,
    Widget parent,
    void *args, Cardinal num_args
);

// args: void* -> ArgList
extern Widget call_XtAppCreateShell(
    char *application_name,
    char *application_class,
    WidgetClass widget_class,
    Display *display,
    void *args, Cardinal num_args
);


extern void set_option_rec(void *options_base, int index, char *opt, char *spec, int kind, char *val);

// callback handling
extern void goCallbackDispatcher(Widget w, XtPointer client_data, XtPointer call_data);

extern void call_XtAddCallback(Widget w, char *name, uintptr_t go_func_id) ;

// Action handling
extern XtActionProc get_bridge_ptr(int id);

void set_c_action_entry(XtActionsRec *table, int index, const char *name, XtActionProc proc);

// event handling
// Declaration of global C->Go bridge for actions
extern void goEventHandlerBridge(Widget w, XtPointer client_data, XEvent* event, Boolean* continue_to_dispatch);

extern void call_XtAddEventHandler(Widget w, EventMask mask, Boolean non_maskable,
    XtEventHandler proc, XtPointer client_data);

extern void call_XtAddEventHandler(Widget w, EventMask mask, Boolean non_maskable, XtEventHandler proc,
    XtPointer client_data);

// Other functions
extern void call_XtManageChildren(Widget **widgets, int num_widgets);

extern void call_XtSetValues( Widget w, void *args, Cardinal n);
