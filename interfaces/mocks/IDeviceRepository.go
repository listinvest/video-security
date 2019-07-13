package mocks

import (
	"github.com/prohevg/video-security/models"
	"github.com/stretchr/testify/mock"
)

//IDeviceRepository devices in db
type IDeviceRepository struct {
	mock.Mock
}

//AddOrUpdate add to db
func (_m *IDeviceRepository) AddOrUpdate(m models.Device) (models.Device, error) {
	args := _m.Called(m)
	return args.Get(0).(models.Device), args.Error(1)
}

//Remove into db
func (_m *IDeviceRepository) Remove(ip string) error{
	args := _m.Called(ip)
	return args.Error(0)
}

//Get from db
func (_m *IDeviceRepository) Get(ip string) (models.Device, error) {
	args := _m.Called(ip)
	return args.Get(0).(models.Device), args.Error(1)
}

//GetAll all
func (_m *IDeviceRepository) GetAll() ([]models.Device, error) {
	args := _m.Called()
	return args.Get(0).([]models.Device), args.Error(1)
}