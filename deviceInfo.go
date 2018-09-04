package main

import (
	"fmt"
)

type DeviceInfo struct {
	swapSize int
	swapFree int
	swapUsed int
	memSize  int
	memFree  int
	memUsed  int
	buffer   int
	cache    int
}

func getDeviceInfo() []string {
	var arr []string
	var d DeviceInfo

	//d.swapSize, d.swapUsed = getSwap()
	getMem(&d)

	arr = append(arr, fmt.Sprintf("KiB Mem : %d total, %d free, %d used, %d buff/cache", d.memSize, d.memFree, d.memUsed, d.buffer+d.cache))
	arr = append(arr, fmt.Sprintf("KiB Swap: %d total, %d free, %d used", d.swapSize, d.swapFree, d.swapUsed))

	return arr
}
