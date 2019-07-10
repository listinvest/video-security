package base

//BizTask base info about task
type BizTask struct {
	ID         string
	Name       string
	IsCanceled bool
	IsCompete  bool
	Result     BizTaskResult
}