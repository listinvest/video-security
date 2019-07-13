package interfaces

import (
	"github.com/prohevg/video-security/models"
)

//IDeviceAuthService password for auth on device
type IDeviceAuthService interface {
	//AddOrUpdate password
	AddOrUpdate(login string, password string) (models.DeviceAuth, error) 
	//Remove password
	Remove(login string) error
	//GetAll all in db
	GetAll() ([]models.DeviceAuth, error)
}