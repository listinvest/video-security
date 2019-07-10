package taskdispatcher

import (
	"videoSecurity/tasks/base"
)

//TaskDispatcher holder contains tasks are executing
type TaskDispatcher struct {
	Tasks []base.IBizTaskRunner
}

//RunTask run task
func (dispatcher *TaskDispatcher) RunTask(task base.IBizTaskRunner) {
	task.Run()
}

//RunAsyncTask run
func (dispatcher *TaskDispatcher) RunAsyncTask(task base.IBizTaskRunner) {
	ch := make(chan base.IBizTaskRunner)
	go runTaskInternal(dispatcher, task, ch)
	go waitTaskInternal(dispatcher, ch)
}

//Wait async task
func (dispatcher *TaskDispatcher) Wait(task base.IBizTaskRunner) {
	for {
		if task.IsCompete() {
			break
		}
	}
}

//GetTask return task if she's running
func (dispatcher *TaskDispatcher) GetTask(taskID string) interface{} {
	return firstOrDefault(dispatcher.Tasks, taskID)
}

//AbortTask stop task
func (dispatcher *TaskDispatcher) AbortTask(taskID string) {
	task := firstOrDefault(dispatcher.Tasks, taskID)
	if task != nil {
		task.Abort()
		dispatcher.Tasks = delete(dispatcher.Tasks, task)
	}
}

//Count tasks in holder
func (dispatcher *TaskDispatcher) Count() int {
	return len(dispatcher.Tasks)
}

// {{{ inner implementation

//runTaskInternal run task in gorutine
func runTaskInternal(dispatcher *TaskDispatcher, task base.IBizTaskRunner, ch chan base.IBizTaskRunner) {
	defer close(ch)
	dispatcher.Tasks = append(dispatcher.Tasks, task)
	task.Run()
	ch <- task
}

//waitTask wait task in gorutine
func waitTaskInternal(dispatcher *TaskDispatcher, ch chan base.IBizTaskRunner) {
	runner := <-ch
	taskID := runner.GetID()
	task := firstOrDefault(dispatcher.Tasks, taskID)
	if task != nil {
		dispatcher.Tasks = delete(dispatcher.Tasks, task)
	}
}

//firstOrDefault
func firstOrDefault(tasks []base.IBizTaskRunner, taskID string) base.IBizTaskRunner {
	for _, bizTask := range tasks {
		if bizTask.GetID() == taskID {
			return bizTask
		}
	}
	return nil
}

//delete
func delete(tasks []base.IBizTaskRunner, task base.IBizTaskRunner) []base.IBizTaskRunner {
	index := indexOf(tasks, task)
	return removeByIndex(tasks, index)
}

//indexOf
func indexOf(tasks []base.IBizTaskRunner, task base.IBizTaskRunner) int {
	for index, bizTask := range tasks {
		if bizTask.GetID() == task.GetID() {
			return index
		}
	}
	return -1
}

//removeByIndex
func removeByIndex(s []base.IBizTaskRunner, i int) []base.IBizTaskRunner {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// }}}
