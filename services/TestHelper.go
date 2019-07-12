package services

import (
	"videoSecurity/deviceonvif"
	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
)

//TestHelper creator test models
type TestHelper struct {
}

//CreateTestDeviceService create
func (h *TestHelper) CreateTestDeviceService(rep interfaces.IDeviceRepository, repAuth interfaces.IDeviceAuthRepository, deviceOnvif deviceonvif.IDeviceOnvif) interfaces.IDeviceService {
	return &DeviceService{
		&logwriter.Logger{},
		rep,
		repAuth,
		deviceOnvif,
	}
}

//CreateTestDeviceAuthService create
func (h *TestHelper) CreateTestDeviceAuthService(rep interfaces.IDeviceAuthRepository) interfaces.IDeviceAuthService {
	return &DeviceAuthService{
		&logwriter.Logger{},
		rep,
	}
}