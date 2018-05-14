package main

import (
	"github.com/jroimartin/gocui"
)

type Panel struct {
	x1   int
	y1   int
	x2   int
	y2   int
	name string
	g    *gocui.Gui
}
