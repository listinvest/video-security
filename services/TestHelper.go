package services

import (
	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
)

//TestHelper creator test models
type TestHelper struct {
}

//CreateTestDeviceService create
func (h *TestHelper) CreateTestDeviceService(rep interfaces.IDeviceRepository) interfaces.IDeviceService {
	return &DeviceService{
		&logwriter.Logger{},
		rep,
	}
}