package autosearch

import (
	"videoSecurity/tasks/base"
	"videoSecurity/tasks/models"

	"github.com/yakovlevdmv/goonvif"
)

//DeviceTask auto search devices in netwotk
type DeviceTask struct {
	Task   base.BizTask
	Result models.DeviceTaskResult
}

//GetID ID task
func (task *DeviceTask) GetID() string {
	return task.Task.ID
}

//Run run task
func (task *DeviceTask) Run() {
	devices := goonvif.GetAvailableDevicesAtSpecificEthernetInterface("0.0.0.0")

	result := make([]models.Device, len(devices))

	for i, dev := range devices {
		newDevice := models.Device{
			Xaddr: dev.GetEndpoint("Device"),
		}
		result[i] = newDevice
	}

	task.Task.IsCompete = true
	task.Result = models.DeviceTaskResult{
		Devices: result,
		Result: base.BizTaskResult{
			IsOk: true,
		},
	}
}

//Abort executing task
func (task *DeviceTask) Abort() {
	task.Task.IsCanceled = true
}

//IsCompete true if complete
func (task *DeviceTask) IsCompete() bool {
	return task.Task.IsCompete
}
