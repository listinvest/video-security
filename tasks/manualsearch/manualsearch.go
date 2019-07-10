package manualsearch

import (
	"fmt"
	"net/http"
	"time"

	"videoSecurity/tasks/base"
	"videoSecurity/tasks/models"
	"videoSecurity/tasks/manualsearch/ipparse"
	"videoSecurity/tasks/manualsearch/portparse"
)

const (
	urlPattern = "http://%s:%v/onvif/device_service"
)

//DeviceTask search devices in network by paramerters
type DeviceTask struct {
	Ips    string
	Ports  string
	Task   base.BizTask
	Result models.DeviceTaskResult
}

//GetID ID task
func (task *DeviceTask) GetID() string {
	return task.Task.ID
}

//Run run task
func (task *DeviceTask) Run() {
	result := make([]models.Device, 0)

	ips, err := ipparse.GetArrayIP(task.Ips)
	if err != nil {
		task.createErrorResult(result, err)
		return
	}

	ports, err := portparse.GetArrayPort(task.Ports)
	if err != nil {
		task.createErrorResult(result, err)
		return
	}

	for _, ip := range ips {
		for _, port := range ports {

			if task.Task.IsCanceled {
				task.createSuccessResult(result)
				return
			}

			endpoint := fmt.Sprintf(urlPattern, ip, port)

			fmt.Println("Request ", endpoint)

			err := task.ping(endpoint)
			if err != nil {
				fmt.Println(err)
				continue
			}

			newDevice := models.Device {
				Xaddr:  endpoint,
			}
			result = append(result, newDevice)
		}
	}

	task.Task.IsCompete = true
	task.createSuccessResult(result)
}

//Abort executing task
func (task *DeviceTask) Abort() {
	task.Task.IsCanceled = true
}

//IsCompete true if complete
func (task *DeviceTask) IsCompete() bool {
	return task.Task.IsCompete
}

// {{{ inner implementation

//Ping url to search for a device
func (task *DeviceTask) ping(endpoint string) (err error) {
	httpClient := new(http.Client)
	httpClient.Timeout = time.Duration(1 * time.Second)

	_, err = httpClient.Get(endpoint)
	if err != nil {
		return err
	}

	return nil
}

//createResult error
func (task *DeviceTask) createErrorResult(devices []models.Device, err error) {
	task.Result = models.DeviceTaskResult {
		Devices: devices,
		Result: base.BizTaskResult {
			IsOk: false,
			IsError: true,
			Error: err.Error(),
		},
	}
}

//createResult success
func (task *DeviceTask) createSuccessResult(devices []models.Device) {
	task.Result = models.DeviceTaskResult {
		Devices: devices,
		Result: base.BizTaskResult {
			IsOk: true,
		},
	}
}

// }}}