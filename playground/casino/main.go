package main

import (
	"context"
	"fmt"
	"github.com/AmpyFin/yfinance-go"
	"go-x11/pkg/x11"
	"log"
	"math"
	"os"
	"time"
)

/*
"time"
"github.com/AmpyFin/yfinance-go"
*/

var canvas x11.Widget

var client = yfinance.NewClient()

func finance() {
	// Erstellt einen neuen Yahoo Finance Client

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
	for _, bar := range bars.Bars {
		//test := math.Abs(1.0)
		price := float64(bar.Close.Scaled) / math.Pow(10, float64(bar.Close.Scale))
		fmt.Printf("Date: %s, Close: %.4f %s\n",
			bar.EventTime.Format("2006-01-02"),
			price, bar.CurrencyCode)
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

func drawIt(w x11.Widget, clientData x11.XtPointer, callData x11.XtPointer) {
	println("Some drawing primitives...")

	// FillRectangle
	cw := x11.XtWindow(canvas)
	drawable := x11.Drawable(cw)
	d := x11.XtDisplay(canvas)
	gc := x11.XCreateGC(d, drawable, 0, nil)
	//x11.XSetForeground(d, gc, x11.BlackPixelOfScreen(x11.XtScreen(canvas)))
	x11.XSetForeground(d, gc, 0x0000ff)
	//x11.XDrawRectangle(d, drawable, gc, 10, 10, 90, 70)
	x11.XFillRectangle(d, drawable, gc, 10, 10, 90, 70)

	// Lines
	x11.XSetForeground(d, gc, 0x00ff00)
	x11.XDrawLine(d, drawable, gc, 20, 25, 70, 20)
	x11.XSetForeground(d, gc, 0x00ffff)
	x11.XDrawLine(d, drawable, gc, 30, 30, 70, 30)

	// Arc/Circle
	x11.XSetForeground(d, gc, 0xff0000)
	x11.XFillArc(d, drawable, gc, 30, 30, 50, 50, 0, 360*64)

	// String
	x11.XSetForeground(d, gc, 0xffffff)
	x11.XDrawString(d, drawable, gc, 40, 20, "HELLO")

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
	panel := x11.XtCreateManagedWidget("message", x11.FormWidgetClass(), shell, &x11.ArgList{})

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
	canvas = x11.XtCreateManagedWidget("canvas", x11.DrawingAreaWidgetClass(), panel, args)

	bu1 := x11.XtCreateManagedWidget("button1", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("button2", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("button3", x11.PushButtonWidgetClass(), commands, nil)

	x11.XtCreateManagedWidget("button1", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("button2", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("button3", x11.PushButtonWidgetClass(), options, nil)

	x11.XtAddCallback(bu1, x11.XmNactivateCallback, drawIt, x11.XtPointer(nil))

	finance()

	x11.XtRealizeWidget(shell)

	x11.XtAppMainLoop(app)

}
