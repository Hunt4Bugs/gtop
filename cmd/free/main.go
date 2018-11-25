package main

import (
	"github.com/Hunt4Bugs/gtop/pkg/procps"
	"fmt"
)

func main(){
	arr := procps.GetDeviceInfo()
	fmt.Println(arr[0])
	fmt.Println(arr[1])
}