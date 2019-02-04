package manualsearch

import (
	"fmt"
	"net/http"
	"time"

	taskDispatcher "../task-dispatcher"
	ipparse "./ipparse"
	portparse "./portparse"
)

const (
	url = "http://%s:%v/onvif/device_service"
)

//Device info about device
type Device struct {
	Xaddr string
}

//DeviceTask auto search devices in netwotk
type DeviceTask struct {
	Ips    string
	Ports  string
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
	fmt.Println("Manual search task GetID")
	return task.Task.ID
}

//Run run task
func (task *DeviceTask) Run() {
	fmt.Println("Manual search task Run")

	result := []Device{}

	ips, err := ipparse.GetArrayIP(task.Ips)
	if err != nil {
		fmt.Println(err)
		task.Result = createResult(false, result)
		return
	}

	ports, err := portparse.GetArrayPort(task.Ports)
	if err != nil {
		fmt.Println(err)
		task.Result = createResult(false, result)
		return
	}

	for _, ip := range ips {
		for _, port := range ports {

			endpoint := fmt.Sprintf(url, ip, port)

			fmt.Println("Request ", endpoint)

			err := ping(endpoint)
			if err != nil {
				fmt.Println(err)
				continue
			}

			newDevice := Device{}
			newDevice.Xaddr = endpoint
			result = append(result, newDevice)
		}
	}

	task.Result = createResult(true, result)
}

//Abort executing task
func (task DeviceTask) Abort() {
	fmt.Println("Manual search task Abort")
}

//Ping url to search for a device
func ping(endpoint string) (err error) {
	httpClient := new(http.Client)
	httpClient.Timeout = time.Duration(1 * time.Second)

	_, err = httpClient.Get(endpoint)
	if err != nil {
		return err
	}

	return nil
}

//createResult retutn result
func createResult(isOk bool, devices []Device) DeviceTaskResult {
	return DeviceTaskResult{
		Devices: devices,
		Result: taskDispatcher.BizTaskResult{
			IsOk: isOk,
		},
	}
}
