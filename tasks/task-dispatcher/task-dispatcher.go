package taskdispatcher

import (
	"sync"
)

//BizTaskRunner interface task
type BizTaskRunner interface {
	GetID() string
	Run()
	Abort()
	IsCompete() bool
}

//BizTask base info about task
type BizTask struct {
	ID         string
	Name       string
	IsCanceled bool
	IsCompete  bool
	Result     BizTaskResult
}

//BizTaskResult result of the task
type BizTaskResult struct {
	IsOk bool
}

//taskDispatcher holder contains tasks are executing
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
}

//RunTask run task async
func (dispatcher *taskDispatcher) RunAsyncTask(task BizTaskRunner) {
	ch := make(chan BizTaskRunner)
	go runTaskInternal(dispatcher, task, ch)
	go waitTaskInternal(dispatcher, ch)
}

//RunTask run task async
func (dispatcher *taskDispatcher) Wait(task BizTaskRunner) {
	for {
		if task.IsCompete() {
			break
		}
	}
}


//GetTask return task if she's running
func (dispatcher *taskDispatcher) GetTask(taskID string) interface{} {
	return firstOrDefault(dispatcher.Tasks, taskID)
}

//AbortTask stop task
func (dispatcher *taskDispatcher) AbortTask(taskID string) {
	task := firstOrDefault(dispatcher.Tasks, taskID)
	if task != nil {
		task.Abort()
		dispatcher.Tasks = delete(dispatcher.Tasks, task)
	}
}

//Count count tasks in holder
func (dispatcher *taskDispatcher) Count() int {
	return len(dispatcher.Tasks)
}

// {{{ inner implementation

//runTaskInternal run task in gorutine
func runTaskInternal(dispatcher *taskDispatcher, task BizTaskRunner, ch chan BizTaskRunner) {
	defer close(ch)
	dispatcher.Tasks = append(dispatcher.Tasks, task)
	task.Run()
	ch <- task
}

//waitTask wait task in gorutine
func waitTaskInternal(dispatcher *taskDispatcher, ch chan BizTaskRunner) {
	runner := <-ch
	taskID := runner.GetID()
	task := firstOrDefault(dispatcher.Tasks, taskID)
	if task != nil {
		dispatcher.Tasks = delete(dispatcher.Tasks, task)
	}
}

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
