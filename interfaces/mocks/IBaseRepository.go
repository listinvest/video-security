package mocks

import (
	"errors"
	"videoSecurity/interfaces"
	"github.com/stretchr/testify/mock"
)

//IBaseRepository repository with basic methods
type IBaseRepository struct {
	mock.Mock
	interfaces.IDbHandler
	KeyPrefix string
}

//GetKeyPrefix prefix on key name
func (_m *IBaseRepository) GetKeyPrefix() (string, error) {
	args := _m.Called()
	return args.String(0), args.Error(1)
}

//GetKey prefix on key name
func (_m *IBaseRepository) GetKey(postfix string) (string, error) {
	args := _m.Called(postfix)
	return args.String(0), args.Error(1)
}

//AddOrUpdate record to db
func (_m *IBaseRepository) AddOrUpdate(keyName string, e interface{}) error {
	args := _m.Called(keyName, e)
	return args.Error(0)
}

//Get record from db
func (_m *IBaseRepository) Get(keyName string, delegate interfaces.ICreateModelDelegate) (interface{}, error) {
	args := _m.Called(keyName, delegate)
	return args.Get(0), args.Error(1)
}

//Remove record into db
func (_m *IBaseRepository) Remove(keyName string) error {
	args := _m.Called(keyName)
	return args.Error(0)
}

//GetAll all records
func (_m *IBaseRepository) GetAll(delegate interfaces.ICreateModelDelegate) ([]interface{}, error) {
	args := _m.Called(delegate)
	return args.Get(0).([]interface{}), args.Error(1)
}

//Encode model to []byte
func (_m *IBaseRepository) Encode(e interface{}) ([]byte, error)  {
	args := _m.Called(e)
	return args.Get(0).([]byte), args.Error(1)
}

//Decode []byte to model
func (_m *IBaseRepository) Decode(data []byte, e interface{}) error {
    if data == nil {
		return errors.New("data is null")
	}
	
	args := _m.Called(data, e)
	return args.Error(0)
}

//Create create model
func (_m *IBaseRepository) Create() interface{} {
	args := _m.Called()
	return args.Get(0)
}
