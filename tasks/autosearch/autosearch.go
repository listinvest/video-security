package autosearch

import (
	"fmt"

	taskDispatcher "../task-dispatcher"
	"github.com/yakovlevdmv/goonvif"
)

//Device info about device
type Device struct {
	Xaddr    string
	Login    string
	Password string
}

//DeviceTask auto search devices in netwotk
type DeviceTask struct {
	Task   taskDispatcher.BizTask
	Result DeviceTaskResult
}

//DeviceTaskResult search results
type DeviceTaskResult struct {
	Devices []Device
	Result  taskDispatcher.BizTaskResult
}

//GetID ID task
func (task *DeviceTask) GetID() string {
	fmt.Println("...and DeviceTask GetTaskID")
	return task.Task.ID
}

//Run run task
func (task *DeviceTask) Run() {
	fmt.Println("DeviceTask RunTask")

	devices := goonvif.GetAvailableDevicesAtSpecificEthernetInterface("0.0.0.0")

	/* result := []Device{
		Device{
			Xaddr: "dev.GetEndpoint",
		},
	} */

	result := []Device{}

	for _, dev := range devices {
		newDevice := Device{}
		newDevice.Xaddr = dev.GetEndpoint("Device")
		result = append(result, newDevice)
	}

	task.Result = DeviceTaskResult{
		Devices: result,
		Result: taskDispatcher.BizTaskResult{
			IsOk: true,
		},
	}
}

//Abort executing task
func (task DeviceTask) Abort() {
	fmt.Println("...and DeviceTask AbortTask")
}
