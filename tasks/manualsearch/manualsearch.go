package manualsearch

import (
	"fmt"

	taskDispatcher "../task-dispatcher"
	ipparse "./ipparse"
	portparse "./portparse"
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

}

//Abort executing task
func (task DeviceTask) Abort() {
	fmt.Println("...and DeviceTask AbortTask")
}
