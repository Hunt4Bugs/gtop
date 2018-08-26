package main

import (
	ui "github.com/gizak/termui"
	"strconv"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	//maxx := ui.TermWidth()
	maxy := ui.TermHeight()

	quarter := int(float64(maxy) * 0.30)

	defer ui.Close()

	var strs []string

	pids := Listpids()

	i := 0
	for i = 0; i < len(pids); i++ {
		strs = append(strs, strconv.Itoa(pids[i].PID))
	}

	ls := ui.NewList()
	ls.Border = false
	ls.Width = 5
	ls.Height = maxy - quarter
	ls.Items = strs
	ls.X = 0
	ls.Y = quarter

	ui.Render(ls)

	//Handlers below

	// quits when q is pressed
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	// interval loop to update everything
	ui.Handle("/timer/1s", func(e ui.Event) {
		//t := e.Data.(ui.EvtTimer)
		i := 0
		strs := make([]string, 0)
		for i = 0; i < len(pids); i++ {
			strs = append(strs, strconv.Itoa(pids[i].PID))
		}
		ls.Items = strs
		ui.Render(ls)
	})
	ui.Loop()
}
