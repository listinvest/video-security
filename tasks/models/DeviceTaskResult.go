package models

import (
	"videoSecurity/tasks/base"
)

//DeviceTaskResult search results
type DeviceTaskResult struct {
	Devices []Device
	Result  base.BizTaskResult
}