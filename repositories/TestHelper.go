package repositories

import (
	"videoSecurity/interfaces"
)

//TestHelper creator test models
type TestHelper struct {
}

//CreateDeviceRepository create
func (h *TestHelper) CreateDeviceRepository(baserep interfaces.IBaseRepository) interfaces.IDeviceRepository {
	return &DeviceRepository{
		baserep,
	}
}

//CreateDeviceAuthRepository create
func (h *TestHelper) CreateDeviceAuthRepository(baserep interfaces.IBaseRepository) interfaces.IDeviceAuthRepository {
	return &DeviceAuthRepository{
		baserep,
	}
}
