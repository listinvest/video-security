package interfaces

import (
	"github.com/prohevg/video-security/models"
)

//IDeviceRepository devices in db
type IDeviceRepository interface {
	//AddOrUpdate to db
	AddOrUpdate(models.Device) (models.Device, error)
	//Remove into db
	Remove(ip string) error
	//Get from db
	Get(ip string) (models.Device, error)
	//GetAll all in db
	GetAll() ([]models.Device, error)
}