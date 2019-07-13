package services

import (
	"errors"
	"testing"

	"github.com/prohevg/video-security/interfaces/mocks"
	"github.com/prohevg/video-security/models"

	"github.com/stretchr/testify/assert"
)

//TestDeviceAuthService_AddOrUpdate_EmptyLogin add with empty login
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthService_AddOrUpdate_EmptyLogin(t *testing.T) {
	s := DeviceAuthService{}
	_, err := s.AddOrUpdate("", "")
	assert.Error(t, err)
}

//Test_DeviceAuthService_AddOrUpdate_EmptyPassword add with empty password
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthService_AddOrUpdate_EmptyPassword(t *testing.T) {
	s := DeviceAuthService{}
	_, err := s.AddOrUpdate("admin", "")
	assert.Error(t, err)
}

//Test_DeviceAuthService_AddOrUpdate_Success add success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthService_AddOrUpdate_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceAuthRepository)
	s := h.CreateTestDeviceAuthService(rep)

	login := "login"
	pass := "pass"

	expected := models.DeviceAuth{
		Login:   login,
		Password: pass,
	}

	rep.On("AddOrUpdate", expected).Return(expected, nil)

	real, err := s.AddOrUpdate(login, pass)

	assert.NoError(t, err)
	assert.Equal(t, expected, real)
}

//Test_DeviceAuthService_Remove_EmptyLogin remove with empty Login
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthService_Remove_EmptyLogin(t *testing.T) {
	s := DeviceAuthService{}
	err := s.Remove("")
	assert.Error(t, err)
}

//Test_DeviceAuthService_Remove_NotFoundByLogin remove not found by Login
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthService_Remove_NotFoundByLogin(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceAuthRepository)
	s := h.CreateTestDeviceAuthService(rep)

	login := ""
	rep.On("Get", login).Return(models.DeviceAuth{}, errors.New("not found"))

	err := s.Remove(login)
	assert.Error(t, err)
}

//Test_DeviceAuthService_Remove_Success remove success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthService_Remove_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceAuthRepository)
	s := h.CreateTestDeviceAuthService(rep)

	login := "admin"
	rep.On("Remove", login).Return(nil)

	err := s.Remove(login)
	assert.NoError(t, err)
}

//Test_DeviceAuthService_GetAll_Empty all success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthService_GetAll_Empty(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceAuthRepository)
	s := h.CreateTestDeviceAuthService(rep)

	devices := make([]models.DeviceAuth, 0)
	rep.On("GetAll").Return(devices, nil)

	all, err := s.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, len(all), len(devices))
}

//Test_DeviceAuthService_GetAll_Success all success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthService_GetAll_Success(t *testing.T) {
	h := TestHelper{}
	rep := new(mocks.IDeviceAuthRepository)
	s := h.CreateTestDeviceAuthService(rep)

	devices := make([]models.DeviceAuth, 2)
	devices[0] = models.DeviceAuth{
		Login:   "admin",
		Password:   "pass",
	}

	devices[1] = models.DeviceAuth{
		Login:   "admin222",
		Password:   "pass",
	}

	rep.On("GetAll").Return(devices, nil)

	real, err := s.GetAll()

	assert.NoError(t, err)
	assert.True(t, len(real) == 2)
	assert.Equal(t, real[0], devices[0])
	assert.Equal(t, real[1], devices[1])
}
