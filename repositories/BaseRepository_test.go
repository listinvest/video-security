package repositories

import (
	"errors"
	"videoSecurity/interfaces"
	"videoSecurity/models"
	"testing"

	"videoSecurity/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

//Test_BaseRepository_Get_KeyPrefixEmpty get empty prefix
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_Get_KeyPrefixEmpty(t *testing.T) {
	rep := BaseRepository{}
	_, err := rep.GetKeyPrefix()
	assert.Error(t, err)
}

//Test_BaseRepository_KeyPrefixNotEmpty get isn't empty prefix
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_Get_KeyPrefixNotEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = "some_prefix"
	real, err := rep.GetKeyPrefix()
	assert.NoError(t, err)
	assert.True(t, real != "")
	assert.Equal(t, real, rep.KeyPrefix)
}

//Test_BaseRepository_Get_KeyEmpty get isn't empty key
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_Get_KeyEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = "some_prefix"
	_, err := rep.GetKey("")
	assert.Error(t, err)
}

//Test_BaseRepository_Get_KeyNotEmpty get isn't empty key
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_Get_KeyNotEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = "some_prefix"
	keyName := "some_key"
	expected := rep.KeyPrefix + keyName
	real, err := rep.GetKey(keyName)
	assert.NoError(t, err)
	assert.True(t, real != "")
	assert.Equal(t, real, expected)
}

//Test_BaseRepository_EncodeError encode error
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_EncodeError(t *testing.T) {
	rep := BaseRepository{}
	_, err := rep.Encode(nil)
	assert.Error(t, err)
}

//Test_BaseRepository_EncodeSuccess encode success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_EncodeSuccess(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	_, err := rep.Encode(m)
	assert.NoError(t, err)
}

//Test_BaseRepository_DecodeDataError decode error
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_DecodeDataError(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	err := rep.Decode(nil, &m)
	assert.Error(t, err)
}

//Test_BaseRepository_DecodeModelError decode error
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_DecodeModelError(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	bytes, err := rep.Encode(m)
	assert.NoError(t, err)

	err = rep.Decode(bytes, nil)
	assert.Error(t, err)
}

//Test_BaseRepository_DecodeSuccess encode success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_DecodeSuccess(t *testing.T) {
	rep := BaseRepository{}
	m := models.Device{}
	bytes, err := rep.Encode(m)
	assert.NoError(t, err)

	err = rep.Decode(bytes, &m)
	assert.NoError(t, err)
}

//Test_BaseRepository_AddOrUpdate_KeyEmpty save with empty key
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_AddOrUpdate_KeyEmpty(t *testing.T) {
	rep := BaseRepository{}

	var e interface{}
	keyName := ""

	err := rep.AddOrUpdate(keyName, e)
	assert.Error(t, err)
}

//Test_BaseRepository_AddOrUpdate_Success save success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_AddOrUpdate_Success(t *testing.T) {
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

//Test_BaseRepository_Get_ModelKeyEmpty get with empty key
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_Get_ModelKeyEmpty(t *testing.T) {
	rep := BaseRepository{}
	mock := new(mocks.IBaseRepository)

	keyName := ""

	_, err := rep.Get(keyName, mock)
	assert.Error(t, err)
}

//Test_BaseRepository_Get_KeyFake get with fake key
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_Get_KeyFake(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	mock := new(mocks.IBaseRepository)

	keyName := "fake"
	dbHandler.On("Get", keyName, rep).Return(nil, nil)

	_, err := rep.Get(keyName, mock)
	assert.Error(t, err)
}

//Test_BaseRepository_Get_Success get with empty key
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_Get_Success(t *testing.T) {
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

//Test_BaseRepository_Remove_KeyPrefixEmpty remove if prefix empty
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_Remove_KeyPrefixEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = ""

	keyName := "some_key"
	err := rep.Remove(keyName)
	assert.Error(t, err)
}

//Test_BaseRepository_Remove_KeyNameEmpty remove if key empty
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_Remove_KeyNameEmpty(t *testing.T) {
	rep := BaseRepository{}
	rep.KeyPrefix = ""

	keyName := ""
	err := rep.Remove(keyName)
	assert.Error(t, err)
}

//Test_BaseRepository_Remove_Success remove success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_Remove_Success(t *testing.T) {
	dbHandler := new(mocks.IDbHandler)
	rep := BaseRepository{}
	rep.IDbHandler = dbHandler
	rep.KeyPrefix = "some_prefix"

	keyName, _ := rep.GetKey("some_key")
	dbHandler.On("Remove", keyName).Return(nil)

	err := rep.Remove(keyName)
	assert.NoError(t, err)
}

//Test_BaseRepository_Remove_Success get all with empty prefix
//SUCCESS IF RETURN ERRORS
func Test_BaseRepository_GetAll_PrefixEmpty(t *testing.T) {
	rep := BaseRepository{}
	mock := new(mocks.IBaseRepository)
	_, err := rep.GetAll(mock)
	assert.Error(t, err)
}

//Test_BaseRepository_GetAll_ScanError get all with error from db
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_GetAll_ScanError(t *testing.T) {
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

//Test_BaseRepository_GetAll_Success get all success
//SUCCESS IF RETURN WITHOUT ERRORS
func Test_BaseRepository_GetAll_Success(t *testing.T) {
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