package mocks

import (
	"videoSecurity/interfaces"
	"github.com/stretchr/testify/mock"
)

//IDbHandler mock handler for db
type IDbHandler struct {
	mock.Mock
}

//Get from db
func (_m *IDbHandler) Get(keyName string) ([]byte, error) {
	args := _m.Called(keyName)
	return args.Get(0).([]byte), args.Error(1)
}

//AddOrUpdate put into db
func (_m *IDbHandler) AddOrUpdate(keyName string, value []byte) error {
	args := _m.Called(keyName, value)
	return args.Error(0)
}

//Remove into db
func (_m *IDbHandler) Remove(keyName string) error {
	args := _m.Called(keyName)
	return args.Error(0)
}

//Scan find records into db by prefix
func (_m *IDbHandler) Scan(prefix string) ([]interfaces.Row, error) {
	args := _m.Called(prefix)
	return args.Get(0).([]interfaces.Row), args.Error(1)
}