package main

import (
	
	ui "github.com/gizak/termui"
	proc "github.com/Hunt4Bugs/gtop/pkg/procps"
	"time"
	//"strconv"
)

func getUI(header, arr []string) (*ui.List,*ui.List){
	if err := ui.Init(); err != nil {
		panic(err)
	}
	maxx := ui.TermWidth()
	maxy := ui.TermHeight()

	quarter := 4

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

	// quits when q is pressed
	ui.Handle("q", func(ui.Event) {
		ui.StopLoop()
	})
	return h,table
}

func main() {

	pids := proc.InitialScan()

	arr := proc.Format(pids)
	header := proc.GetDeviceInfo()

	h,table := getUI(header,arr)

	defer ui.Close()
	
	drawTicker := time.NewTicker(time.Second)

	go func(){
		for{
			proc.Scan(pids)
			arr = proc.Format(pids)
			header = proc.GetDeviceInfo()
			table.Items = arr
			h.Items = header
			ui.Render(table, h)
			time.Sleep(time.Second)
			<-drawTicker.C
		}
	}()

	ui.Loop()
}
