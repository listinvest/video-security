package autofind

import (
	"fmt"

	"../tasks"
	//"github.com/yakovlevdmv/goonvif"
)

type Device struct {
	Xaddr    string
	Login    string
	Password string
}

//DeviceTask description task
type DeviceTask struct {
	Task   tasks.BizTask
	Result DeviceTaskResult
}

//DeviceTaskResult description task
type DeviceTaskResult struct {
	Devices []Device
	Result  tasks.BizTaskResult
}

//GetID run task
func (task *DeviceTask) GetID() string {
	fmt.Println("...and DeviceTask GetTaskID")
	return task.Task.ID
}

//Run run task
func (task *DeviceTask) Run() {
	fmt.Println("DeviceTask RunTask")

	//devices := goonvif.GetAvailableDevicesAtSpecificEthernetInterface("0.0.0.0")

	result := []Device{
		Device{
			Xaddr: "dev.GetEndpoint",
		},
	}

	/*
		result := []Device{}

		for _, dev := range devices {
			newDevice := device{}
			newDevice.Xaddr = dev.GetEndpoint("Device")
			result = append(result, newDevice)
		}
	*/

	task.Result = DeviceTaskResult{
		Devices: result,
		Result: tasks.BizTaskResult{
			IsOk: true,
		},
	}
}

//Abort run task
func (task DeviceTask) Abort() {
	fmt.Println("...and DeviceTask AbortTask")
}
