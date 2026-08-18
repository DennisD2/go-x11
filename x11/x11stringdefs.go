package x11

/* #cgo CFLAGS: -std=c99 -Wno-incompatible-pointer-types
#cgo LDFLAGS: -lXm -lXt -lX11
#include <stdint.h>
#include <stdlib.h>

// Xt includes
#include <X11/Intrinsic.h>

// Xm includes
#include <Xm/Xm.h>
#include "Xm/ArrowB.h"
#include "Xm/ArrowBG.h"
#include "Xm/BulletinB.h"
#include "Xm/CascadeB.h"
#include "Xm/CascadeBG.h"
#include "Xm/ClipWindowP.h"
#include "Xm/ComboBox.h"
#include "Xm/Command.h"
#include "Xm/DesktopP.h"
#include "Xm/DialogS.h"
#include "Xm/DialogSEP.h"
#include "Xm/DragC.h"
#include "Xm/DragIcon.h"
#include "Xm/DrawingA.h"
#include "Xm/DrawnB.h"
#include "Xm/DropDown.h"
#include "Xm/DropSMgr.h"
#include "Xm/DropTrans.h"
#include "Xm/Ext18List.h"
#include "Xm/ExtObjectP.h"
#include "Xm/FileSB.h"
#include "Xm/Form.h"
#include "Xm/Frame.h"
#include "Xm/Gadget.h"
#include "Xm/GrabShell.h"
#include "Xm/Label.h"
#include "Xm/LabelG.h"
#include "Xm/List.h"
#include "Xm/MainW.h"
#include "Xm/Manager.h"
#include "Xm/MenuShell.h"
#include "Xm/MessageB.h"
#include "Xm/MultiList.h"
#include "Xm/Notebook.h"
#include "Xm/PanedW.h"
#include "Xm/Primitive.h"
//#include "Xm/Print.h" xprint extension is required, no packages there today for it
#include "Xm/ProtocolsP.h"
#include "Xm/PushB.h"
#include "Xm/PushBG.h"
#include "Xm/RowColumn.h"
#include "Xm/SSpinB.h"
#include "Xm/SashP.h"
#include "Xm/Scale.h"
#include "Xm/ScrollBar.h"
#include "Xm/ScrolledW.h"
#include "Xm/SelectioB.h"
#include "Xm/SeparatoG.h"
#include "Xm/Separator.h"
#include "Xm/ShellEP.h"
#include "Xm/SpinB.h"
#include "Xm/TearOffBP.h"
#include "Xm/ToggleB.h"
#include "Xm/ToggleBG.h"
#include "Xm/VendorS.h"
#include "Xm/VendorSEP.h"


#include "cheader.h"
#include "wrapperInfo.h"

*/
import "C"

// String defines
var XtNaccelerators = C.GoString(C.XtNaccelerators)
var XtNallowHoriz = C.GoString(C.XtNallowHoriz)
var XtNallowVert = C.GoString(C.XtNallowVert)
var XtNancestorSensitive = C.GoString(C.XtNancestorSensitive)
var XtNbackground = C.GoString(C.XtNbackground)
var XtNbackgroundPixmap = C.GoString(C.XtNbackgroundPixmap)
var XtNbitmap = C.GoString(C.XtNbitmap)
var XtNborderColor = C.GoString(C.XtNborderColor)
var XtNborder = C.GoString(C.XtNborder)
var XtNborderPixmap = C.GoString(C.XtNborderPixmap)
var XtNborderWidth = C.GoString(C.XtNborderWidth)
var XtNcallback = C.GoString(C.XtNcallback)
var XtNchildren = C.GoString(C.XtNchildren)
var XtNcolormap = C.GoString(C.XtNcolormap)
var XtNdepth = C.GoString(C.XtNdepth)
var XtNdestroyCallback = C.GoString(C.XtNdestroyCallback)
var XtNeditType = C.GoString(C.XtNeditType)
var XtNfile = C.GoString(C.XtNfile)
var XtNfont = C.GoString(C.XtNfont)
var XtNforceBars = C.GoString(C.XtNforceBars)
var XtNforeground = C.GoString(C.XtNforeground)
var XtNfunction = C.GoString(C.XtNfunction)
var XtNheight = C.GoString(C.XtNheight)
var XtNhighlight = C.GoString(C.XtNhighlight)
var XtNhSpace = C.GoString(C.XtNhSpace)
var XtNindex = C.GoString(C.XtNindex)
var XtNinitialResourcesPersistent = C.GoString(C.XtNinitialResourcesPersistent)
var XtNinnerHeight = C.GoString(C.XtNinnerHeight)
var XtNinnerWidth = C.GoString(C.XtNinnerWidth)
var XtNinnerWindow = C.GoString(C.XtNinnerWindow)
var XtNinsertPosition = C.GoString(C.XtNinsertPosition)
var XtNinternalHeight = C.GoString(C.XtNinternalHeight)
var XtNinternalWidth = C.GoString(C.XtNinternalWidth)
var XtNjumpProc = C.GoString(C.XtNjumpProc)
var XtNjustify = C.GoString(C.XtNjustify)
var XtNknobHeight = C.GoString(C.XtNknobHeight)
var XtNknobIndent = C.GoString(C.XtNknobIndent)
var XtNknobPixel = C.GoString(C.XtNknobPixel)
var XtNknobWidth = C.GoString(C.XtNknobWidth)
var XtNlabel = C.GoString(C.XtNlabel)
var XtNlength = C.GoString(C.XtNlength)
var XtNlowerRight = C.GoString(C.XtNlowerRight)
var XtNmappedWhenManaged = C.GoString(C.XtNmappedWhenManaged)
var XtNmenuEntry = C.GoString(C.XtNmenuEntry)
var XtNname = C.GoString(C.XtNname)
var XtNnotify = C.GoString(C.XtNnotify)
var XtNnumChildren = C.GoString(C.XtNnumChildren)
var XtNorientation = C.GoString(C.XtNorientation)
var XtNparameter = C.GoString(C.XtNparameter)
var XtNpixmap = C.GoString(C.XtNpixmap)
var XtNpopupCallback = C.GoString(C.XtNpopupCallback)
var XtNpopdownCallback = C.GoString(C.XtNpopdownCallback)
var XtNresize = C.GoString(C.XtNresize)
var XtNreverseVideo = C.GoString(C.XtNreverseVideo)
var XtNscreen = C.GoString(C.XtNscreen)
var XtNscrollProc = C.GoString(C.XtNscrollProc)
var XtNscrollDCursor = C.GoString(C.XtNscrollDCursor)
var XtNscrollHCursor = C.GoString(C.XtNscrollHCursor)
var XtNscrollLCursor = C.GoString(C.XtNscrollLCursor)
var XtNscrollRCursor = C.GoString(C.XtNscrollRCursor)
var XtNscrollUCursor = C.GoString(C.XtNscrollUCursor)
var XtNscrollVCursor = C.GoString(C.XtNscrollVCursor)
var XtNselection = C.GoString(C.XtNselection)
var XtNselectionArray = C.GoString(C.XtNselectionArray)
var XtNsensitive = C.GoString(C.XtNsensitive)
var XtNshown = C.GoString(C.XtNshown)
var XtNspace = C.GoString(C.XtNspace)
var XtNstring = C.GoString(C.XtNstring)
var XtNtextOptions = C.GoString(C.XtNtextOptions)
var XtNtextSink = C.GoString(C.XtNtextSink)
var XtNtextSource = C.GoString(C.XtNtextSource)
var XtNthickness = C.GoString(C.XtNthickness)
var XtNthumb = C.GoString(C.XtNthumb)
var XtNthumbProc = C.GoString(C.XtNthumbProc)
var XtNtop = C.GoString(C.XtNtop)
var XtNtranslations = C.GoString(C.XtNtranslations)
var XtNunrealizeCallback = C.GoString(C.XtNunrealizeCallback)
var XtNupdate = C.GoString(C.XtNupdate)
var XtNuseBottom = C.GoString(C.XtNuseBottom)
var XtNuseRight = C.GoString(C.XtNuseRight)
var XtNvalue = C.GoString(C.XtNvalue)
var XtNvSpace = C.GoString(C.XtNvSpace)
var XtNwidth = C.GoString(C.XtNwidth)
var XtNwindow = C.GoString(C.XtNwindow)
var XtNx = C.GoString(C.XtNx)
var XtNy = C.GoString(C.XtNy)
var XtNfontSet = C.GoString(C.XtNfontSet)
var XtNcreateHook = C.GoString(C.XtNcreateHook)
var XtNchangeHook = C.GoString(C.XtNchangeHook)
var XtNconfigureHook = C.GoString(C.XtNconfigureHook)
var XtNgeometryHook = C.GoString(C.XtNgeometryHook)
var XtNdestroyHook = C.GoString(C.XtNdestroyHook)
var XtNshells = C.GoString(C.XtNshells)
var XtNnumShells = C.GoString(C.XtNnumShells)

// string defines
var XmNlabelString = C.GoString(C.XmNlabelString)

func ArrowButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmArrowButtonWidgetClass}
}
func ArrowButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmArrowButtonGadgetClass}
}
func BulletinBoardWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmBulletinBoardWidgetClass}
}
func CascadeButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmCascadeButtonWidgetClass}
}
func CascadeButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmCascadeButtonGadgetClass}
}
func ClipWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmClipWindowWidgetClass}
}
func ComboBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmComboBoxWidgetClass}
}
func CommandWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmCommandWidgetClass}
}
func DesktopClass() WidgetClass {
	return WidgetClass{c: C.xmDesktopClass}
}
func DialogShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDialogShellWidgetClass}
}
func DialogShellExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDialogShellExtObjectClass}
}
func DragContextClass() WidgetClass {
	return WidgetClass{c: C.xmDragContextClass}
}
func DragIconObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDragIconObjectClass}
}
func DrawingAreaWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDrawingAreaWidgetClass}
}
func DrawnButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDrawnButtonWidgetClass}
}
func DropDownWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmDropDownWidgetClass}
}
func DropSiteManagerObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDropSiteManagerObjectClass}
}
func DropTransferObjectClass() WidgetClass {
	return WidgetClass{c: C.xmDropTransferObjectClass}
}
func Ext18ListWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmExt18ListWidgetClass}
}
func ExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmExtObjectClass}
}
func FileSelectionBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmFileSelectionBoxWidgetClass}
}
func FormWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmFormWidgetClass}
}
func FrameWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmFrameWidgetClass}
}
func GadgetClass() WidgetClass {
	return WidgetClass{c: C.xmGadgetClass}
}
func GrabShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmGrabShellWidgetClass}
}
func LabelWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmLabelWidgetClass}
}
func LabelGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmLabelGadgetClass}
}
func ListWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmListWidgetClass}
}
func MainWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMainWindowWidgetClass}
}
func ManagerWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmManagerWidgetClass}
}
func MenuShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMenuShellWidgetClass}
}
func MessageBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMessageBoxWidgetClass}
}
func MultiListWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmMultiListWidgetClass}
}
func NotebookWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmNotebookWidgetClass}
}
func PanedWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPanedWindowWidgetClass}
}
func PrimitiveWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPrimitiveWidgetClass}
}

//	func PrintShellWidgetClass() WidgetClass {
//		return WidgetClass{c: C.xmPrintShellWidgetClass}
//	}

func ProtocolObjectClass() WidgetClass {
	return WidgetClass{c: C.xmProtocolObjectClass}
}
func PushButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmPushButtonWidgetClass}
}
func PushButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmPushButtonGadgetClass}
}
func RowColumnWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmRowColumnWidgetClass}
}
func SimpleSpinBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSimpleSpinBoxWidgetClass}
}
func SashWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSashWidgetClass}
}
func ScaleWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmScaleWidgetClass}
}
func ScrollBarWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmScrollBarWidgetClass}
}
func ScrolledWindowWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmScrolledWindowWidgetClass}
}
func SelectionBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSelectionBoxWidgetClass}
}
func SeparatorGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmSeparatorGadgetClass}
}
func SeparatorWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSeparatorWidgetClass}
}
func ShellExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmShellExtObjectClass}
}
func SpinBoxWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmSpinBoxWidgetClass}
}
func TearOffButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmTearOffButtonWidgetClass}
}
func ToggleButtonWidgetClass() WidgetClass {
	return WidgetClass{c: C.xmToggleButtonWidgetClass}
}
func ToggleButtonGadgetClass() WidgetClass {
	return WidgetClass{c: C.xmToggleButtonGadgetClass}
}
func vendorShellWidgetClass() WidgetClass {
	return WidgetClass{c: C.vendorShellWidgetClass}
}
func VendorShellExtObjectClass() WidgetClass {
	return WidgetClass{c: C.xmVendorShellExtObjectClass}
}
