package tasks

import (
	"sync"
)

//BizTaskRunner interface to execute task
type BizTaskRunner interface {
	GetID() string
	Run()
	Abort()
}

//BizTask description task
type BizTask struct {
	ID     string
	Name   string
	Result BizTaskResult
}

//BizTaskResult result task
type BizTaskResult struct {
	IsOk bool
}

type taskDispatcher struct {
	Tasks []BizTaskRunner
}

var instance *taskDispatcher
var once sync.Once

//GetInstance singlenton
func GetInstance() *taskDispatcher {
	once.Do(func() {
		instance = &taskDispatcher{}
	})
	return instance
}

//RunTask run task
func (dispatcher *taskDispatcher) RunTask(task BizTaskRunner) {
	task.Run()
	dispatcher.Tasks = append(dispatcher.Tasks, task)
}

//AbortTask stop task
func (dispatcher *taskDispatcher) AbortTask(taskID string) {
	task := firstOrDefault(dispatcher.Tasks, taskID)
	if task != nil {
		task.Abort()
		dispatcher.Tasks = delete(dispatcher.Tasks, task)
	}
}

//AbortTask stop execute task
func (dispatcher *taskDispatcher) Count() int {
	return len(dispatcher.Tasks)
}

// {{{ inner implementation

//firstOrDefault
func firstOrDefault(tasks []BizTaskRunner, taskID string) BizTaskRunner {
	for _, bizTask := range tasks {
		if bizTask.GetID() == taskID {
			return bizTask
		}
	}
	return nil
}

//delete
func delete(tasks []BizTaskRunner, task BizTaskRunner) []BizTaskRunner {
	index := indexOf(tasks, task)
	return removeByIndex(tasks, index)
}

//indexOf
func indexOf(tasks []BizTaskRunner, task BizTaskRunner) int {
	for index, bizTask := range tasks {
		if bizTask.GetID() == task.GetID() {
			return index
		}
	}
	return -1
}

//removeByIndex
func removeByIndex(s []BizTaskRunner, i int) []BizTaskRunner {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// }}}
