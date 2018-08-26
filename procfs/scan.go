package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func initialScan() map[int]*Process {
	items, err := ioutil.ReadDir(proc)
	pids := make(map[int]*Process)
	if err != nil {
		return pids //[]Process{}
	}

	//pids[0] = *newProcess()
	//i := 0

	for _, item := range items {
		pid := newProcess()
		pid.Pid = atoiOr(item.Name(), -1)
		if pid.Pid > 0 {
			scanStatus(pid)
			scanCmdLine(pid)

			if pid.Username != "" && pid.Cmd != "" {
				scanStat(pid)
				prev := pid.Utime + pid.Stime
				pid.Previous = prev
				pids[pid.Pid] = pid
			}
		}
	}

	time.Sleep(time.Second)

	Scan(pids)

	fmt.Print("========================\n")

	for k, v := range pids {
		if v.Cpuusage > 0 {
			fmt.Print(k)
			fmt.Print(" ")
			fmt.Print(v.Cpuusage)
			fmt.Print("\n")
		}
	}

	return pids //[]Process{}
}

func Scan(items map[int]*Process) {
	for k := range items {
		p := items[k]
		scanStat(p)
		calculateCPU(p)
	}
}

func atoiOr(s string, alt int) int {
	value, err := strconv.Atoi(s)
	if err == nil {
		return value
	}
	return alt
}

func scanCmdLine(pid *Process) {
	cmdstr := ""
	b, err := ioutil.ReadFile(fmt.Sprintf(pidpath, pid.Pid, "cmdline"))

	if err == nil {
		cmdstr = string(b)
		pid.Cmd = cmdstr
	}
}

func scanStat(pid *Process) {
	b, err := os.Open(fmt.Sprintf(pidpath, pid.Pid, stat))
	defer b.Close()

	if err == nil {
		scanner := bufio.NewScanner(b)

		if scanner.Scan() {
			arr := strings.Fields(scanner.Text())
			if len(arr) > 22 {
				for i := 13; i < 17; i++ {
					val := readPos(i, arr)
					switch i {
					case 13:
						pid.Utime = val

					case 14:
						pid.Stime = val

					case 15:
						pid.Cutime = val

					case 16:
						pid.Cstime = val
					}

					pid.Starttime = readPos(21, arr)
				}
			}
		}
	}
}

func scanStatus(pid *Process) {
	uidstr := "Uid:"

	statuscontent, err := os.Open(fmt.Sprintf(pidpath, pid.Pid, status))

	if err == nil {
		defer statuscontent.Close()
		scanner := bufio.NewScanner(statuscontent)

		if scanner.Scan() {

			for scanner.Scan() {
				text := strings.Fields(scanner.Text())
				switch text[0] {
				case uidstr:
					pid.Uid = text[1]
					pid.Username = getUsers()[pid.Uid]
				}
			}
		}
	}
}
