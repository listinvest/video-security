package main

import (
	"fmt"

	"github.com/yakovlevdmv/goonvif"
)

func main() {

	_, err := goonvif.NewDevice("192.168.11.180:80")
	if err != nil {
		fmt.Println(err)
	}

	//devices := goonfiv.GetAvailableDevicesAtSpecificEthernetInterface("0.0.0.0")

	//for _, device := range devices {
	//	xaddr := device.xaddr
	//	fmt.Println("endpoint: %s", xaddr)
	//endpoint := device.GetEndpoint("")
	//fmt.Println("endpoint: %s", endpoint)
	//}
}
