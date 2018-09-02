package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func unique(intSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func format(items map[int]*Process) ([]string, []string, []string, []string) {
	// convert important struct values to columns
	var pid []string
	pid = append(pid, "PID")
	var uid []string
	uid = append(uid, "User")
	var cpu []string
	cpu = append(cpu, "CPU %")
	var command []string
	command = append(command, "Command")
	cpumap := make(map[string][]Process)

	keys := make([]float64, 0, len(items))
	for _, e := range items {
		keys = append(keys, e.Cpu)
		cpumap[e.Cpuusage] = append(cpumap[e.Cpuusage], *e)
	}

	sort.Float64s(keys)

	keys = unique(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		k := fmt.Sprintf("%.2f", keys[i])
		for _, e := range cpumap[k] {
			if e.Cpuusage != "NaN" {
				cpu = append(cpu, k)
				pid = append(pid, strconv.Itoa(e.Pid))
				uid = append(uid, e.Username)
				command = append(command, e.Cmd)
			}
		}
	}

	return pid, uid, cpu, command
}

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
				pid.Puptime = getUptime()
				pid.Previous = prev
				pids[pid.Pid] = pid
			}
		}
	}

	time.Sleep(time.Second)

	Scan(pids)

	/*fmt.Print("========================\n")

	for k, v := range pids {
		if v.Cpuusage > 0 {
			fmt.Print(k)
			fmt.Print(" ")
			fmt.Print(v.Cpuusage)
			fmt.Print("\n")
		}
	}*/

	return pids //[]Process{}
}

func Scan(items map[int]*Process) {
	for k := range items {
		p := items[k]
		p.Pstime = p.Stime
		p.Putime = p.Utime
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
