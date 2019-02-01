package manualsearch

import (
	"fmt"
	"strconv"

	taskDispatcher "../task-dispatcher"
	ipparse "./ipparse"
	portparse "./portparse"
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
	fmt.Println("...and DeviceTask GetTaskID")
	return task.Task.ID
}

//Run run task
func (task *DeviceTask) Run() {
	fmt.Println("DeviceTask RunTask")

	ips := ipparse.GetArrayIP(task.Ips)

	for _, ip := range ips {
		fmt.Println("ip in Run=", ip)
	}

	ports := portparse.GetArrayPort(task.Ports)
	for _, port := range ports {
		fmt.Println("port in Run=", port)
	}

	result := []Device{}

	for _, ip := range ips {
		for _, port := range ports {
			xaddr := ip + ":" + strconv.Itoa(port)
			_, err := goonvif.NewDevice(xaddr)
			if err != nil {
				fmt.Println(err)
				continue
			}

			newDevice := Device{}
			newDevice.Xaddr = xaddr
			result = append(result, newDevice)
		}
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
