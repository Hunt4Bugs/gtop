package main

import (
	ui "github.com/gizak/termui"
	//"strconv"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	//maxx := ui.TermWidth()
	maxy := ui.TermHeight()

	quarter := 0 //int(float64(maxy) * 0.30)

	defer ui.Close()

	pids := initialScan()

	pid, cpu, coms := format(pids)

	pcol := ui.NewList()
	pcol.Border = false
	pcol.Width = 6
	pcol.Height = maxy - quarter
	pcol.Items = pid
	pcol.X = 0
	pcol.Y = quarter

	ccol := ui.NewList()
	ccol.Border = false
	ccol.Width = 7
	ccol.Height = maxy - quarter
	ccol.Items = cpu
	ccol.X = 8
	ccol.Y = quarter

	comcol := ui.NewList()
	comcol.Border = false
	comcol.Width = 20
	comcol.Height = maxy - quarter
	comcol.Items = coms
	comcol.X = 16
	comcol.Y = quarter
	ui.Render(pcol, ccol, comcol)

	//Handlers below

	// quits when q is pressed
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	// interval loop to update everything
	ui.Handle("/timer/1s", func(e ui.Event) {
		//t := e.Data.(ui.EvtTimer)
		//i := 0
		Scan(pids)
		pid, cpu, coms = format(pids)
		pcol.Items = pid
		ccol.Items = cpu
		comcol.Items = coms
		ui.Render(pcol, ccol, comcol)
	})
	ui.Loop()
}
