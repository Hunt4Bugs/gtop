package main

import (
	"./procfs"
	"fmt"
	"github.com/jroimartin/gocui"
	"math/rand"
	"time"
)

const (
	proc    = "/proc"
	path    = "/proc/%s"
	pidpath = "/proc/%d/%s"
)

type ProcessList struct {
	Panel
}

func (pl *ProcessList) Layout(g *gocui.Gui) error {
	v, err := g.SetView(pl.name, pl.x1, pl.y1, pl.x2, pl.y2)
	pids := procfs.Listpids()
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, pids)
	}
	return nil
}

func ProcessList_new(name string, g *gocui.Gui, x1, y1, x2, y2 int) *ProcessList {
	return &ProcessList{Panel: Panel{x1, y1, x2, y2, name, g}}
}

func (pl *ProcessList) scanroutine(g *gocui.Gui) {
	// start goroutine to scan processes continuously on delay
	for {
		// continuously scan
		select {
		case <-done:
			return
		case <-time.After(500 * time.Millisecond):
			//rescan stuff
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View(pl.name)
				n := rand.Intn(5000)
				if err != nil {
					return err
				}
				v.Clear()
				fmt.Fprintln(v, n)
				return nil
			})
		}
	}
}
