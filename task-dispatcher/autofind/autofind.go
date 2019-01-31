package autofind

import (
	"fmt"

	"../tasks"
	"github.com/yakovlevdmv/goonvif"
	//"github.com/yakovlevdmv/goonvif"
)

type device struct {
	xaddr    string
	login    string
	password string
}

//DeviceTask description task
type DeviceTask struct {
	BizTask tasks.BizTask
}

//DeviceTaskResult description task
type DeviceTaskResult struct {
	Devices []device
}

//GetID run task
func (task DeviceTask) GetID() string {
	fmt.Println("...and DeviceTask GetTaskID")
	return task.BizTask.ID
}

//Run run task
func (task DeviceTask) Run() tasks.BizTask {
	fmt.Println("DeviceTask RunTask")

	devices := goonvif.GetAvailableDevicesAtSpecificEthernetInterface("0.0.0.0")

	//result := []device{
	//	device{
	//		xaddr: "dev.GetEndpoint",
	//	},
	//}

	result := []device{}

	for _, dev := range devices {
		newDevice := device{}
		newDevice.xaddr = dev.GetEndpoint("Device")
		result = append(result, newDevice)
	}

	task.BizTask.Result = DeviceTaskResult{
		Devices: result,
	}

	return task.BizTask
}

//Abort run task
func (task DeviceTask) Abort() {
	fmt.Println("...and DeviceTask AbortTask")
}
