package models

import (
	"github.com/prohevg/video-security/tasks/base"
)

//DeviceTaskResult search results
type DeviceTaskResult struct {
	Devices []Device
	Result  base.BizTaskResult
}