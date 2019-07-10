package services

import (
	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
)

//DeviceService devices in application
type DeviceService struct {
	Logger *logwriter.Logger
	IDeviceRepository interfaces.IDeviceRepository
}
