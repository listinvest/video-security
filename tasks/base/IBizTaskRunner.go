package base

//IBizTaskRunner interface task
type IBizTaskRunner interface {
	GetID() string
	Run()
	Abort()
	IsCompete() bool
}