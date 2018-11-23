package procps

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	//"time"
)

const (
	proc      = "/proc"
	stat      = "stat"
	status    = "status"
	pidpath   = "/proc/%d/%s"
	uidstr    = "Uid:"
	userspath = "/etc/passwd"
)

type Process struct {
	Cmd       string
	Cpuusage  string
	Cpu       float64
	Username  string
	Uid       string
	Pid       int
	Utime     float64
	Putime    float64
	Stime     float64
	Pstime    float64
	Cutime    float64
	Cstime    float64
	Previous  float64
	Current   float64
	Starttime float64
	Vmrss     int
	Puptime   float64
}

func newProcess() *Process {
	proc := new(Process)
	proc.Cmd = ""
	proc.Username = ""
	proc.Utime = -1
	proc.Stime = -1
	proc.Vmrss = -1
	proc.Pid = -1
	return proc
}

func calculateCPU(pid *Process) {
	curr := float64(pid.Utime + pid.Stime)
	stime := pid.Stime - pid.Pstime
	utime := pid.Utime - pid.Putime
	uptime := getUptime() - pid.Puptime
	pid.Cpu = float64(float64(stime+utime)/float64(100.0)) / uptime * float64(100.0)
	pid.Cpuusage = fmt.Sprintf("%.1f", pid.Cpu)
	pid.Puptime = getUptime()
	pid.Previous = curr
}

func getUptime() float64 {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return -1
	}

	times := strings.Split(string(b), " ")

	s, err := strconv.ParseFloat(times[0], 64)

	if err != nil {
		return -1
	}
	return s
}

//curtosy of stackoverflow.com/questions/11356330/getting-cpu-usage-with-golang
func getHertz() float64 {
	contents, err := ioutil.ReadFile("/proc/stat")
	total := 0
	if err != nil {
		return -1
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += int(val)
			}
			return float64(total)
		}
	}
	return float64(total)
}

func readPos(pos int, arr []string) float64 {
	val, err := strconv.Atoi(arr[pos])
	if err != nil {
		return -1
	}
	return float64(val)
}

func getUsers() map[string]string {
	users := make(map[string]string)
	namepos := 0
	uidpos := 2
	usercontent, err := os.Open(userspath)
	if err == nil {
		scanner := bufio.NewScanner(usercontent)
		if scanner.Scan() {
			for scanner.Scan() {
				text := strings.Split(scanner.Text(), ":")
				users[text[uidpos]] = text[namepos]
			}
		}
	}
	return users
}
