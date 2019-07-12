package services

import (
	"errors"
	"testing"

	"videoSecurity/interfaces/mocks"
	"videoSecurity/models"

	"github.com/stretchr/testify/assert"
)

//TestDeviceService_AddOrUpdate_EmptyIP add with empty ip
//SUCCESS IF RETURN ERRORS
func Test_DeviceService_AddOrUpdate_EmptyIP(t *testing.T) {
	s := DeviceService{}
	_, err := s.AddOrUpdate("", 37777)
	assert.Error(t, err)
}

//Test_DeviceService_AddOrUpdate_NotValidIP add with not valid ip
//SUCCESS IF RETURN ERRORS
func Test_DeviceService_AddOrUpdate_NotValidIP(t *testing.T) {
	s := DeviceService{}
	_, err := s.AddOrUpdate("122.55.45.999", 37777)
	assert.Error(t, err)
}

//Test_DeviceService_AddOrUpdate_EmptyPort add with empty port
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceService_AddOrUpdate_EmptyPort(t *testing.T) {
	s := DeviceService{}
	_, err := s.AddOrUpdate("192.168.11.4", 0)
	assert.Error(t, err)
}

//Test_DeviceService_AddOrUpdate_Success add success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceService_AddOrUpdate_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	ip := "192.168.11.4"
	port := 37777

	dev := models.Device{
		IP:   ip,
		Port: port,
	}
	rep.On("AddOrUpdate", dev).Return(dev, nil)

	real, err := s.AddOrUpdate(ip, port)

	assert.NoError(t, err)
	assert.Equal(t, dev, real)
}

//Test_DeviceService_Get_EmptyIP get with empty ip
//SUCCESS IF RETURN ERRORS
func Test_DeviceService_Get_EmptyIP(t *testing.T) {
	s := DeviceService{}
	_, err := s.Get("")
	assert.Error(t, err)
}

//Test_DeviceService_Get_NotFoundByIp get not found by guid
//SUCCESS IF RETURN ERRORS
func Test_DeviceService_Get_NotFoundByIp(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	ip := "127.0.0.1"
	rep.On("Get", ip).Return(models.Device{}, errors.New("not found"))

	_, err := s.Get(ip)
	assert.Error(t, err)
}

//Test_DeviceService_Get_Success get
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceService_Get_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	ip := "192.168.11.4"
	port := 37777

	expected := models.Device{
		IP:   ip,
		Port: port,
	}

	rep.On("Get", ip).Return(expected, nil)

	real, err := s.Get(ip)
	assert.NoError(t, err)
	assert.Equal(t, real, expected)
}

//Test_DeviceService_Remove_EmptyIp remove with empty ip
//SUCCESS IF RETURN ERRORS
func Test_DeviceService_Remove_EmptyIp(t *testing.T) {
	s := DeviceService{}
	err := s.Remove("")
	assert.Error(t, err)
}

//Test_DeviceService_Remove_NotFoundByIp remove not found by ip
//SUCCESS IF RETURN ERRORS
func Test_DeviceService_Remove_NotFoundByIp(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	ip := ""
	rep.On("Get", ip).Return(models.Device{}, errors.New("not found"))

	err := s.Remove(ip)
	assert.Error(t, err)
}

//Test_DeviceService_Remove_Success remove success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceService_Remove_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	ip := "127.0.0.1"
	rep.On("Remove", ip).Return(nil)

	err := s.Remove(ip)
	assert.NoError(t, err)
}

//Test_DeviceService_GetAll_Empty all success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceService_GetAll_Empty(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	devices := make([]models.Device, 0)
	rep.On("GetAll").Return(devices, nil)

	all, err := s.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, len(all), len(devices))
}

//Test_DeviceService_GetAll_Success all success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceService_GetAll_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceRepository)
	s := h.CreateTestDeviceService(rep)

	devices := make([]models.Device, 2)
	devices[0] = models.Device{
		IP:   "192.168.11.4",
		Port: 37777,
	}

	devices[1] = models.Device{
		IP:   "192.168.11.5",
		Port: 37777,
	}

	rep.On("GetAll").Return(devices, nil)

	real, err := s.GetAll()

	assert.NoError(t, err)
	assert.True(t, len(real) == 2)
	assert.Equal(t, real[0], devices[0])
	assert.Equal(t, real[1], devices[1])
}
