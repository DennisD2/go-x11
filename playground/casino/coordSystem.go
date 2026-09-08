package main

import (
	"go-x11/pkg/x11"
	"strconv"
	"unsafe"
)

type VArea struct {
	x       int // view coordinate system offset x
	y       int // view coordinate system offset y
	lower_x int // view coordinate system lower x value to use (defines window inside s)
	upper_x int // view coordinate system upper x value to use (defines window inside s)
	lower_y int // view coordinate system lower y value to use (defines window inside s)
	upper_y int // view coordinate system lower y value to use (defines window inside s)
}

type AreaDimension struct {
	width         int
	height        int
	divisionSizeX int
	divisionSizeY int
}

type CoordSystemInfo struct {
	source        AreaDimension // source coordinate system dimension
	view          VArea         // v is window inside s
	legendView    VArea         // legend is also a view into legend of s
	target        AreaDimension // target coordinate system dimension (pixels)
	margin_l      int
	margin_r      int
	margin_top    int
	margin_bottom int
	num_ticks_x   int
	num_ticks_y   int
	division_x    int // size of x division in source units
	division_y    int // size of y division in source units
	len_tick      int
	legend_x      []string
	legend_y      []string
}

// init a coordinate system
func (coo *CoordSystemInfo) init() {
	coo.legend_x = make([]string, coo.num_ticks_x+1)
	for i := 0; i <= coo.num_ticks_x; i++ {
		y := 2010 + i
		coo.legend_x[i] = strconv.Itoa(y)
	}
	coo.legend_y = make([]string, coo.num_ticks_y+1)
	for i := 0; i <= coo.num_ticks_y; i++ {
		coo.legend_y[i] = strconv.Itoa(i)
	}
}

// Transforms source coordinate x,y to target coordinates
func (coo *CoordSystemInfo) transform(x int, y int) (int, int) {
	vheight := coo.view.upper_y - coo.view.lower_y
	vwidth := coo.view.upper_x - coo.view.lower_x
	tx := (x - coo.view.x) * (coo.target.width - coo.margin_l - coo.margin_r) / vwidth
	ty := (vheight - (y - coo.view.lower_y) - coo.view.y) * (coo.target.height - coo.margin_l - coo.margin_r) / vheight
	return tx, ty
}

// drawCoordSystem Draw a XY coordinate system
func (coo *CoordSystemInfo) drawCoordSystem(d *x11.Display, drawable x11.Drawable, gc x11.GC) {
	coo.drawXAxis(d, drawable, gc)
	coo.drawYAxis(d, drawable, gc)
}

// drawXAxis draws X axis of a coordinate system
func (coo *CoordSystemInfo) drawXAxis(d *x11.Display, drawable x11.Drawable, gc x11.GC) {
	gContextId := x11.XGContextFromGC(gc)
	font := x11.XQueryFont(d, *(*x11.XID)(unsafe.Pointer(&gContextId)))
	tickLen := coo.len_tick / 2

	//TODO
	// check numeric values below (15) these offsets need to be calculated by font size

	coo_x_start := coo.margin_l
	coo_x_stop := uint(coo.target.width) - uint(coo.margin_r)
	coo_y_start := coo.target.height - coo.margin_bottom
	coo_y_stop := uint(coo.target.height) - uint(coo.margin_bottom)
	x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
	// ticks
	divisionSize := coordSystem.source.divisionSizeX * (coo.source.width - coo.margin_l - coo.margin_r) / (coo.view.upper_x - coo.view.lower_x)
	//fmt.Printf("division size x: %d\n", divisionSize)
	if divisionSize < coordSystem.source.divisionSizeX {
		divisionSize = coordSystem.source.divisionSizeX
	}
	legend_i := 0

	for i := coo.legendView.lower_x; i <= coo.legendView.upper_x; i++ {
		// draw the tick
		coo_x_start = coo.margin_l + (i-coo.legendView.lower_x)*divisionSize
		coo_x_stop = uint(coo_x_start)
		// len of tick = 10 (5+5)
		coo_y_start = coo.target.height - coo.margin_bottom + tickLen
		coo_y_stop = uint(coo.target.height) - uint(coo.margin_bottom) - uint(tickLen)
		x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
		// draw the legend
		var textDimensions x11.XCharStruct
		var dir int
		var ascent int
		var descent int
		if i >= coordSystem.legendView.lower_x && i < coordSystem.legendView.upper_x {
			ctext := coo.legend_x[legend_i+coordSystem.legendView.lower_x]
			x11.XTextExtents(font, ctext, len(ctext), &dir, &ascent, &descent, &textDimensions)
			txt_x_start := coo_x_start - int(textDimensions.Width/2)
			x11.XDrawString(d, drawable, gc, txt_x_start, coo_y_start+15, ctext)
			legend_i++
		}
	}
}

// drawXAxis draws Y axis of a coordinate system
func (coo *CoordSystemInfo) drawYAxis(d *x11.Display, drawable x11.Drawable, gc x11.GC) {
	gContextId := x11.XGContextFromGC(gc)
	font := x11.XQueryFont(d, *(*x11.XID)(unsafe.Pointer(&gContextId)))
	tickLen := coo.len_tick / 2

	//TODO
	// check numeric values below (5,15) these offsets need to be calculated by font size

	coo_x_start := coo.margin_l
	coo_x_stop := uint(coo.margin_l)
	coo_y_start := coo.target.height - coo.margin_bottom
	coo_y_stop := uint(coo.margin_top)
	x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
	// ticks
	divisionSize := coordSystem.source.divisionSizeY * (coo.source.height - coo.margin_bottom - coo.margin_top) / (coo.view.upper_y - coo.view.lower_y)
	//fmt.Printf("division size y: %d\n", divisionSize)
	if divisionSize < coordSystem.source.divisionSizeY {
		divisionSize = coordSystem.source.divisionSizeY
	}
	legend_i := 0
	for i := coo.legendView.lower_y; i < coo.legendView.upper_y; i++ {
		// draw the tick
		coo_y_start = coo.target.height - coo.margin_bottom - (i-coo.legendView.lower_y+1)*divisionSize
		coo_y_stop = uint(coo_y_start)
		// len of tick = 10 (5+5)
		coo_x_start = int(coo.margin_l) + tickLen
		coo_x_stop = uint(coo.margin_l) - uint(tickLen)
		x11.XDrawLine(d, drawable, gc, int(coo_x_start), coo_y_start, coo_x_stop, coo_y_stop)
		// draw the legend
		var textDimensions x11.XCharStruct
		var dir int
		var ascent int
		var descent int
		if i >= coordSystem.legendView.lower_y && i <= coordSystem.legendView.upper_y {
			ctext := coo.legend_y[legend_i+coordSystem.legendView.lower_y+1]
			x11.XTextExtents(font, ctext, len(ctext), &dir, &ascent, &descent, &textDimensions)
			txt_x_start := int(coo.margin_l) - int(textDimensions.Width) - 5
			x11.XDrawString(d, drawable, gc, txt_x_start, coo_y_start+15, ctext)
			legend_i++
		}
	}
}
