package main

import (
	"context"
	"fmt"
	"github.com/AmpyFin/yfinance-go"
	"go-x11/pkg/x11"
	"log"
	"math"
	"os"
	"strconv"
	"time"
	"unsafe"
)

/*
"time"
"github.com/AmpyFin/yfinance-go"
*/

var canvas x11.Widget

var client = yfinance.NewClient()

type QuoteData struct {
	Date         time.Time
	Close        float64
	CurrencyCode string
}

var v_lower_y = 200
var v_upper_y = 500
var coordSystem = CoordSystemInfo{
	swidth:        2000,
	sheight:       1000,
	vx:            0,
	vy:            0,
	vwidth:        2000,
	v_lower_y:     v_lower_y,
	v_upper_y:     v_upper_y,
	width:         2000,
	height:        1000,
	margin_l:      20,
	margin_r:      20,
	margin_top:    20,
	margin_bottom: 40,
	num_ticks_x:   20,
	num_ticks_y:   (v_upper_y - v_lower_y) / 100,
	len_tick:      10,
	legend_x:      nil,
	legend_y:      nil,
}

type CoordSystemInfo struct {
	swidth        int // source coordinate system width
	sheight       int // source coordinate system height
	vx            int // view coordinate system offset x
	vy            int // view coordinate system offset y
	vwidth        int // view coordinate system width
	v_upper_y     int // view coordinate system upper y value to use (defines window inside v)
	v_lower_y     int // view coordinate system lower y value to use (defines window inside v)
	width         int // target coordinate system width (pixel coordinates)
	height        int // target coordinate system width (pixel coordinates)
	margin_l      int
	margin_r      int
	margin_top    int
	margin_bottom int
	num_ticks_x   int
	num_ticks_y   int
	len_tick      int
	legend_x      []string
	legend_y      []string
}

var quoteData []QuoteData

func finance() {
	// Create a new client
	ctx := context.Background()

	//who := "ALV.DE"
	who := "AAPL"
	// Fetch daily bars
	start := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)

	//end := time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)
	end := time.Now()

	runId := "my-run-id"
	bars, err := client.FetchDailyBars(ctx, who, start, end, true, runId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fetched %d bars for %v\n", len(bars.Bars), who)
	quoteData = nil
	for _, bar := range bars.Bars {
		//test := math.Abs(1.0)
		price := float64(bar.Close.Scaled) / math.Pow(10, float64(bar.Close.Scale))
		fmt.Printf("Date: %s, Close: %.4f %s\n",
			bar.EventTime.Format("2006-01-02"),
			price, bar.CurrencyCode)
		newQuote := QuoteData{
			Date:         bar.EventTime,
			Close:        price,
			CurrencyCode: bar.CurrencyCode,
		}
		quoteData = append(quoteData, newQuote)
	}

	quote, err := client.FetchQuote(ctx, who, runId)
	if err != nil {
		log.Fatal(err)
	}
	currency := quote.CurrencyCode
	q := *quote.RegularMarketVolume
	fmt.Printf("Quote: Volume: %v\n", q)
	fmt.Printf("Quote: %v\n", quote)

	p := quote.RegularMarketPrice
	//var q norm.ScaledDecimal = p
	//var p ScaledDecimal
	var price float64
	if p != nil {
		price = float64(p.Scaled) / math.Pow(10, float64(p.Scale))
		fmt.Printf("Quote: Regular %v %s\n", price, currency)
	}
	if p != nil {
		p = quote.RegularMarketHigh
		price = float64(p.Scaled) / math.Pow(10, float64(p.Scale))
		fmt.Printf("Quote: Hight   %v %s\n", price, currency)
	}
	if p != nil {
		p = quote.RegularMarketLow
		price = float64(p.Scaled) / math.Pow(10, float64(p.Scale))
		fmt.Printf("Quote: Low     %v %s\n", price, currency)
	}
	if p == nil {
		fmt.Printf("No RegularMarketPrice values found\n")
	}
}

func transform(coo *CoordSystemInfo, x int, y int) (int, int) {
	vheight := coo.v_upper_y - coo.v_lower_y
	tx := (x - coo.vx) * coo.width / coo.vwidth
	ty := (vheight - (y - coo.v_lower_y) - coo.vy) * coo.height / vheight
	return tx, ty
}

func redisplay(w x11.Widget, clientData x11.XtPointer, callData x11.XtPointer) {
	println("Some drawing primitives...")

	var height uint16
	var width uint16
	// get pointer to var
	ptrToHeight := unsafe.Pointer(&height)
	ptrToWidth := unsafe.Pointer(&width)
	//convert to uintptr to put it into ArgList
	heightAsUintptr := uintptr(ptrToHeight)
	widthAsUintptr := uintptr(ptrToWidth)
	inargs := x11.AppendArgList(nil, x11.XmNheight, heightAsUintptr)
	inargs = x11.AppendArgList(inargs, x11.XmNwidth, widthAsUintptr)
	x11.XtGetValues(canvas, inargs)

	fmt.Printf("canvas width*height: %d x %d\n", width, height)
	// update coordsystem struct
	coordSystem.width = int(width)
	coordSystem.height = int(height)

	// FillRectangle
	cw := x11.XtWindow(canvas)
	drawable := x11.Drawable(cw)
	d := x11.XtDisplay(canvas)
	gc := x11.XCreateGC(d, drawable, 0, nil)
	//x11.XSetForeground(d, gc, x11.BlackPixelOfScreen(x11.XtScreen(canvas)))
	x11.XSetForeground(d, gc, 0x0000ff)
	//x11.XDrawRectangle(d, drawable, gc, 10, 10, 90, 70)
	x11.XFillRectangle(d, drawable, gc, 0, 0, uint(width), uint(height))

	// Lines
	x11.XSetForeground(d, gc, 0x00ff00)
	x11.XDrawLine(d, drawable, gc, 530, 100, 700, 50)
	x11.XSetForeground(d, gc, 0x00ffff)
	x11.XDrawLine(d, drawable, gc, 530, 100, 700, 300)

	// Arc/Circle
	x11.XSetForeground(d, gc, 0xff0000)
	x11.XFillArc(d, drawable, gc, 300, 500, 300, 300, 0, 360*64)

	// String
	x11.XSetForeground(d, gc, 0xffffff)
	x11.XDrawString(d, drawable, gc, 40, 200, "HELLO")

	// Pixels
	x11.XSetForeground(d, gc, 0xffff00)
	x11.XDrawPoint(d, drawable, gc, 50, 50)
	x11.XDrawPoint(d, drawable, gc, 52, 50)
	x11.XDrawPoint(d, drawable, gc, 54, 50)
	x11.XDrawPoint(d, drawable, gc, 56, 50)
	x11.XDrawPoint(d, drawable, gc, 58, 50)
	x11.XDrawPoint(d, drawable, gc, 60, 50)

	// Array of pixels, CoordModeOrigin - coords based on widget 0,0
	points := []x11.XPoint{
		{80, 10},
		{81, 11},
		{82, 12},
		{83, 13},
	}
	x11.XDrawPoints(d, drawable, gc, points, x11.CoordModeOrigin)

	// Array of pixels, CoordModeOrigin - coords based on previous point
	x11.XSetForeground(d, gc, 0x00ffff)
	points = []x11.XPoint{
		{90, 20},
		{1, 2},
		{1, 2},
		{1, 2},
	}
	x11.XDrawPoints(d, drawable, gc, points, x11.CoordModePrevious)

	x11.XSetForeground(d, gc, 0xffff00)
	arcs := []x11.XArc{
		{500, 40, 20, 20, 0, 360 * 64},
		{500, 70, 20, 20, 0, 360 * 64},
		{500, 100, 20, 20, 0, 360 * 64},
	}
	x11.XDrawArcs(d, drawable, gc, arcs)

	coordSystem.legend_x = make([]string, coordSystem.num_ticks_x+1)
	for i := 0; i < coordSystem.num_ticks_x+1; i++ {
		y := 2010 + i
		coordSystem.legend_x[i] = strconv.Itoa(y)
	}
	coordSystem.legend_y = make([]string, coordSystem.num_ticks_y+1)
	start_index := coordSystem.v_lower_y / 100
	for i := start_index; i < coordSystem.num_ticks_y+1+start_index; i++ {
		coordSystem.legend_y[i-start_index] = strconv.Itoa(i)
	}

	drawCoordSystem(d, drawable, gc, &coordSystem)

	// some quote data points
	divisionLength := (int(width) - (coordSystem.margin_r + coordSystem.margin_l)) / len(quoteData)
	x11.XSetForeground(d, gc, 0xe7e78d)
	var arcsize uint = 20
	for i, l := range quoteData {
		fmt.Printf("Date: %s, Close: %.4f %s\n", l.Date, l.Close, l.CurrencyCode)
		qx := i*divisionLength + coordSystem.margin_l - int(arcsize)/2
		qy := int(l.Close)
		tx, ty := transform(&coordSystem, qx, qy)
		x11.XFillArc(d, drawable, gc, tx, ty, arcsize, arcsize, 0, 360*64)
	}
}

// drawCoordSystem Draw a XY coordinate system
func drawCoordSystem(d *x11.Display, drawable x11.Drawable, gc x11.GC, coo *CoordSystemInfo) {
	drawXAxis(d, drawable, gc, coo)
	drawYAxis(d, drawable, gc, coo)
}

// drawXAxis draws X axis of a coordinate system
func drawXAxis(d *x11.Display, drawable x11.Drawable, gc x11.GC, coo *CoordSystemInfo) {

	gContextId := x11.XGContextFromGC(gc)
	font := x11.XQueryFont(d, *(*x11.XID)(unsafe.Pointer(&gContextId)))
	var textDimensions x11.XCharStruct // Alloziert den Speicher in Go
	var dir int
	var ascent int
	var descent int

	//TODO
	// check numeric values below (15) these offsets need to be calculated by font size

	tickLen := coo.len_tick / 2

	coo_x_start := coo.margin_l
	coo_x_stop := uint(coo.width) - uint(coo.margin_r)
	coo_y_start := coo.height - coo.margin_bottom
	coo_y_stop := uint(coo.height) - uint(coo.margin_bottom)
	x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
	// ticks
	divisionLength := (coo.width - (coo.margin_r + coo.margin_l)) / coo.num_ticks_x
	for i := 0; i <= coo.num_ticks_x; i++ {
		coo_x_start = coo.margin_l + i*divisionLength
		coo_x_stop = uint(coo.margin_l) + uint(i*divisionLength)
		// len of tick = 10 (5+5)
		coo_y_start = coo.height - coo.margin_bottom + tickLen
		coo_y_stop = uint(coo.height) - uint(coo.margin_bottom) - uint(tickLen)
		x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
		// legend
		ctext := coo.legend_x[i]
		x11.XTextExtents(font, ctext, len(ctext), &dir, &ascent, &descent, &textDimensions)
		txt_x_start := coo_x_start - int(textDimensions.Width/2)
		x11.XDrawString(d, drawable, gc, txt_x_start, coo_y_start+15, ctext)
	}
}

// drawXAxis draws Y axis of a coordinate system
func drawYAxis(d *x11.Display, drawable x11.Drawable, gc x11.GC, coo *CoordSystemInfo) {
	gContextId := x11.XGContextFromGC(gc)
	font := x11.XQueryFont(d, *(*x11.XID)(unsafe.Pointer(&gContextId)))
	var textDimensions x11.XCharStruct // Alloziert den Speicher in Go
	var dir int
	var ascent int
	var descent int

	//TODO
	// check numeric values below (5,15) these offsets need to be calculated by font size

	tickLen := coo.len_tick / 2

	coo_x_start := coo.margin_l
	coo_x_stop := uint(coo.margin_l)
	coo_y_start := coo.height - coo.margin_bottom
	coo_y_stop := uint(coo.margin_top)
	x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
	// ticks
	divisionLength := (coo.height - (coo.margin_top + coo.margin_bottom)) / coo.num_ticks_y
	for i := 0; i <= coo.num_ticks_y; i++ {
		coo_y_start = coo.height - coo.margin_bottom - i*divisionLength
		coo_y_stop = uint(coo.height) - uint(coo.margin_bottom) - uint(i*divisionLength)
		// len of tick = 10 (5+5)
		coo_x_start = int(coo.margin_l) + tickLen
		coo_x_stop = uint(coo.margin_l) - uint(tickLen)
		x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
		// legend
		ctext := coo.legend_y[i]
		x11.XTextExtents(font, ctext, len(ctext), &dir, &ascent, &descent, &textDimensions)
		txt_x_start := int(coo.margin_l) - int(textDimensions.Width) - 5
		x11.XDrawString(d, drawable, gc, txt_x_start, coo_y_start+15, ctext)
	}
}

func viewYSliderCallback(w x11.Widget, clientData x11.XtPointer, callData x11.XtPointer) {
	println("viewYSliderCallback")
	cb := (*x11.XmScaleCallbackStruct)(unsafe.Pointer(callData))
	fmt.Printf("value: %v\n", cb.Value)
}

type Data struct {
	x int
}

func main() {
	x11.WrapperInfo()

	var app x11.XtAppContext
	var xrmOptions []x11.XrmOptionDescRec
	fallbacks := []string{""}
	initArgs := []x11.Arg{}

	shell := x11.XtAppInitialize(
		&app,
		"Editor",
		xrmOptions,
		os.Args,
		fallbacks,
		initArgs,
	)

	// The overall window layout is handled by an XmForm widget
	panel := x11.XtCreateManagedWidget("panel", x11.FormWidgetClass(), shell, &x11.ArgList{})

	//  An XmRowColumn widget holds the buttons along the top of the window
	args := x11.AppendArgList(nil, x11.XmNnumColumns, 3)
	args = x11.AppendArgList(args, x11.XmNorientation, uintptr(x11.XmHORIZONTAL))
	args = x11.AppendArgList(args, x11.XmNtopAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNrightAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNleftAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNbottomAttachment, uintptr(x11.XmATTACH_NONE))
	commands := x11.XtCreateManagedWidget("commands", x11.RowColumnWidgetClass(), panel, args)

	// Another XmRowColumn widget contains a column of buttons along the left side of the window
	args = x11.AppendArgList(nil, x11.XmNnumColumns, 1)
	args = x11.AppendArgList(args, x11.XmNorientation, uintptr(x11.XmVERTICAL))
	args = x11.AppendArgList(args, x11.XmNtopAttachment, uintptr(x11.XmATTACH_WIDGET))
	args = x11.AppendArgList(args, x11.XmNtopWidget, uintptr(commands))
	args = x11.AppendArgList(args, x11.XmNrightAttachment, uintptr(x11.XmATTACH_NONE))
	args = x11.AppendArgList(args, x11.XmNleftAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNbottomAttachment, uintptr(x11.XmATTACH_FORM))
	options := x11.XtCreateManagedWidget("options", x11.RowColumnWidgetClass(), panel, args)

	// The middle window, in which the application can display text or graphics is an XmDrawingArea widget.
	args = x11.AppendArgList(nil, x11.XmNtopAttachment, uintptr(x11.XmATTACH_WIDGET))
	args = x11.AppendArgList(args, x11.XmNtopWidget, uintptr(commands))
	args = x11.AppendArgList(args, x11.XmNrightAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XmNleftWidget, uintptr(options))
	args = x11.AppendArgList(args, x11.XmNleftAttachment, uintptr(x11.XmATTACH_WIDGET))
	args = x11.AppendArgList(args, x11.XmNbottomAttachment, uintptr(x11.XmATTACH_FORM))
	args = x11.AppendArgList(args, x11.XtNwidth, 2000)
	args = x11.AppendArgList(args, x11.XtNheight, 1000)
	canvas = x11.XtCreateManagedWidget("canvas", x11.DrawingAreaWidgetClass(), panel, args)

	x11.XtCreateManagedWidget("command1", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("command2", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("command3", x11.PushButtonWidgetClass(), commands, nil)

	x11.XtCreateManagedWidget("option1", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("option2", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("option3", x11.PushButtonWidgetClass(), options, nil)

	var data Data
	x11.XtAddCallback(canvas, x11.XmNexposeCallback, redisplay, (x11.XtPointer)(unsafe.Pointer(&data)))

	args = x11.AppendArgList(nil, x11.XmNminimum, uintptr(coordSystem.v_lower_y))
	args = x11.AppendArgList(args, x11.XmNmaximum, uintptr(coordSystem.v_upper_y))
	args = x11.AppendArgList(args, x11.XmNshowValue, 1)
	args = x11.AppendArgList(args, x11.XmNorientation, uintptr(x11.XmHORIZONTAL))
	viewYSlider := x11.XtCreateManagedWidget("view_y", x11.ScaleWidgetClass(), commands, args)
	x11.XtAddCallback(viewYSlider, x11.XmNvalueChangedCallback, viewYSliderCallback, nil)
	x11.XtAddCallback(viewYSlider, x11.XmNdragCallback, viewYSliderCallback, nil)

	finance()

	x11.XtRealizeWidget(shell)

	x11.XtAppMainLoop(app)

}
