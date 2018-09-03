package main

import ()

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

func getDeviceInfo() DeviceInfo {
	var d DeviceInfo

	//d.swapSize, d.swapUsed = getSwap()
	getMem(&d)
	return d
}
