package taskdispatcher

import (
	"videoSecurity/tasks/base"
)

//ITaskDispatcher interface
type ITaskDispatcher interface {
	//run
	RunTask(task base.IBizTaskRunner)
	//run async
	RunAsyncTask(task base.IBizTaskRunner)
	//wait async task
	Wait(task base.IBizTaskRunner)
	//return task if she's running
	GetTask(taskID string) interface{}
	//stop task
	AbortTask(taskID string)
	//tasks in holder
	Count() int
}