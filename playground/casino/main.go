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

	p := quote.RegularMarketPrice
	//var q norm.ScaledDecimal = p
	fmt.Printf("Quote: %s\n", quote)
	price := float64(p.Scaled) / math.Pow(10, float64(p.Scale))
	fmt.Printf("Quote: Regular %v %s\n", price, currency)

	p = quote.RegularMarketHigh
	price = float64(p.Scaled) / math.Pow(10, float64(p.Scale))
	fmt.Printf("Quote: Hight   %v %s\n", price, currency)

	p = quote.RegularMarketLow
	price = float64(p.Scaled) / math.Pow(10, float64(p.Scale))
	fmt.Printf("Quote: Low     %v %s\n", price, currency)
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
	canvas := x11.XtCreateManagedWidget("canvas", x11.DrawingAreaWidgetClass(), panel, args)
	_ = canvas

	x11.XtCreateManagedWidget("button1", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("button2", x11.PushButtonWidgetClass(), commands, nil)
	x11.XtCreateManagedWidget("button3", x11.PushButtonWidgetClass(), commands, nil)

	x11.XtCreateManagedWidget("button1", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("button2", x11.PushButtonWidgetClass(), options, nil)
	x11.XtCreateManagedWidget("button3", x11.PushButtonWidgetClass(), options, nil)

	finance()

	x11.XtRealizeWidget(shell)
	x11.XtAppMainLoop(app)

}
