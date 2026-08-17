
#include <X11/Intrinsic.h>
#include <Xm/Xm.h>

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

extern void set_option_rec(void *options_base, int index, char *opt, char *spec, int kind, char *val);

// callback handling
extern void goCallbackDispatcher(Widget w, XtPointer client_data, XtPointer call_data);

extern void call_XtAddCallback(Widget w, char *name, uintptr_t go_func_id) ;