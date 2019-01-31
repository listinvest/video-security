package main

import (
	"fmt"

	"./autofind"
	"./tasks"
)

func main() {
	bizTask := autofind.DeviceTask{}

	dispatcher := tasks.GetInstance()
	f := dispatcher.RunTask(bizTask)

	fmt.Println("count tasks=", dispatcher.Count())

	deviceTaskResult := f.Result.(autofind.DeviceTaskResult)
	fmt.Println("count tasks=", deviceTaskResult)

	//	result := bizTask.GetResult()
	//	deviceTaskResult := result.(autofind.DeviceTaskResult)
	//	fmt.Println("result tasks=", deviceTaskResult)

	dispatcher.AbortTask(bizTask.BizTask.ID)

	fmt.Println("count tasks=", dispatcher.Count())
}
