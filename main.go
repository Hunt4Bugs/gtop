package main

import (
	ui "github.com/gizak/termui"
	"time"
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
	header := getDeviceInfo()

	h := ui.NewList()
	h.Border = false
	h.Width = maxx
	h.Height = 4
	h.Items = header
	h.X = 0
	h.Y = 0

	table := ui.NewList()
	table.Border = false
	//table.Separator = false
	table.Width = maxx
	table.Height = maxy - quarter
	table.Items = arr
	table.X = 0
	table.Y = quarter

	ui.Render(table, h)

	// quits when q is pressed
	ui.Handle("q", func(ui.Event) {
		ui.StopLoop()
	})

	
	drawTicker := time.NewTicker(time.Second)

	go func(){
		for{
			Scan(pids)
			arr = format(pids)
			header = getDeviceInfo()
			table.Items = arr
			h.Items = header
			ui.Render(table, h)
			<-drawTicker.C
		}
	}()

	/*/ interval loop to update everything
	ui.Handle("/timer/1s", func(e ui.Event) {
		//t := e.Data.(ui.EvtTimer)
		//i := 0
		Scan(pids)
		//pid, uid, cpu, mem, coms = format(pids)
		arr = format(pids)
		header = getDeviceInfo()
		/*pcol.Items = pid
		ucol.Items = uid
		ccol.Items = cpu
		comcol.Items = coms
		memcol.Items = mem *
		table.Items = arr
		h.Items = header
		ui.Render(table, h) //pcol, ucol, ccol, memcol, comcol)
	})*/
	ui.Loop()
}
