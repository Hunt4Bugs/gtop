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

	pid, uid, cpu, mem, coms := format(pids)

	// pid column
	pcol := ui.NewList()
	pcol.Border = false
	pcol.Width = 6
	pcol.Height = maxy - quarter
	pcol.Items = pid
	pcol.X = 0
	pcol.Y = quarter

	// cpu column
	ccol := ui.NewList()
	ccol.Border = false
	ccol.Width = 7
	ccol.Height = maxy - quarter
	ccol.Items = cpu
	ccol.X = 16
	ccol.Y = quarter

	//command column
	comcol := ui.NewList()
	comcol.Border = false
	comcol.Width = 20
	comcol.Height = maxy - quarter
	comcol.Items = coms
	comcol.X = 32
	comcol.Y = quarter

	// username column
	ucol := ui.NewList()
	ucol.Border = false
	ucol.Width = 7
	ucol.Height = maxy - quarter
	ucol.Items = uid
	ucol.X = 8
	ucol.Y = quarter

	memcol := ui.NewList()
	memcol.Border = false
	memcol.Width = 7
	memcol.Height = maxy - quarter
	memcol.Items = mem
	memcol.X = 24
	memcol.Y = quarter

	//Render view
	ui.Render(pcol, ucol, ccol, memcol, comcol)

	// quits when q is pressed
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	// interval loop to update everything
	ui.Handle("/timer/1s", func(e ui.Event) {
		//t := e.Data.(ui.EvtTimer)
		//i := 0
		Scan(pids)
		pid, uid, cpu, mem, coms = format(pids)
		pcol.Items = pid
		ucol.Items = uid
		ccol.Items = cpu
		comcol.Items = coms
		memcol.Items = mem
		ui.Render(pcol, ucol, ccol, memcol, comcol)
	})
	ui.Loop()
}
