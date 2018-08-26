package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	proc    = "/proc"
	status  = "status"
	pidpath = "/proc/%d/%s"
)

type Process struct {
	Name  string
	User  string
	UID   string
	PID   int
	UTime int
	STime int
	VSize int
}

func newProcess() *Process {
	proc := new(Process)
	proc.Name = ""
	proc.User = ""
	proc.UTime = -1
	proc.STime = -1
	proc.VSize = -1
	proc.PID = -1
	return proc
}

func Listpids() []Process {
	items, err := ioutil.ReadDir(proc)
	if err != nil {
		return []Process{}
	}

	pids := make([]Process, len(items))
	pids[0] = *newProcess()
	i := 0

	for _, item := range items {
		pid := newProcess()
		pid.PID = atoiOr(item.Name(), -1)
		if pid.PID > 0 {
			pids[i] = *pid
			i++
		}
	}

	if pids[0].PID > 0 {
		return pids[0:i]
	}

	return []Process{}
}

func atoiOr(s string, alt int) int {
	value, err := strconv.Atoi(s)
	if err == nil {
		return value
	}
	return alt
}

func UID(pid Process) string {
	uid := ""
	statuscontent, err := os.Open(fmt.Sprintf(pidpath, pid.PID, status))

	if err != nil {
		return uid
	}

	defer statuscontent.Close()
	scanner := bufio.NewScanner(statuscontent)

	if !scanner.Scan() {
		return uid
	}

	// TODO: read lines and find uid using bufio and possibly strings
	//parts := strings.Fields(scanner.Text())

	return uid
}

/*func StatOf(pid int) *Stat {
	result := newStat()

	stat, err := os.Open(fmt.Sprintf(pidpath))
	defer stat.Close()

	if err != nil {
		return result
	}

	scanner := bufio.NewScanner(stat)
	if !scanner.Scan() {
		return result
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < statHighestIndex {
		return result
	}
	result.Name = parts[statName][1 : len(parts[statName])-1]

	return result
}*/
