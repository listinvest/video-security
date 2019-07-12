package repositories

import (
	"errors"
	"fmt"
	"testing"

	"videoSecurity/interfaces/mocks"
	"videoSecurity/models"

	"github.com/stretchr/testify/assert"
)

//TestDeviceRepositoryAddOrUpdateErrorIfKeyEmpty save item with empty key
//SUCCESS IF RETURN ERRORS
func TestDeviceRepositoryAddOrUpdateErrorIfKeyEmpty(t *testing.T) {
	expected := models.Device {
		IP: "",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler

	h := TestHelper{}
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("AddOrUpdate", expected.IP, &expected).Return(errors.New("key is empty"))

	_, err := rep.AddOrUpdate(expected)
	assert.Error(t, err)
}

//TestDeviceRepositoryAddOrUpdateSuccess save item success
//SUCCESS IF RETURN WITHOUT ERRORS
func TestDeviceRepositoryAddOrUpdateSuccess(t *testing.T) {
	expected := models.Device {
		IP: "192.168.11.4",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler

	h := TestHelper{}
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("AddOrUpdate", expected.IP, &expected).Return(nil)

	_, err := rep.AddOrUpdate(expected)
	assert.NoError(t, err)
}

//TestDeviceRepositoryGetErrorIfKeyEmpty get item with empty key
//SUCCESS IF RETURN ERRORS
func TestDeviceRepositoryGetErrorIfKeyEmpty(t *testing.T) {
	h := TestHelper{}

	expected := models.Device {
		IP: "",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("Get", expected.IP, rep).Return(&expected, errors.New("key is empty"))

	_, err := rep.Get(expected.IP)
	assert.Error(t, err)
}

//TestDeviceRepositoryGetSuccess get item by key
//SUCCESS IF RETURN WITHOUT ERRORS
func TestDeviceRepositoryGetSuccess(t *testing.T) {
	h := TestHelper{}

	expected := models.Device {
		IP: "192.168.11.4",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("Get", expected.IP, rep).Return(&expected, nil)

	real, err := rep.Get(expected.IP)
	assert.NoError(t, err)
	assert.True(t, expected.IP == real.IP)
}

//TestDeviceRepositoryRemoveErrorIfKeyEmpty remove item by empty key
//SUCCESS IF RETURN ERRORS
func TestDeviceRepositoryRemoveErrorIfKeyEmpty(t *testing.T) {
	h := TestHelper{}

	expected := models.Device {
		IP: "",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("Remove", expected.IP).Return(errors.New("key is empty"))

	err := rep.Remove(expected.IP)
	assert.Error(t, err)
}

//TestDeviceRepositoryRemoveErrorIfKeyFake remove item by fake key
//SUCCESS IF RETURN ERRORS
func TestDeviceRepositoryRemoveErrorIfKeyFake(t *testing.T) {
	h := TestHelper{}

	expected := models.Device {
		IP: "",
	}

	expError := fmt.Errorf("not found record with key %s", expected.IP)

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("Remove", expected.IP).Return(expError)

	err := rep.Remove(expected.IP)
	assert.Error(t, err)
}

//TestDeviceRepositoryRemoveSuccess remove item by key
//SUCCESS IF RETURN WITHOUT ERRORS
func TestDeviceRepositoryRemoveSuccess(t *testing.T) {
	h := TestHelper{}

	expected := models.Device {
		IP: "192.168.11.4",
	}

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("Remove", expected.IP).Return(nil)

	err := rep.Remove(expected.IP)
	assert.NoError(t, err)
}

//TestDeviceRepositoryGetAllWasEmpty return empty slice
//SUCCESS IF RETURN WITHOUT ERRORS
func TestDeviceRepositoryGetAllWasEmpty(t *testing.T) {
	h := TestHelper{}

	expected := make([]interface{}, 0)

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("GetAll", rep).Return(expected, nil)

	real, err := rep.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(real) == 0)
}

//TestDeviceRepositoryGetAllNotEmpty return slice len = 2
//SUCCESS IF RETURN WITHOUT ERRORS
func TestDeviceRepositoryGetAllNotEmpty(t *testing.T) {
	h := TestHelper{}

	expected := make([]interface{}, 0)
	exp0 := models.Device {
		IP: "192.168.11.4",
	}
	exp1 := models.Device {
		IP: "192.168.11.5",
	}

	expected = append(expected, &exp0)
	expected = append(expected, &exp1)

	baserep := new(mocks.IBaseRepository)
	handler := new(mocks.IDbHandler)
	baserep.IDbHandler = handler
	rep := h.CreateDeviceRepository(baserep)

	baserep.On("GetAll", rep).Return(expected, nil)

	real, err := rep.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(real) == 2)
	assert.Equal(t, real[0], exp0)
	assert.Equal(t, real[1], exp1)
}