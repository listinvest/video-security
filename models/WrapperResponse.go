package models

//WrapperResponse response
type WrapperResponse struct {
	IsError      bool
	ErrorMessage string
	Data         interface{}
}
