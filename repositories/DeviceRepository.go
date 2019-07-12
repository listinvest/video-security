package repositories

import (
	"videoSecurity/interfaces"
	"videoSecurity/models"
)

//DeviceRepository repository for devices
type DeviceRepository struct {
	interfaces.IBaseRepository
}


//AddOrUpdate add to db
func (rep *DeviceRepository) AddOrUpdate(m models.Device) (models.Device, error) {
	err := rep.IBaseRepository.AddOrUpdate(m.IP, &m)
	return m, err
}

//Get get from db
func (rep *DeviceRepository) Get(ip string) (models.Device, error) {
	get, err := rep.IBaseRepository.Get(ip, rep)
	if get == nil {
		return models.Device{}, err
	}
	res := get.(*models.Device)
	return *res, err
}

//Remove remove into db
func (rep *DeviceRepository) Remove(ip string) error {
	return rep.IBaseRepository.Remove(ip)
}

//GetAll all
func (rep *DeviceRepository) GetAll() ([]models.Device, error) {
	all, err := rep.IBaseRepository.GetAll(rep)
	return rep.getArray(all), err
}

//Create create model
func (rep *DeviceRepository) Create() interface{} {
	return &models.Device{}
}

//getArray convert array interface{} to array models
func (rep *DeviceRepository) getArray(arr []interface{}) []models.Device {
	res := make([]models.Device, len(arr))
	for index, i := range arr {
		if i != nil {
			v := i.(*models.Device)
			res[index] = *v
		}
	}
	return res
}