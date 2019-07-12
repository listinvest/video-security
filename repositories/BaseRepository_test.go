package repositories

import (
	"errors"
	"videoSecurity/interfaces"
	"videoSecurity/models"
	"testing"

	"videoSecurity/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

//TestBaseGetKeyPrefixEmpty get empty prefix
//SUCCESS IF RETURN ERRORS
func TestBaseGetKeyPrefixEmpty(t *testing.T) {
	rep := BaseRepository{}
	_, err := rep.GetKeyPrefix()
	assert.Error(t, err)
}

//TestBaseGetKeyPrefixNotEmpty get isn't empty prefix
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseGetKeyPrefixNotEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = "some_prefix"
	real, err := rep.GetKeyPrefix()
	assert.NoError(t, err)
	assert.True(t, real != "")
	assert.Equal(t, real, rep.KeyPrefix)
}

//TestBaseGetKeyEmpty get isn't empty key
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseGetKeyEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = "some_prefix"
	_, err := rep.GetKey("")
	assert.Error(t, err)
}

//TestBaseGetKeyNotEmpty get isn't empty key
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseGetKeyNotEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = "some_prefix"
	keyName := "some_key"
	expected := rep.KeyPrefix + keyName
	real, err := rep.GetKey(keyName)
	assert.NoError(t, err)
	assert.True(t, real != "")
	assert.Equal(t, real, expected)
}

//TestBaseEncodeError encode error
//SUCCESS IF RETURN ERRORS
func TestBaseEncodeError(t *testing.T) {
	rep := BaseRepository{}
	_, err := rep.Encode(nil)
	assert.Error(t, err)
}

//TestBaseEncodeSuccess encode success
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseEncodeSuccess(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	_, err := rep.Encode(m)
	assert.NoError(t, err)
}

//TestBaseDecodeDataError decode error
//SUCCESS IF RETURN ERRORS
func TestBaseDecodeDataError(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	err := rep.Decode(nil, &m)
	assert.Error(t, err)
}

//TestBaseDecodeModelError decode error
//SUCCESS IF RETURN ERRORS
func TestBaseDecodeModelError(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	bytes, err := rep.Encode(m)
	assert.NoError(t, err)

	err = rep.Decode(bytes, nil)
	assert.Error(t, err)
}

//TestBaseDecodeSuccess encode success
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseDecodeSuccess(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	bytes, err := rep.Encode(m)
	assert.NoError(t, err)

	err = rep.Decode(bytes, &m)
	assert.NoError(t, err)
}

//TestBaseAddOrUpdateKeyEmpty save with empty key
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseAddOrUpdateKeyEmpty(t *testing.T) {
	rep := BaseRepository{}

	var e interface{}
	keyName := ""

	err := rep.AddOrUpdate(keyName, e)
	assert.Error(t, err)
}

//TestBaseAddOrUpdateSuccess save success
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseAddOrUpdateSuccess(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)

	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	rep.KeyPrefix = "some_prefix"

	m := &models.Device{}
	bytes, _ := rep.Encode(m)

	keyName, _ := rep.GetKey("some_key")
	dbHandler.On("AddOrUpdate", keyName, bytes).Return(nil)

	err := rep.AddOrUpdate(keyName, m)
	assert.NoError(t, err)
}

//TestBaseGetModelKeyEmpty get with empty key
//SUCCESS IF RETURN ERRORS
func TestBaseGetModelKeyEmpty(t *testing.T) {
	rep := BaseRepository{}
	mock := new(mocks.IBaseRepository)

	keyName := ""

	_, err := rep.Get(keyName, mock)
	assert.Error(t, err)
}

//TestBaseGetKeyFake get with fake key
//SUCCESS IF RETURN ERRORS
func TestBaseGetKeyFake(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	mock := new(mocks.IBaseRepository)

	keyName := "fake"
	dbHandler.On("Get", keyName, rep).Return(nil, nil)

	_, err := rep.Get(keyName, mock)
	assert.Error(t, err)
}

//TestBaseGetSuccess get with empty key
//SUCCESS IF RETURN ERRORS
func TestBaseGetSuccess(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	rep.KeyPrefix = "some_prefix"
	mock := new(mocks.IBaseRepository)

	keyName, _ := rep.GetKey("fake")
	m := &models.Device{}
	bytes, _ := rep.Encode(m)

	dbHandler.On("Get", keyName).Return(bytes, nil)
	mock.On("Create").Return(m)

	res, err := rep.Get(keyName, mock)
	assert.NoError(t, err)
	assert.Equal(t, m, res)
}

//TestBaseRemoveKeyPrefixEmpty remove if prefix empty
//SUCCESS IF RETURN ERRORS
func TestBaseRemoveKeyPrefixEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = ""

	keyName := "some_key"
	err := rep.Remove(keyName)
	assert.Error(t, err)
}

//TestBaseRemoveKeyNameEmpty remove if key empty
//SUCCESS IF RETURN ERRORS
func TestBaseRemoveKeyNameEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = ""

	keyName := ""
	err := rep.Remove(keyName)
	assert.Error(t, err)
}

//TestBaseRemoveSuccess remove success
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseRemoveSuccess(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	rep.KeyPrefix = "some_prefix"

	keyName, _ := rep.GetKey("some_key")
	dbHandler.On("Remove", keyName).Return(nil)

	err := rep.Remove(keyName)
	assert.NoError(t, err)
}

//TestBaseRemoveSuccess get all with empty prefix
//SUCCESS IF RETURN ERRORS
func TestBaseGetAllPrefixEmpty(t *testing.T) {
	rep := BaseRepository{}
	mock := new(mocks.IBaseRepository)
	_, err := rep.GetAll(mock)
	assert.Error(t, err)
}

//TestBaseGetAllScanError get all with error from db
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseGetAllScanError(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	mock := new(mocks.IBaseRepository)

	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	rep.KeyPrefix = "some_prefix"

	rows := make([]interfaces.Row, 0)
	dbHandler.On("Scan", rep.KeyPrefix).Return(rows, errors.New("some error in db"))

	_, err := rep.GetAll(mock)
	assert.Error(t, err)
}

//TestBaseGetAllSuccess get all success
//SUCCESS IF RETURN WITHOUT ERRORS
func TestBaseGetAllSuccess(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	mock := new(mocks.IBaseRepository)

	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	rep.KeyPrefix = "some_prefix"

	m := &models.Device{
		IP: "192.168.11.4",
		Port: 37777,
	}
	bytes, _ := rep.Encode(m)

	rows := make([]interfaces.Row, 0)
	rows = append(rows, interfaces.Row {
		KeyName: "some_key",
		Value: bytes,
	})
	dbHandler.On("Scan", rep.KeyPrefix).Return(rows, nil)

	mock.On("Create").Return(m)

	res, err := rep.GetAll(mock)
	assert.NoError(t, err)
	assert.True(t, len(res) == 1)
	assert.Equal(t, res[0], m)
}