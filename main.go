package main

import (
	ui "github.com/gizak/termui"
	//tb "github.com/nsf/termbox-go"
	//"strconv"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	maxx := ui.TermWidth()
	maxy := ui.TermHeight()

	quarter := 4 //int(float64(maxy) * 0.20)

	defer ui.Close()

	pids := initialScan()

	arr := format(pids)

	//TODO: add block to top of page with system info (i.e. cpu, memory and swap)

	table := ui.NewList()
	table.Border = false
	//table.Separator = false
	table.Width = maxx
	table.Height = maxy - quarter
	table.Items = arr
	table.X = 0
	table.Y = quarter

	ui.Render(table)

	// quits when q is pressed
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	// interval loop to update everything
	ui.Handle("/timer/1s", func(e ui.Event) {
		//t := e.Data.(ui.EvtTimer)
		//i := 0
		Scan(pids)
		//pid, uid, cpu, mem, coms = format(pids)
		arr = format(pids)
		/*pcol.Items = pid
		ucol.Items = uid
		ccol.Items = cpu
		comcol.Items = coms
		memcol.Items = mem*/
		table.Items = arr
		ui.Render(table) //pcol, ucol, ccol, memcol, comcol)
	})
	ui.Loop()
}
