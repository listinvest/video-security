package main

import (
	"fmt"

	"./autofind"
	"./tasks"
)

func main() {
	bizTask := autofind.DeviceTask{
		Task: tasks.BizTask{
			ID:   "1",
			Name: "test",
		},
	}

	dispatcher := tasks.GetInstance()
	dispatcher.RunTask(&bizTask)

	fmt.Println("count tasks=", dispatcher.Count())

	devices := bizTask.Result.Devices
	for _, device := range devices {
		fmt.Println("device xaddres=", device.Xaddr)
	}

	dispatcher.AbortTask(bizTask.Task.ID)

	fmt.Println("count tasks=", dispatcher.Count())
}
