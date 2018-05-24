package procfs

import (
	"io/ioutil"
	"strconv"
)

const (
	proc    = "/proc"
	pidpath = "/proc/%d/%s"
)

func Listpids() []int {
	items, err := ioutil.ReadDir(proc)
	if err != nil {
		return []int{}
	}

	pids := make([]int, len(items))
	pids[0] = -1
	i := 0

	for _, item := range items {
		pid := atoiOr(item.Name(), -1)
		if pid > 0 {
			pids[i] = pid
			i++
		}
	}

	if pids[0] > 0 {
		return pids[0:i]
	}

	return []int{}
}

func atoiOr(s string, alt int) int {
	value, err := strconv.Atoi(s)
	if err == nil {
		return value
	}
	return alt
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
