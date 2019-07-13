package interfaces

import (
	"github.com/prohevg/video-security/models"
)

//IDeviceService service for devices
type IDeviceService interface {
	//AddOrUpdate device
	AddOrUpdate(ip string, port int) (models.Device, error) 
	//Get device
	Get(ip string) (models.Device, error) 
	//Remove device
	Remove(ip string) error
	//GetAll all in db
	GetAll() ([]models.Device, error)
}