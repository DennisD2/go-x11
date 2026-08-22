package x11

/*
#include <stdint.h>
#include <stdlib.h>

// Xm includes
#include <Xm/XmStrDefs.h>
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
*/
import "C"
import "unsafe"

// WidgetClass defines - need to be functions

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

// Defines for Xm widgets

var XmNaccelerator = C.GoString(C.XmNaccelerator)
var XmNacceleratorText = C.GoString(C.XmNacceleratorText)
var XmNactivateCallback = C.GoString(C.XmNactivateCallback)
var XmNadjustLast = C.GoString(C.XmNadjustLast)
var XmNadjustMargin = C.GoString(C.XmNadjustMargin)
var XmNalignment = C.GoString(C.XmNalignment)
var XmNallowOverlap = C.GoString(C.XmNallowOverlap)
var XmNallowResize = C.GoString(C.XmNallowResize)
var XmNallowUnusedSpace = C.GoString(C.XmNallowUnusedSpace)
var XmNanimationMask = C.GoString(C.XmNanimationMask)
var XmNanimationPixmap = C.GoString(C.XmNanimationPixmap)
var XmNanimationPixmapDepth = C.GoString(C.XmNanimationPixmapDepth)
var XmNanimationStyle = C.GoString(C.XmNanimationStyle)
var XmNapplyCallback = C.GoString(C.XmNapplyCallback)
var XmNapplyLabelString = C.GoString(C.XmNapplyLabelString)
var XmNarmCallback = C.GoString(C.XmNarmCallback)
var XmNarmColor = C.GoString(C.XmNarmColor)
var XmNarmPixmap = C.GoString(C.XmNarmPixmap)
var XmNarrowDirection = C.GoString(C.XmNarrowDirection)
var XmNattachment = C.GoString(C.XmNattachment)
var XmNaudibleWarning = C.GoString(C.XmNaudibleWarning)
var XmNautoShowCursorPosition = C.GoString(C.XmNautoShowCursorPosition)
var XmNautoUnmanage = C.GoString(C.XmNautoUnmanage)
var XmNautomaticSelection = C.GoString(C.XmNautomaticSelection)
var XmNavailability = C.GoString(C.XmNavailability)
var XmNblendModel = C.GoString(C.XmNblendModel)
var XmNblinkRate = C.GoString(C.XmNblinkRate)
var XmNbottomAttachment = C.GoString(C.XmNbottomAttachment)
var XmNbottomOffset = C.GoString(C.XmNbottomOffset)
var XmNbottomPosition = C.GoString(C.XmNbottomPosition)
var XmNbottomShadowColor = C.GoString(C.XmNbottomShadowColor)
var XmNbottomShadowPixmap = C.GoString(C.XmNbottomShadowPixmap)
var XmNbottomWidget = C.GoString(C.XmNbottomWidget)
var XmNbrowseSelectionCallback = C.GoString(C.XmNbrowseSelectionCallback)
var XmNbuttonAcceleratorText = C.GoString(C.XmNbuttonAcceleratorText)
var XmNbuttonAccelerators = C.GoString(C.XmNbuttonAccelerators)
var XmNbuttonCount = C.GoString(C.XmNbuttonCount)
var XmNbuttonFontList = C.GoString(C.XmNbuttonFontList)
var XmNbuttonMnemonicCharSets = C.GoString(C.XmNbuttonMnemonicCharSets)
var XmNbuttonMnemonics = C.GoString(C.XmNbuttonMnemonics)
var XmNbuttonSet = C.GoString(C.XmNbuttonSet)
var XmNbuttonType = C.GoString(C.XmNbuttonType)
var XmNbuttons = C.GoString(C.XmNbuttons)
var XmNcancelButton = C.GoString(C.XmNcancelButton)
var XmNcancelCallback = C.GoString(C.XmNcancelCallback)
var XmNcancelLabelString = C.GoString(C.XmNcancelLabelString)
var XmNcascadePixmap = C.GoString(C.XmNcascadePixmap)
var XmNcascadingCallback = C.GoString(C.XmNcascadingCallback)
var XmNchildHorizontalAlignment = C.GoString(C.XmNchildHorizontalAlignment)
var XmNchildHorizontalSpacing = C.GoString(C.XmNchildHorizontalSpacing)
var XmNchildPlacement = C.GoString(C.XmNchildPlacement)
var XmNchildPosition = C.GoString(C.XmNchildPosition)
var XmNchildType = C.GoString(C.XmNchildType)
var XmNchildVerticalAlignment = C.GoString(C.XmNchildVerticalAlignment)
var XmNclientData = C.GoString(C.XmNclientData)
var XmNclipWindow = C.GoString(C.XmNclipWindow)
var XmNcolumns = C.GoString(C.XmNcolumns)
var XmNcommand = C.GoString(C.XmNcommand)
var XmNcommandChangedCallback = C.GoString(C.XmNcommandChangedCallback)
var XmNcommandEnteredCallback = C.GoString(C.XmNcommandEnteredCallback)
var XmNcommandWindow = C.GoString(C.XmNcommandWindow)
var XmNcommandWindowLocation = C.GoString(C.XmNcommandWindowLocation)
var XmNconvertProc = C.GoString(C.XmNconvertProc)
var XmNcursorBackground = C.GoString(C.XmNcursorBackground)
var XmNcursorForeground = C.GoString(C.XmNcursorForeground)
var XmNcursorPosition = C.GoString(C.XmNcursorPosition)
var XmNcursorPositionVisible = C.GoString(C.XmNcursorPositionVisible)
var XmNdarkThreshold = C.GoString(C.XmNdarkThreshold)
var XmNdecimalPoints = C.GoString(C.XmNdecimalPoints)
var XmNdecrementCallback = C.GoString(C.XmNdecrementCallback)
var XmNdefaultActionCallback = C.GoString(C.XmNdefaultActionCallback)
var XmNdefaultButton = C.GoString(C.XmNdefaultButton)
var XmNdefaultButtonShadowThickness = C.GoString(C.XmNdefaultButtonShadowThickness)
var XmNdefaultButtonType = C.GoString(C.XmNdefaultButtonType)
var XmNdefaultCopyCursorIcon = C.GoString(C.XmNdefaultCopyCursorIcon)
var XmNdefaultFontList = C.GoString(C.XmNdefaultFontList)
var XmNdefaultInvalidCursorIcon = C.GoString(C.XmNdefaultInvalidCursorIcon)
var XmNdefaultLinkCursorIcon = C.GoString(C.XmNdefaultLinkCursorIcon)
var XmNdefaultMoveCursorIcon = C.GoString(C.XmNdefaultMoveCursorIcon)
var XmNdefaultNoneCursorIcon = C.GoString(C.XmNdefaultNoneCursorIcon)
var XmNdefaultPosition = C.GoString(C.XmNdefaultPosition)
var XmNdefaultSourceCursorIcon = C.GoString(C.XmNdefaultSourceCursorIcon)
var XmNdefaultValidCursorIcon = C.GoString(C.XmNdefaultValidCursorIcon)
var XmNdeleteResponse = C.GoString(C.XmNdeleteResponse)
var XmNdesktopParent = C.GoString(C.XmNdesktopParent)
var XmNdialogStyle = C.GoString(C.XmNdialogStyle)
var XmNdialogTitle = C.GoString(C.XmNdialogTitle)
var XmNdialogType = C.GoString(C.XmNdialogType)
var XmNdirListItemCount = C.GoString(C.XmNdirListItemCount)
var XmNdirListItems = C.GoString(C.XmNdirListItems)
var XmNdirListLabelString = C.GoString(C.XmNdirListLabelString)
var XmNdirMask = C.GoString(C.XmNdirMask)
var XmNdirSearchProc = C.GoString(C.XmNdirSearchProc)
var XmNdirSpec = C.GoString(C.XmNdirSpec)
var XmNdirectory = C.GoString(C.XmNdirectory)
var XmNdirectoryValid = C.GoString(C.XmNdirectoryValid)
var XmNdisarmCallback = C.GoString(C.XmNdisarmCallback)
var XmNdoubleClickInterval = C.GoString(C.XmNdoubleClickInterval)
var XmNdragCallback = C.GoString(C.XmNdragCallback)
var XmNdragContextClass = C.GoString(C.XmNdragContextClass)
var XmNdragDropFinishCallback = C.GoString(C.XmNdragDropFinishCallback)
var XmNdragIconClass = C.GoString(C.XmNdragIconClass)
var XmNdragInitiatorProtocolStyle = C.GoString(C.XmNdragInitiatorProtocolStyle)
var XmNdragMotionCallback = C.GoString(C.XmNdragMotionCallback)
var XmNdragOperations = C.GoString(C.XmNdragOperations)
var XmNdragOverMode = C.GoString(C.XmNdragOverMode)
var XmNdragProc = C.GoString(C.XmNdragProc)
var XmNdragReceiverProtocolStyle = C.GoString(C.XmNdragReceiverProtocolStyle)
var XmNdropFinishCallback = C.GoString(C.XmNdropFinishCallback)
var XmNdropProc = C.GoString(C.XmNdropProc)
var XmNdropRectangles = C.GoString(C.XmNdropRectangles)
var XmNdropSiteActivity = C.GoString(C.XmNdropSiteActivity)
var XmNdropSiteEnterCallback = C.GoString(C.XmNdropSiteEnterCallback)
var XmNdropSiteLeaveCallback = C.GoString(C.XmNdropSiteLeaveCallback)
var XmNdropSiteManagerClass = C.GoString(C.XmNdropSiteManagerClass)
var XmNdropSiteOperations = C.GoString(C.XmNdropSiteOperations)
var XmNdropSiteType = C.GoString(C.XmNdropSiteType)
var XmNdropStartCallback = C.GoString(C.XmNdropStartCallback)
var XmNdropTransferClass = C.GoString(C.XmNdropTransferClass)
var XmNdropTransfers = C.GoString(C.XmNdropTransfers)
var XmNeditMode = C.GoString(C.XmNeditMode)
var XmNeditable = C.GoString(C.XmNeditable)
var XmNentryAlignment = C.GoString(C.XmNentryAlignment)
var XmNentryBorder = C.GoString(C.XmNentryBorder)
var XmNentryCallback = C.GoString(C.XmNentryCallback)
var XmNentryClass = C.GoString(C.XmNentryClass)
var XmNentryVerticalAlignment = C.GoString(C.XmNentryVerticalAlignment)
var XmNexportTargets = C.GoString(C.XmNexportTargets)
var XmNexposeCallback = C.GoString(C.XmNexposeCallback)
var XmNextendedSelectionCallback = C.GoString(C.XmNextendedSelectionCallback)
var XmNextensionType = C.GoString(C.XmNextensionType)
var XmNfileListItemCount = C.GoString(C.XmNfileListItemCount)
var XmNfileListItems = C.GoString(C.XmNfileListItems)
var XmNfileListLabelString = C.GoString(C.XmNfileListLabelString)
var XmNfileSearchProc = C.GoString(C.XmNfileSearchProc)
var XmNfileTypeMask = C.GoString(C.XmNfileTypeMask)
var XmNfillOnArm = C.GoString(C.XmNfillOnArm)
var XmNfillOnSelect = C.GoString(C.XmNfillOnSelect)
var XmNfilterLabelString = C.GoString(C.XmNfilterLabelString)
var XmNfocusCallback = C.GoString(C.XmNfocusCallback)
var XmNfocusMovedCallback = C.GoString(C.XmNfocusMovedCallback)
var XmNfocusPolicyChanged = C.GoString(C.XmNfocusPolicyChanged)
var XmNfontList = C.GoString(C.XmNfontList)
var XmNforegroundThreshold = C.GoString(C.XmNforegroundThreshold)
var XmNfractionBase = C.GoString(C.XmNfractionBase)
var XmNgainPrimaryCallback = C.GoString(C.XmNgainPrimaryCallback)
var XmNhelpCallback = C.GoString(C.XmNhelpCallback)
var XmNhelpLabelString = C.GoString(C.XmNhelpLabelString)
var XmNhighlightColor = C.GoString(C.XmNhighlightColor)
var XmNhighlightOnEnter = C.GoString(C.XmNhighlightOnEnter)
var XmNhighlightPixmap = C.GoString(C.XmNhighlightPixmap)
var XmNhighlightThickness = C.GoString(C.XmNhighlightThickness)
var XmNhistoryItemCount = C.GoString(C.XmNhistoryItemCount)
var XmNhistoryItems = C.GoString(C.XmNhistoryItems)
var XmNhistoryMaxItems = C.GoString(C.XmNhistoryMaxItems)
var XmNhistoryVisibleItemCount = C.GoString(C.XmNhistoryVisibleItemCount)
var XmNhorizontalFontUnit = C.GoString(C.XmNhorizontalFontUnit)
var XmNhorizontalScrollBar = C.GoString(C.XmNhorizontalScrollBar)
var XmNhorizontalSpacing = C.GoString(C.XmNhorizontalSpacing)
var XmNhotX = C.GoString(C.XmNhotX)
var XmNhotY = C.GoString(C.XmNhotY)
var XmNiccHandle = C.GoString(C.XmNiccHandle)
var XmNimportTargets = C.GoString(C.XmNimportTargets)
var XmNincrement = C.GoString(C.XmNincrement)
var XmNincrementCallback = C.GoString(C.XmNincrementCallback)
var XmNincremental = C.GoString(C.XmNincremental)
var XmNindicatorOn = C.GoString(C.XmNindicatorOn)
var XmNindicatorSize = C.GoString(C.XmNindicatorSize)
var XmNindicatorType = C.GoString(C.XmNindicatorType)
var XmNinitialDelay = C.GoString(C.XmNinitialDelay)
var XmNinitialFocus = C.GoString(C.XmNinitialFocus)
var XmNinputCallback = C.GoString(C.XmNinputCallback)
var XmNinputCreate = C.GoString(C.XmNinputCreate)
var XmNinputMethod = C.GoString(C.XmNinputMethod)
var XmNinvalidCursorForeground = C.GoString(C.XmNinvalidCursorForeground)
var XmNisAligned = C.GoString(C.XmNisAligned)
var XmNisHomogeneous = C.GoString(C.XmNisHomogeneous)
var XmNitemCount = C.GoString(C.XmNitemCount)
var XmNitems = C.GoString(C.XmNitems)
var XmNkeyboardFocusPolicy = C.GoString(C.XmNkeyboardFocusPolicy)
var XmNlabelFontList = C.GoString(C.XmNlabelFontList)
var XmNlabelInsensitivePixmap = C.GoString(C.XmNlabelInsensitivePixmap)
var XmNlabelPixmap = C.GoString(C.XmNlabelPixmap)
var XmNlabelString = C.GoString(C.XmNlabelString)
var XmNlabelType = C.GoString(C.XmNlabelType)
var XmNleftAttachment = C.GoString(C.XmNleftAttachment)
var XmNleftOffset = C.GoString(C.XmNleftOffset)
var XmNleftPosition = C.GoString(C.XmNleftPosition)
var XmNleftWidget = C.GoString(C.XmNleftWidget)
var XmNlightThreshold = C.GoString(C.XmNlightThreshold)
var XmNlineSpace = C.GoString(C.XmNlineSpace)
var XmNlistItemCount = C.GoString(C.XmNlistItemCount)
var XmNlistItems = C.GoString(C.XmNlistItems)
var XmNlistLabelString = C.GoString(C.XmNlistLabelString)
var XmNlistMarginHeight = C.GoString(C.XmNlistMarginHeight)
var XmNlistMarginWidth = C.GoString(C.XmNlistMarginWidth)
var XmNlistSizePolicy = C.GoString(C.XmNlistSizePolicy)
var XmNlistSpacing = C.GoString(C.XmNlistSpacing)
var XmNlistUpdated = C.GoString(C.XmNlistUpdated)
var XmNlistVisibleItemCount = C.GoString(C.XmNlistVisibleItemCount)
var XmNlogicalParent = C.GoString(C.XmNlogicalParent)
var XmNlosePrimaryCallback = C.GoString(C.XmNlosePrimaryCallback)
var XmNlosingFocusCallback = C.GoString(C.XmNlosingFocusCallback)
var XmNmainWindowMarginHeight = C.GoString(C.XmNmainWindowMarginHeight)
var XmNmainWindowMarginWidth = C.GoString(C.XmNmainWindowMarginWidth)
var XmNmapCallback = C.GoString(C.XmNmapCallback)
var XmNmappingDelay = C.GoString(C.XmNmappingDelay)
var XmNmargin = C.GoString(C.XmNmargin)
var XmNmarginBottom = C.GoString(C.XmNmarginBottom)
var XmNmarginHeight = C.GoString(C.XmNmarginHeight)
var XmNmarginLeft = C.GoString(C.XmNmarginLeft)
var XmNmarginRight = C.GoString(C.XmNmarginRight)
var XmNmarginTop = C.GoString(C.XmNmarginTop)
var XmNmarginWidth = C.GoString(C.XmNmarginWidth)
var XmNmask = C.GoString(C.XmNmask)
var XmNmaxLength = C.GoString(C.XmNmaxLength)
var XmNmaximum = C.GoString(C.XmNmaximum)
var XmNmenuAccelerator = C.GoString(C.XmNmenuAccelerator)
var XmNmenuBar = C.GoString(C.XmNmenuBar)
var XmNmenuCursor = C.GoString(C.XmNmenuCursor)
var XmNmenuHelpWidget = C.GoString(C.XmNmenuHelpWidget)
var XmNmenuHistory = C.GoString(C.XmNmenuHistory)
var XmNmenuPost = C.GoString(C.XmNmenuPost)
var XmNmessageAlignment = C.GoString(C.XmNmessageAlignment)
var XmNmessageProc = C.GoString(C.XmNmessageProc)
var XmNmessageString = C.GoString(C.XmNmessageString)
var XmNmessageWindow = C.GoString(C.XmNmessageWindow)
var XmNminimizeButtons = C.GoString(C.XmNminimizeButtons)
var XmNminimum = C.GoString(C.XmNminimum)
var XmNmnemonic = C.GoString(C.XmNmnemonic)
var XmNmnemonicCharSet = C.GoString(C.XmNmnemonicCharSet)
var XmNmodifyVerifyCallback = C.GoString(C.XmNmodifyVerifyCallback)
var XmNmodifyVerifyCallbackWcs = C.GoString(C.XmNmodifyVerifyCallbackWcs)
var XmNmotionVerifyCallback = C.GoString(C.XmNmotionVerifyCallback)
var XmNmoveOpaque = C.GoString(C.XmNmoveOpaque)
var XmNmultiClick = C.GoString(C.XmNmultiClick)
var XmNmultipleSelectionCallback = C.GoString(C.XmNmultipleSelectionCallback)
var XmNmustMatch = C.GoString(C.XmNmustMatch)
var XmNmwmDecorations = C.GoString(C.XmNmwmDecorations)
var XmNmwmFunctions = C.GoString(C.XmNmwmFunctions)
var XmNmwmInputMode = C.GoString(C.XmNmwmInputMode)
var XmNmwmMenu = C.GoString(C.XmNmwmMenu)
var XmNmwmMessages = C.GoString(C.XmNmwmMessages)
var XmNnavigationType = C.GoString(C.XmNnavigationType)
var XmNneedsMotion = C.GoString(C.XmNneedsMotion)
var XmNnoMatchCallback = C.GoString(C.XmNnoMatchCallback)
var XmNnoMatchString = C.GoString(C.XmNnoMatchString)
var XmNnoResize = C.GoString(C.XmNnoResize)
var XmNnoneCursorForeground = C.GoString(C.XmNnoneCursorForeground)
var XmNnotifyProc = C.GoString(C.XmNnotifyProc)
var XmNnumColumns = C.GoString(C.XmNnumColumns)
var XmNnumDropRectangles = C.GoString(C.XmNnumDropRectangles)
var XmNnumDropTransfers = C.GoString(C.XmNnumDropTransfers)
var XmNnumExportTargets = C.GoString(C.XmNnumExportTargets)
var XmNnumImportTargets = C.GoString(C.XmNnumImportTargets)
var XmNnumRectangles = C.GoString(C.XmNnumRectangles)
var XmNoffsetX = C.GoString(C.XmNoffsetX)
var XmNoffsetY = C.GoString(C.XmNoffsetY)
var XmNokCallback = C.GoString(C.XmNokCallback)
var XmNokLabelString = C.GoString(C.XmNokLabelString)
var XmNoperationChangedCallback = C.GoString(C.XmNoperationChangedCallback)
var XmNoperationCursorIcon = C.GoString(C.XmNoperationCursorIcon)
var XmNoptionLabel = C.GoString(C.XmNoptionLabel)
var XmNoptionMnemonic = C.GoString(C.XmNoptionMnemonic)
var XmNoutputCreate = C.GoString(C.XmNoutputCreate)
var XmNpacking = C.GoString(C.XmNpacking)
var XmNpageDecrementCallback = C.GoString(C.XmNpageDecrementCallback)
var XmNpageIncrement = C.GoString(C.XmNpageIncrement)
var XmNpageIncrementCallback = C.GoString(C.XmNpageIncrementCallback)
var XmNpaneMaximum = C.GoString(C.XmNpaneMaximum)
var XmNpaneMinimum = C.GoString(C.XmNpaneMinimum)
var XmNpattern = C.GoString(C.XmNpattern)
var XmNpendingDelete = C.GoString(C.XmNpendingDelete)
var XmNpopupEnabled = C.GoString(C.XmNpopupEnabled)
var XmNpositionIndex = C.GoString(C.XmNpositionIndex)
var XmNpostFromButton = C.GoString(C.XmNpostFromButton)
var XmNpostFromCount = C.GoString(C.XmNpostFromCount)
var XmNpostFromList = C.GoString(C.XmNpostFromList)
var XmNpreeditType = C.GoString(C.XmNpreeditType)
var XmNprocessingDirection = C.GoString(C.XmNprocessingDirection)
var XmNpromptString = C.GoString(C.XmNpromptString)
var XmNprotocolCallback = C.GoString(C.XmNprotocolCallback)
var XmNpushButtonEnabled = C.GoString(C.XmNpushButtonEnabled)
var XmNqualifySearchDataProc = C.GoString(C.XmNqualifySearchDataProc)
var XmNradioAlwaysOne = C.GoString(C.XmNradioAlwaysOne)
var XmNradioBehavior = C.GoString(C.XmNradioBehavior)
var XmNrealizeCallback = C.GoString(C.XmNrealizeCallback)
var XmNrecomputeSize = C.GoString(C.XmNrecomputeSize)
var XmNrectangles = C.GoString(C.XmNrectangles)
var XmNrefigureMode = C.GoString(C.XmNrefigureMode)
var XmNrepeatDelay = C.GoString(C.XmNrepeatDelay)
var XmNresizable = C.GoString(C.XmNresizable)
var XmNresizeCallback = C.GoString(C.XmNresizeCallback)
var XmNresizeHeight = C.GoString(C.XmNresizeHeight)
var XmNresizePolicy = C.GoString(C.XmNresizePolicy)
var XmNresizeWidth = C.GoString(C.XmNresizeWidth)
var XmNrightAttachment = C.GoString(C.XmNrightAttachment)
var XmNrightOffset = C.GoString(C.XmNrightOffset)
var XmNrightPosition = C.GoString(C.XmNrightPosition)
var XmNrightWidget = C.GoString(C.XmNrightWidget)
var XmNrowColumnType = C.GoString(C.XmNrowColumnType)
var XmNrows = C.GoString(C.XmNrows)
var XmNrubberPositioning = C.GoString(C.XmNrubberPositioning)
var XmNsashHeight = C.GoString(C.XmNsashHeight)
var XmNsashIndent = C.GoString(C.XmNsashIndent)
var XmNsashShadowThickness = C.GoString(C.XmNsashShadowThickness)
var XmNsashWidth = C.GoString(C.XmNsashWidth)
var XmNscaleHeight = C.GoString(C.XmNscaleHeight)
var XmNscaleMultiple = C.GoString(C.XmNscaleMultiple)
var XmNscaleWidth = C.GoString(C.XmNscaleWidth)
var XmNscrollBarDisplayPolicy = C.GoString(C.XmNscrollBarDisplayPolicy)
var XmNscrollBarPlacement = C.GoString(C.XmNscrollBarPlacement)
var XmNscrollHorizontal = C.GoString(C.XmNscrollHorizontal)
var XmNscrollLeftSide = C.GoString(C.XmNscrollLeftSide)
var XmNscrollTopSide = C.GoString(C.XmNscrollTopSide)
var XmNscrollVertical = C.GoString(C.XmNscrollVertical)
var XmNscrolledWindowMarginHeight = C.GoString(C.XmNscrolledWindowMarginHeight)
var XmNscrolledWindowMarginWidth = C.GoString(C.XmNscrolledWindowMarginWidth)
var XmNscrollingPolicy = C.GoString(C.XmNscrollingPolicy)
var XmNselectColor = C.GoString(C.XmNselectColor)
var XmNselectInsensitivePixmap = C.GoString(C.XmNselectInsensitivePixmap)
var XmNselectPixmap = C.GoString(C.XmNselectPixmap)
var XmNselectThreshold = C.GoString(C.XmNselectThreshold)
var XmNselectedItemCount = C.GoString(C.XmNselectedItemCount)
var XmNselectedItems = C.GoString(C.XmNselectedItems)
var XmNselectionArrayCount = C.GoString(C.XmNselectionArrayCount)
var XmNselectionLabelString = C.GoString(C.XmNselectionLabelString)
var XmNselectionPolicy = C.GoString(C.XmNselectionPolicy)
var XmNseparatorOn = C.GoString(C.XmNseparatorOn)
var XmNseparatorType = C.GoString(C.XmNseparatorType)
var XmNset = C.GoString(C.XmNset)
var XmNshadow = C.GoString(C.XmNshadow)
var XmNshadowThickness = C.GoString(C.XmNshadowThickness)
var XmNshadowType = C.GoString(C.XmNshadowType)
var XmNshellUnitType = C.GoString(C.XmNshellUnitType)
var XmNshowArrows = C.GoString(C.XmNshowArrows)
var XmNshowAsDefault = C.GoString(C.XmNshowAsDefault)
var XmNshowSeparator = C.GoString(C.XmNshowSeparator)
var XmNshowValue = C.GoString(C.XmNshowValue)
var XmNsimpleCallback = C.GoString(C.XmNsimpleCallback)
var XmNsingleSelectionCallback = C.GoString(C.XmNsingleSelectionCallback)
var XmNsizePolicy = C.GoString(C.XmNsizePolicy)
var XmNskipAdjust = C.GoString(C.XmNskipAdjust)
var XmNsliderSize = C.GoString(C.XmNsliderSize)
var XmNsource = C.GoString(C.XmNsource)
var XmNsourceCursorIcon = C.GoString(C.XmNsourceCursorIcon)
var XmNsourceIsExternal = C.GoString(C.XmNsourceIsExternal)
var XmNsourcePixmapIcon = C.GoString(C.XmNsourcePixmapIcon)
var XmNsourceWidget = C.GoString(C.XmNsourceWidget)
var XmNsourceWindow = C.GoString(C.XmNsourceWindow)
var XmNspacing = C.GoString(C.XmNspacing)
var XmNspotLocation = C.GoString(C.XmNspotLocation)
var XmNstartTime = C.GoString(C.XmNstartTime)
var XmNstateCursorIcon = C.GoString(C.XmNstateCursorIcon)
var XmNstringDirection = C.GoString(C.XmNstringDirection)
var XmNsubMenuId = C.GoString(C.XmNsubMenuId)
var XmNsymbolPixmap = C.GoString(C.XmNsymbolPixmap)
var XmNtearOffMenuActivateCallback = C.GoString(C.XmNtearOffMenuActivateCallback)
var XmNtearOffMenuDeactivateCallback = C.GoString(C.XmNtearOffMenuDeactivateCallback)
var XmNtearOffModel = C.GoString(C.XmNtearOffModel)
var XmNtextAccelerators = C.GoString(C.XmNtextAccelerators)
var XmNtextColumns = C.GoString(C.XmNtextColumns)
var XmNtextFontList = C.GoString(C.XmNtextFontList)
var XmNtextString = C.GoString(C.XmNtextString)
var XmNtextTranslations = C.GoString(C.XmNtextTranslations)
var XmNtextValue = C.GoString(C.XmNtextValue)
var XmNtitleString = C.GoString(C.XmNtitleString)
var XmNtoBottomCallback = C.GoString(C.XmNtoBottomCallback)
var XmNtoPositionCallback = C.GoString(C.XmNtoPositionCallback)
var XmNtoTopCallback = C.GoString(C.XmNtoTopCallback)
var XmNtopAttachment = C.GoString(C.XmNtopAttachment)
var XmNtopCharacter = C.GoString(C.XmNtopCharacter)
var XmNtopItemPosition = C.GoString(C.XmNtopItemPosition)
var XmNtopLevelEnterCallback = C.GoString(C.XmNtopLevelEnterCallback)
var XmNtopLevelLeaveCallback = C.GoString(C.XmNtopLevelLeaveCallback)
var XmNtopOffset = C.GoString(C.XmNtopOffset)
var XmNtopPosition = C.GoString(C.XmNtopPosition)
var XmNtopShadowColor = C.GoString(C.XmNtopShadowColor)
var XmNtopShadowPixmap = C.GoString(C.XmNtopShadowPixmap)
var XmNtopWidget = C.GoString(C.XmNtopWidget)
var XmNtransferProc = C.GoString(C.XmNtransferProc)
var XmNtransferStatus = C.GoString(C.XmNtransferStatus)
var XmNtraversalCallback = C.GoString(C.XmNtraversalCallback)
var XmNtraversalOn = C.GoString(C.XmNtraversalOn)
var XmNtraversalType = C.GoString(C.XmNtraversalType)
var XmNtraverseObscuredCallback = C.GoString(C.XmNtraverseObscuredCallback)
var XmNtreeUpdateProc = C.GoString(C.XmNtreeUpdateProc)
var XmNtroughColor = C.GoString(C.XmNtroughColor)
var XmNunitType = C.GoString(C.XmNunitType)
var XmNunmapCallback = C.GoString(C.XmNunmapCallback)
var XmNunpostBehavior = C.GoString(C.XmNunpostBehavior)
var XmNunselectPixmap = C.GoString(C.XmNunselectPixmap)
var XmNupdateSliderSize = C.GoString(C.XmNupdateSliderSize)
var XmNuseAsyncGeometry = C.GoString(C.XmNuseAsyncGeometry)
var XmNuserData = C.GoString(C.XmNuserData)
var XmNvalidCursorForeground = C.GoString(C.XmNvalidCursorForeground)
var XmNvalueChangedCallback = C.GoString(C.XmNvalueChangedCallback)
var XmNvalueWcs = C.GoString(C.XmNvalueWcs)
var XmNverifyBell = C.GoString(C.XmNverifyBell)
var XmNverticalFontUnit = C.GoString(C.XmNverticalFontUnit)
var XmNverticalScrollBar = C.GoString(C.XmNverticalScrollBar)
var XmNverticalSpacing = C.GoString(C.XmNverticalSpacing)
var XmNvisibleItemCount = C.GoString(C.XmNvisibleItemCount)
var XmNvisibleWhenOff = C.GoString(C.XmNvisibleWhenOff)
var XmNvisualPolicy = C.GoString(C.XmNvisualPolicy)
var XmNwhichButton = C.GoString(C.XmNwhichButton)
var XmNwordWrap = C.GoString(C.XmNwordWrap)
var XmNworkWindow = C.GoString(C.XmNworkWindow)
var XmNtearOffTitle = C.GoString(C.XmNtearOffTitle)
var XmNpopupHandlerCallback = C.GoString(C.XmNpopupHandlerCallback)
var XmNconvertCallback = C.GoString(C.XmNconvertCallback)
var XmNdestinationCallback = C.GoString(C.XmNdestinationCallback)
var XmNselectedItem = C.GoString(C.XmNselectedItem)
var XmNselectionCallback = C.GoString(C.XmNselectionCallback)
var XmNmatchBehavior = C.GoString(C.XmNmatchBehavior)
var XmNnoFontCallback = C.GoString(C.XmNnoFontCallback)
var XmNtextPath = C.GoString(C.XmNtextPath)
var XmNeditingPath = C.GoString(C.XmNeditingPath)
var XmNbidirectionalCursor = C.GoString(C.XmNbidirectionalCursor)
var XmNcollapsedStatePixmap = C.GoString(C.XmNcollapsedStatePixmap)
var XmNdetailColumnHeading = C.GoString(C.XmNdetailColumnHeading)
var XmNdetailCount = C.GoString(C.XmNdetailCount)
var XmNdetailTabList = C.GoString(C.XmNdetailTabList)
var XmNexpandedStatePixmap = C.GoString(C.XmNexpandedStatePixmap)
var XmNlargeCellHeight = C.GoString(C.XmNlargeCellHeight)
var XmNlargeCellWidth = C.GoString(C.XmNlargeCellWidth)
var XmNlayoutType = C.GoString(C.XmNlayoutType)
var XmNoutlineIndentation = C.GoString(C.XmNoutlineIndentation)
var XmNoutlineLineStyle = C.GoString(C.XmNoutlineLineStyle)
var XmNprimaryOwnership = C.GoString(C.XmNprimaryOwnership)
var XmNselectionTechnique = C.GoString(C.XmNselectionTechnique)
var XmNsmallCellHeight = C.GoString(C.XmNsmallCellHeight)
var XmNsmallCellWidth = C.GoString(C.XmNsmallCellWidth)
var XmNspatialStyle = C.GoString(C.XmNspatialStyle)
var XmNentryParent = C.GoString(C.XmNentryParent)
var XmNlargeIconX = C.GoString(C.XmNlargeIconX)
var XmNlargeIconY = C.GoString(C.XmNlargeIconY)
var XmNsmallIconX = C.GoString(C.XmNsmallIconX)
var XmNsmallIconY = C.GoString(C.XmNsmallIconY)
var XmNlargeIcon = C.GoString(C.XmNlargeIcon)
var XmNlargeIconMask = C.GoString(C.XmNlargeIconMask)
var XmNlargeIconPixmap = C.GoString(C.XmNlargeIconPixmap)
var XmNsmallIcon = C.GoString(C.XmNsmallIcon)
var XmNsmallIconMask = C.GoString(C.XmNsmallIconMask)
var XmNsmallIconPixmap = C.GoString(C.XmNsmallIconPixmap)
var XmNcurrentPageNumber = C.GoString(C.XmNcurrentPageNumber)
var XmNfirstPageNumber = C.GoString(C.XmNfirstPageNumber)
var XmNlastPageNumber = C.GoString(C.XmNlastPageNumber)
var XmNbackPagePlacement = C.GoString(C.XmNbackPagePlacement)
var XmNbackPageNumber = C.GoString(C.XmNbackPageNumber)
var XmNbackPageSize = C.GoString(C.XmNbackPageSize)
var XmNbackPageForeground = C.GoString(C.XmNbackPageForeground)
var XmNbackPageBackground = C.GoString(C.XmNbackPageBackground)
var XmNframeBackground = C.GoString(C.XmNframeBackground)
var XmNbindingType = C.GoString(C.XmNbindingType)
var XmNbindingPixmap = C.GoString(C.XmNbindingPixmap)
var XmNbindingWidth = C.GoString(C.XmNbindingWidth)
var XmNmajorTabSpacing = C.GoString(C.XmNmajorTabSpacing)
var XmNminorTabSpacing = C.GoString(C.XmNminorTabSpacing)
var XmNinnerMarginWidth = C.GoString(C.XmNinnerMarginWidth)
var XmNinnerMarginHeight = C.GoString(C.XmNinnerMarginHeight)
var XmNframeShadowThickness = C.GoString(C.XmNframeShadowThickness)
var XmNpageNumber = C.GoString(C.XmNpageNumber)
var XmNarrowLayout = C.GoString(C.XmNarrowLayout)
var XmNarrowSensitivity = C.GoString(C.XmNarrowSensitivity)
var XmNdefaultArrowSensitivity = C.GoString(C.XmNdefaultArrowSensitivity)
var XmNarrowSize = C.GoString(C.XmNarrowSize)
var XmNspinBoxChildType = C.GoString(C.XmNspinBoxChildType)
var XmNposition = C.GoString(C.XmNposition)
var XmNnumValues = C.GoString(C.XmNnumValues)
var XmNvalues = C.GoString(C.XmNvalues)
var XmNminimumValue = C.GoString(C.XmNminimumValue)
var XmNmaximumValue = C.GoString(C.XmNmaximumValue)
var XmNincrementValue = C.GoString(C.XmNincrementValue)
var XmNentryViewType = C.GoString(C.XmNentryViewType)
var XmNinsensitiveStippleBitmap = C.GoString(C.XmNinsensitiveStippleBitmap)
var XmNlayoutDirection = C.GoString(C.XmNlayoutDirection)
var XmNviewType = C.GoString(C.XmNviewType)
var XmNvisualEmphasis = C.GoString(C.XmNvisualEmphasis)
var XmNsnapBackMultiple = C.GoString(C.XmNsnapBackMultiple)
var XmNslidingMode = C.GoString(C.XmNslidingMode)
var XmNsliderVisual = C.GoString(C.XmNsliderVisual)
var XmNautoDragModel = C.GoString(C.XmNautoDragModel)
var XmNcolorCalculationProc = C.GoString(C.XmNcolorCalculationProc)
var XmNbitmapConversionModel = C.GoString(C.XmNbitmapConversionModel)
var XmNcolorAllocationProc = C.GoString(C.XmNcolorAllocationProc)
var XmNselectionMode = C.GoString(C.XmNselectionMode)
var XmNselectedPositions = C.GoString(C.XmNselectedPositions)
var XmNselectedPositionCount = C.GoString(C.XmNselectedPositionCount)
var XmNinputPolicy = C.GoString(C.XmNinputPolicy)
var XmNtoggleMode = C.GoString(C.XmNtoggleMode)
var XmNindeterminatePixmap = C.GoString(C.XmNindeterminatePixmap)
var XmNunselectColor = C.GoString(C.XmNunselectColor)
var XmNselectedPosition = C.GoString(C.XmNselectedPosition)
var XmNarrowSpacing = C.GoString(C.XmNarrowSpacing)
var XmNenableWarp = C.GoString(C.XmNenableWarp)
var XmNmotifVersion = C.GoString(C.XmNmotifVersion)
var XmNdefaultGlyphPixmap = C.GoString(C.XmNdefaultGlyphPixmap)
var XmNtag = C.GoString(C.XmNtag)
var XmNfontName = C.GoString(C.XmNfontName)
var XmNfontType = C.GoString(C.XmNfontType)
var XmNloadModel = C.GoString(C.XmNloadModel)
var XmNtabList = C.GoString(C.XmNtabList)
var XmNunderlineType = C.GoString(C.XmNunderlineType)
var XmNstrikethruType = C.GoString(C.XmNstrikethruType)
var XmNrenderTable = C.GoString(C.XmNrenderTable)
var XmNbuttonRenderTable = C.GoString(C.XmNbuttonRenderTable)
var XmNlabelRenderTable = C.GoString(C.XmNlabelRenderTable)
var XmNtextRenderTable = C.GoString(C.XmNtextRenderTable)
var XmNdragStartCallback = C.GoString(C.XmNdragStartCallback)
var XmNnoRenditionCallback = C.GoString(C.XmNnoRenditionCallback)
var XmNpatternType = C.GoString(C.XmNpatternType)
var XmNsubstitute = C.GoString(C.XmNsubstitute)
var XmNinvokeParseProc = C.GoString(C.XmNinvokeParseProc)
var XmNincludeStatus = C.GoString(C.XmNincludeStatus)
var XmNnotebookChildType = C.GoString(C.XmNnotebookChildType)
var XmNscrolledWindowChildType = C.GoString(C.XmNscrolledWindowChildType)
var XmNselectedObjects = C.GoString(C.XmNselectedObjects)
var XmNselectedObjectCount = C.GoString(C.XmNselectedObjectCount)
var XmNcomboBoxType = C.GoString(C.XmNcomboBoxType)
var XmNtabValue = C.GoString(C.XmNtabValue)
var XmNoffsetModel = C.GoString(C.XmNoffsetModel)
var XmNdecimal = C.GoString(C.XmNdecimal)
var XmNdetail = C.GoString(C.XmNdetail)
var XmNcontainerID = C.GoString(C.XmNcontainerID)
var XmNpathMode = C.GoString(C.XmNpathMode)
var XmNfileFilterStyle = C.GoString(C.XmNfileFilterStyle)
var XmNdirTextLabelString = C.GoString(C.XmNdirTextLabelString)
var XmNenableBtn1Transfer = C.GoString(C.XmNenableBtn1Transfer)
var XmNenableButtonTab = C.GoString(C.XmNenableButtonTab)
var XmNenableEtchedInMenu = C.GoString(C.XmNenableEtchedInMenu)
var XmNdefaultButtonEmphasis = C.GoString(C.XmNdefaultButtonEmphasis)
var XmNenableToggleColor = C.GoString(C.XmNenableToggleColor)
var XmNenableToggleVisual = C.GoString(C.XmNenableToggleVisual)
var XmNenableDragIcon = C.GoString(C.XmNenableDragIcon)
var XmNenableUnselectableDrag = C.GoString(C.XmNenableUnselectableDrag)
var XmNdragOverActiveMode = C.GoString(C.XmNdragOverActiveMode)
var XmNinstallColormap = C.GoString(C.XmNinstallColormap)
var XmNownerEvents = C.GoString(C.XmNownerEvents)
var XmNgrabStyle = C.GoString(C.XmNgrabStyle)
var XmNforegroundState = C.GoString(C.XmNforegroundState)
var XmNbackgroundState = C.GoString(C.XmNbackgroundState)
var XmNoutlineState = C.GoString(C.XmNoutlineState)
var XmNspatialIncludeModel = C.GoString(C.XmNspatialIncludeModel)
var XmNspatialResizeModel = C.GoString(C.XmNspatialResizeModel)
var XmNspatialSnapModel = C.GoString(C.XmNspatialSnapModel)
var XmNdetailColumnHeadingCount = C.GoString(C.XmNdetailColumnHeadingCount)
var XmNdetailOrder = C.GoString(C.XmNdetailOrder)
var XmNdetailOrderCount = C.GoString(C.XmNdetailOrderCount)
var XmNoutlineColumnWidth = C.GoString(C.XmNoutlineColumnWidth)
var XmNoutlineChangedCallback = C.GoString(C.XmNoutlineChangedCallback)
var XmNoutlineButtonPolicy = C.GoString(C.XmNoutlineButtonPolicy)
var XmNdefaultVirtualBindings = C.GoString(C.XmNdefaultVirtualBindings)
var XmNpageChangedCallback = C.GoString(C.XmNpageChangedCallback)
var XmNarea = C.GoString(C.XmNarea)
var XmNdetailShadowThickness = C.GoString(C.XmNdetailShadowThickness)
var XmNsliderMark = C.GoString(C.XmNsliderMark)
var XmNrenditionBackground = C.GoString(C.XmNrenditionBackground)
var XmNrenditionForeground = C.GoString(C.XmNrenditionForeground)
var XmNindeterminateInsensitivePixmap = C.GoString(C.XmNindeterminateInsensitivePixmap)
var XmNframeChildType = C.GoString(C.XmNframeChildType)
var XmNtextField = C.GoString(C.XmNtextField)
var XmNenableThinThickness = C.GoString(C.XmNenableThinThickness)
var XmNprimaryColorSetId = C.GoString(C.XmNprimaryColorSetId)
var XmNsecondaryColorSetId = C.GoString(C.XmNsecondaryColorSetId)
var XmNtextColorSetId = C.GoString(C.XmNtextColorSetId)
var XmNactiveColorSetId = C.GoString(C.XmNactiveColorSetId)
var XmNinactiveColorSetId = C.GoString(C.XmNinactiveColorSetId)
var XmNuseColorObj = C.GoString(C.XmNuseColorObj)
var XmNuseTextColor = C.GoString(C.XmNuseTextColor)
var XmNuseTextColorForList = C.GoString(C.XmNuseTextColorForList)
var XmNuseMask = C.GoString(C.XmNuseMask)
var XmNuseMultiColorIcons = C.GoString(C.XmNuseMultiColorIcons)
var XmNuseIconFileCache = C.GoString(C.XmNuseIconFileCache)
var XmNlist = C.GoString(C.XmNlist)
var XmNarrowOrientation = C.GoString(C.XmNarrowOrientation)
var XmNpositionType = C.GoString(C.XmNpositionType)
var XmNwrap = C.GoString(C.XmNwrap)
var XmNpositionMode = C.GoString(C.XmNpositionMode)
var XmNprintOrientation = C.GoString(C.XmNprintOrientation)
var XmNprintOrientations = C.GoString(C.XmNprintOrientations)
var XmNprintResolution = C.GoString(C.XmNprintResolution)
var XmNprintResolutions = C.GoString(C.XmNprintResolutions)
var XmNdefaultPixmapResolution = C.GoString(C.XmNdefaultPixmapResolution)
var XmNstartJobCallback = C.GoString(C.XmNstartJobCallback)
var XmNendJobCallback = C.GoString(C.XmNendJobCallback)
var XmNpageSetupCallback = C.GoString(C.XmNpageSetupCallback)
var XmNpdmNotificationCallback = C.GoString(C.XmNpdmNotificationCallback)
var XmNminX = C.GoString(C.XmNminX)
var XmNminY = C.GoString(C.XmNminY)
var XmNmaxX = C.GoString(C.XmNmaxX)
var XmNmaxY = C.GoString(C.XmNmaxY)
var XmNpreeditStartCallback = C.GoString(C.XmNpreeditStartCallback)
var XmNpreeditDoneCallback = C.GoString(C.XmNpreeditDoneCallback)
var XmNpreeditDrawCallback = C.GoString(C.XmNpreeditDrawCallback)
var XmNpreeditCaretCallback = C.GoString(C.XmNpreeditCaretCallback)
var XmNverifyPreedit = C.GoString(C.XmNverifyPreedit)
var XmNenableMultiKeyBindings = C.GoString(C.XmNenableMultiKeyBindings)

// these are strings and should be defined as literals according to Xm/XmStrDefs.h
// var XmNtoolTipString = C.GoString(C.XmNtoolTipString)
// var XmNtoolTipPostDelay = C.GoString(C.XmNtoolTipPostDelay)
// var XmNtoolTipPostDuration = C.GoString(C.XmNtoolTipPostDuration)
// var XmNtoolTipEnable = C.GoString(C.XmNtoolTipEnable)

var XmNaccelerators = C.GoString(C.XtNaccelerators)
var XmNallowShellResize = C.GoString(C.XtNallowShellResize)
var XmNancestorSensitive = C.GoString(C.XtNancestorSensitive)
var XmNargc = C.GoString(C.XtNargc)
var XmNargv = C.GoString(C.XtNargv)
var XmNbackground = C.GoString(C.XtNbackground)
var XmNbackgroundPixmap = C.GoString(C.XtNbackgroundPixmap)
var XmNbaseHeight = C.GoString(C.XtNbaseHeight)
var XmNbaseWidth = C.GoString(C.XtNbaseWidth)
var XmNbitmap = C.GoString(C.XtNbitmap)
var XmNborder = C.GoString(C.XtNborder)
var XmNborderColor = C.GoString(C.XtNborderColor)
var XmNborderPixmap = C.GoString(C.XtNborderPixmap)
var XmNborderWidth = C.GoString(C.XtNborderWidth)
var XmNcallback = C.GoString(C.XtNcallback)
var XmNchildren = C.GoString(C.XtNchildren)
var XmNcolormap = C.GoString(C.XtNcolormap)
var XmNcreatePopupChildProc = C.GoString(C.XtNcreatePopupChildProc)
var XmNdepth = C.GoString(C.XtNdepth)
var XmNdestroyCallback = C.GoString(C.XtNdestroyCallback)
var XmNeditType = C.GoString(C.XtNeditType)
var XmNfile = C.GoString(C.XtNfile)
var XmNfont = C.GoString(C.XtNfont)
var XmNfontSet = C.GoString(C.XtNfontSet)
var XmNforceBars = C.GoString(C.XtNforceBars)
var XmNforeground = C.GoString(C.XtNforeground)
var XmNfunction = C.GoString(C.XtNfunction)
var XmNgeometry = C.GoString(C.XtNgeometry)
var XmNheight = C.GoString(C.XtNheight)
var XmNheightInc = C.GoString(C.XtNheightInc)
var XmNhighlight = C.GoString(C.XtNhighlight)
var XmNiconMask = C.GoString(C.XtNiconMask)
var XmNiconName = C.GoString(C.XtNiconName)
var XmNiconNameEncoding = C.GoString(C.XtNiconNameEncoding)
var XmNiconPixmap = C.GoString(C.XtNiconPixmap)
var XmNiconWindow = C.GoString(C.XtNiconWindow)
var XmNiconX = C.GoString(C.XtNiconX)
var XmNiconY = C.GoString(C.XtNiconY)
var XmNiconic = C.GoString(C.XtNiconic)
var XmNindex = C.GoString(C.XtNindex)
var XmNinitialResourcesPersistent = C.GoString(C.XtNinitialResourcesPersistent)
var XmNinitialState = C.GoString(C.XtNinitialState)
var XmNinnerHeight = C.GoString(C.XtNinnerHeight)
var XmNinnerWidth = C.GoString(C.XtNinnerWidth)
var XmNinnerWindow = C.GoString(C.XtNinnerWindow)
var XmNinput = C.GoString(C.XtNinput)
var XmNinsertPosition = C.GoString(C.XtNinsertPosition)
var XmNinternalHeight = C.GoString(C.XtNinternalHeight)
var XmNinternalWidth = C.GoString(C.XtNinternalWidth)
var XmNjumpProc = C.GoString(C.XtNjumpProc)
var XmNjustify = C.GoString(C.XtNjustify)
var XmNlength = C.GoString(C.XtNlength)
var XmNlowerRight = C.GoString(C.XtNlowerRight)
var XmNmappedWhenManaged = C.GoString(C.XtNmappedWhenManaged)
var XmNmaxAspectX = C.GoString(C.XtNmaxAspectX)
var XmNmaxAspectY = C.GoString(C.XtNmaxAspectY)
var XmNmaxHeight = C.GoString(C.XtNmaxHeight)
var XmNmaxWidth = C.GoString(C.XtNmaxWidth)
var XmNmenuEntry = C.GoString(C.XtNmenuEntry)
var XmNminAspectX = C.GoString(C.XtNminAspectX)
var XmNminAspectY = C.GoString(C.XtNminAspectY)
var XmNminHeight = C.GoString(C.XtNminHeight)
var XmNminWidth = C.GoString(C.XtNminWidth)
var XmNname = C.GoString(C.XtNname)
var XmNnotify = C.GoString(C.XtNnotify)
var XmNnumChildren = C.GoString(C.XtNnumChildren)
var XmNorientation = C.GoString(C.XtNorientation)
var XmNoverrideRedirect = C.GoString(C.XtNoverrideRedirect)
var XmNparameter = C.GoString(C.XtNparameter)
var XmNpixmap = C.GoString(C.XtNpixmap)
var XmNpopdownCallback = C.GoString(C.XtNpopdownCallback)
var XmNpopupCallback = C.GoString(C.XtNpopupCallback)
var XmNresize = C.GoString(C.XtNresize)
var XmNreverseVideo = C.GoString(C.XtNreverseVideo)
var XmNsaveUnder = C.GoString(C.XtNsaveUnder)
var XmNscreen = C.GoString(C.XtNscreen)
var XmNscrollDCursor = C.GoString(C.XtNscrollDCursor)
var XmNscrollHCursor = C.GoString(C.XtNscrollHCursor)
var XmNscrollLCursor = C.GoString(C.XtNscrollLCursor)
var XmNscrollProc = C.GoString(C.XtNscrollProc)
var XmNscrollRCursor = C.GoString(C.XtNscrollRCursor)
var XmNscrollUCursor = C.GoString(C.XtNscrollUCursor)
var XmNscrollVCursor = C.GoString(C.XtNscrollVCursor)
var XmNselection = C.GoString(C.XtNselection)
var XmNselectionArray = C.GoString(C.XtNselectionArray)
var XmNsensitive = C.GoString(C.XtNsensitive)
var XmNshown = C.GoString(C.XtNshown)
var XmNspace = C.GoString(C.XtNspace)
var XmNstring = C.GoString(C.XtNstring)
var XmNtextOptions = C.GoString(C.XtNtextOptions)
var XmNtextSink = C.GoString(C.XtNtextSink)
var XmNtextSource = C.GoString(C.XtNtextSource)
var XmNthickness = C.GoString(C.XtNthickness)
var XmNthumb = C.GoString(C.XtNthumb)
var XmNthumbProc = C.GoString(C.XtNthumbProc)
var XmNtitle = C.GoString(C.XtNtitle)
var XmNtitleEncoding = C.GoString(C.XtNtitleEncoding)
var XmNtop = C.GoString(C.XtNtop)
var XmNtransient = C.GoString(C.XtNtransient)
var XmNtransientFor = C.GoString(C.XtNtransientFor)
var XmNtranslations = C.GoString(C.XtNtranslations)
var XmNupdate = C.GoString(C.XtNupdate)
var XmNuseBottom = C.GoString(C.XtNuseBottom)
var XmNuseRight = C.GoString(C.XtNuseRight)
var XmNvalue = C.GoString(C.XtNvalue)
var XmNvisual = C.GoString(C.XtNvisual)
var XmNwaitForWm = C.GoString(C.XtNwaitForWm)
var XmNwidth = C.GoString(C.XtNwidth)
var XmNwidthInc = C.GoString(C.XtNwidthInc)
var XmNwinGravity = C.GoString(C.XtNwinGravity)
var XmNwindow = C.GoString(C.XtNwindow)
var XmNwindowGroup = C.GoString(C.XtNwindowGroup)
var XmNwmTimeout = C.GoString(C.XtNwmTimeout)
var XmNx = C.GoString(C.XtNx)
var XmNy = C.GoString(C.XtNy)

//var XmNanimate = C.GoString(C."animate")

var XmFONTLIST_DEFAULT_TAG = XmStringCharset((*C.char)(unsafe.Pointer(C.XmFONTLIST_DEFAULT_TAG)))
var XmSTRING_DEFAULT_CHARSET = XmStringCharset((*C.char)(unsafe.Pointer(C.XmSTRING_DEFAULT_CHARSET)))
