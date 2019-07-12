package interfaces

import (
	"videoSecurity/models"
)

//IDeviceAuthRepository password for auth on device
type IDeviceAuthRepository interface {
	//AddOrUpdate to db
	AddOrUpdate(models.DeviceAuth) (models.DeviceAuth, error)
	//Remove into db
	Remove(key string) error
	//Get from db
	Get(key string) (models.DeviceAuth, error)
	//GetAll all in db
	GetAll() ([]models.DeviceAuth, error)
}