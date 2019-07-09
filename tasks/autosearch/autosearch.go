package autosearch

import (
	"fmt"
	"net/url"
	"strconv"

	taskDispatcher "../task-dispatcher"
	"github.com/yakovlevdmv/goonvif"
)

//Device info about device
type Device struct {
	Xaddr    string
	Login    string
	Password string
	IP       string
	Port     int
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

	result := []Device{}

	for _, dev := range devices {
		newDevice := Device{}
		newDevice.Xaddr = dev.GetEndpoint("Device")

		u, err := url.Parse(newDevice.Xaddr)
		if err == nil {
			newDevice.IP = u.Hostname()
			number, _ := strconv.Atoi(u.Port())
			newDevice.Port = number
		}

		if newDevice.Port == 0 {
			newDevice.Port = 80
		}

		result = append(result, newDevice)
	}

	task.Task.IsCompete = true
	task.Result = DeviceTaskResult{
		Devices: result,
		Result: taskDispatcher.BizTaskResult{
			IsOk: true,
		},
	}
}

//Abort executing task
func (task *DeviceTask) Abort() {
	task.Task.IsCanceled = true
	fmt.Println("...and DeviceTask AbortTask")
}

//IsCompete true if complete
func (task *DeviceTask) IsCompete() bool {
	return task.Task.IsCompete
}
