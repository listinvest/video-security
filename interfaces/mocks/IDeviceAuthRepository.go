package mocks

import (
	"github.com/prohevg/video-security/models"
	"github.com/stretchr/testify/mock"
)

//IDeviceAuthRepository DeviceAuths in db
type IDeviceAuthRepository struct {
	mock.Mock
}

//AddOrUpdate add to db
func (_m *IDeviceAuthRepository) AddOrUpdate(m models.DeviceAuth) (models.DeviceAuth, error) {
	args := _m.Called(m)
	return args.Get(0).(models.DeviceAuth), args.Error(1)
}

//Remove into db
func (_m *IDeviceAuthRepository) Remove(key string) error{
	args := _m.Called(key)
	return args.Error(0)
}

//Get from db
func (_m *IDeviceAuthRepository) Get(key string) (models.DeviceAuth, error) {
	args := _m.Called(key)
	return args.Get(0).(models.DeviceAuth), args.Error(1)
}

//GetAll all
func (_m *IDeviceAuthRepository) GetAll() ([]models.DeviceAuth, error) {
	args := _m.Called()
	return args.Get(0).([]models.DeviceAuth), args.Error(1)
}