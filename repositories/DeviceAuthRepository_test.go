package repositories

import (
	"errors"
	"fmt"
	"testing"

	"videoSecurity/interfaces/mocks"
	"videoSecurity/models"

	"github.com/stretchr/testify/assert"
)

//Test_DeviceAuthRepository_AddOrUpdate_ErrorIfLoginEmpty save item with empty login
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthRepository_AddOrUpdate_ErrorIfLoginEmpty(t *testing.T) {
	expected := models.DeviceAuth {
		Login: "",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler

	h := TestHelper{}
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("AddOrUpdate", expected.Login, &expected).Return(errors.New("key is empty"))

	_, err := rep.AddOrUpdate(expected)
	assert.Error(t, err)
}

//Test_DeviceAuthRepository_AddOrUpdate_Success save item success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthRepository_AddOrUpdate_Success(t *testing.T) {
	expected := models.DeviceAuth {
		Login: "admin",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler

	h := TestHelper{}
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("AddOrUpdate", expected.Login, &expected).Return(nil)

	_, err := rep.AddOrUpdate(expected)
	assert.NoError(t, err)
}

//Test_DeviceAuthRepository_Get_ErrorIfKeyEmpty get item with empty key
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthRepository_Get_ErrorIfKeyEmpty(t *testing.T) {
	h := TestHelper{}

	expected := models.DeviceAuth {
		Login: "",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("Get", expected.Login, rep).Return(&expected, errors.New("key is empty"))

	_, err := rep.Get(expected.Login)
	assert.Error(t, err)
}

//Test_DeviceAuthRepository_Get_Success get item by key
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthRepository_Get_Success(t *testing.T) {
	h := TestHelper{}

	expected := models.DeviceAuth {
		Login: "admin",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("Get", expected.Login, rep).Return(&expected, nil)

	real, err := rep.Get(expected.Login)
	assert.NoError(t, err)
	assert.True(t, expected.Login == real.Login)
}

//Test_DeviceAuthRepository_Remove_ErrorIfKeyEmpty remove item by empty key
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthRepository_Remove_ErrorIfKeyEmpty(t *testing.T) {
	h := TestHelper{}

	expected := models.DeviceAuth {
		Login: "",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("Remove", expected.Login).Return(errors.New("key is empty"))

	err := rep.Remove(expected.Login)
	assert.Error(t, err)
}

//Test_DeviceAuthRepository_Remove_ErrorIfKeyFake remove item by fake key
//SUCCESS IF RETURN ERRORS
func Test_DeviceAuthRepository_Remove_ErrorIfKeyFake(t *testing.T) {
	h := TestHelper{}

	expected := models.DeviceAuth {
		Login: "",
	}

	expError := fmt.Errorf("not found record with key %s", expected.Login)

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("Remove", expected.Login).Return(expError)

	err := rep.Remove(expected.Login)
	assert.Error(t, err)
}

//Test_DeviceAuthRepository_Remove_Success remove item by key
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthRepository_Remove_Success(t *testing.T) {
	h := TestHelper{}

	expected := models.DeviceAuth {
		Login: "admin",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("Remove", expected.Login).Return(nil)

	err := rep.Remove(expected.Login)
	assert.NoError(t, err)
}

//Test_DeviceAuthRepository__GetAll_AllWasEmpty return empty slice
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthRepository_GetAll_AllWasEmpty(t *testing.T) {
	h := TestHelper{}

	expected := make([]interface{}, 0)

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("GetAll", rep).Return(expected, nil)

	real, err := rep.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(real) == 0)
}

//Test_DeviceAuthRepository__GetAll_AllNotEmpty return slice len = 2
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_DeviceAuthRepository__GetAll_AllNotEmpty(t *testing.T) {
	h := TestHelper{}

	expected := make([]interface{}, 0)
	exp0 := models.DeviceAuth {
		Login: "admin",
	}
	exp1 := models.DeviceAuth {
		Login: "admin222",
	}

	expected = append(expected, &exp0)
	expected = append(expected, &exp1)

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceAuthRepository(baserep)

	baserep.On("GetAll", rep).Return(expected, nil)

	real, err := rep.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(real) == 2)
	assert.Equal(t, real[0], exp0)
	assert.Equal(t, real[1], exp1)
}