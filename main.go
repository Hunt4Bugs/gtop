package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"sync"
	//"os/user"
)

var (
	done  = make(chan struct{})
	delay = kingpin.Flag("delay", "Set time delay in tenths of seconds").Short('d').Int()
)

const plname = "ProcessList"

/*
https://github.com/jroimartin/gocui/blob/master/_example
widgets.go is a good example of how to setup this project.
Widgets for cpu cores, memory, processes and more.
Basically it is a struct of fields with a Layout function
attached.
*/

func main() {
	var wg sync.WaitGroup
	kingpin.Parse()

	g, err := gocui.NewGui(gocui.OutputNormal)
	maxX, maxY := g.Size()
	splity := int(float64(maxY) * 0.25)
	if err != nil {
		log.Panicln(err)
	}

	pl := ProcessList_new(plname, g, 0, splity, maxX, maxY)
	fmt.Println(pl)

	defer g.Close()

	g.SetManager(pl)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	wg.Add(1)
	go func(d *ProcessList) {
		defer wg.Done()
		d.scanroutine(g)
	}(pl)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	wg.Wait()
}

/*func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	splity := int(float64(maxY) * 0.25)
	if v, err := g.SetView("top", 0, 0, maxX, splity); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		uname, errr := user.LookupId("1000")
		if errr != nil {
			return errr
		}
		fmt.Fprintln(v, uname.Username)
	}
	if v, err := g.SetView("bottom", 0, splity+1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world")
	}
	return nil
}*/

func quit(g *gocui.Gui, v *gocui.View) error {
	close(done)
	return gocui.ErrQuit
}
