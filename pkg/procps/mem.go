package procps

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetMem(dev *DeviceInfo) { // (int, int, int, int, int) {
	// read /proc for the mem info
	f, err := os.Open("/proc/meminfo")
	defer f.Close()

	if err == nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			text := strings.Fields(scanner.Text())
			switch text[0] {
			case "MemTotal:":
				dev.memSize, err = strconv.Atoi(text[1])
			case "MemFree:":
				dev.memFree, err = strconv.Atoi(text[1])
			case "Buffers:":
				dev.buffer, err = strconv.Atoi(text[1])
			case "Cached:":
				dev.cache, err = strconv.Atoi(text[1])
			case "SwapTotal:":
				dev.swapSize, err = strconv.Atoi(text[1])
			case "SwapFree:":
				dev.swapFree, err = strconv.Atoi(text[1])
			}
		}
		dev.swapUsed = dev.swapSize - dev.swapFree
		dev.memUsed = dev.memSize - dev.memFree
	}

}
