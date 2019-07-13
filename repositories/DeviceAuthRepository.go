package repositories

import (
	"github.com/prohevg/video-security/interfaces"
	"github.com/prohevg/video-security/models"
)

//DeviceAuthRepository repository for devices
type DeviceAuthRepository struct {
	interfaces.IBaseRepository
}

//AddOrUpdate add to db
func (rep *DeviceAuthRepository) AddOrUpdate(m models.DeviceAuth) (models.DeviceAuth, error) {
	err := rep.IBaseRepository.AddOrUpdate(m.Login, &m)
	return m, err
}

//Get get from db
func (rep *DeviceAuthRepository) Get(key string) (models.DeviceAuth, error) {
	get, err := rep.IBaseRepository.Get(key, rep)
	if get == nil {
		return models.DeviceAuth{}, err
	}
	res := get.(*models.DeviceAuth)
	return *res, err
}

//Remove remove into db
func (rep *DeviceAuthRepository) Remove(key string) error {
	return rep.IBaseRepository.Remove(key)
}

//GetAll all
func (rep *DeviceAuthRepository) GetAll() ([]models.DeviceAuth, error) {
	all, err := rep.IBaseRepository.GetAll(rep)
	return rep.getArray(all), err
}

//Create create model
func (rep *DeviceAuthRepository) Create() interface{} {
	return &models.DeviceAuth{}
}

//getArray convert array interface{} to array models
func (rep *DeviceAuthRepository) getArray(arr []interface{}) []models.DeviceAuth {
	res := make([]models.DeviceAuth, len(arr))
	for index, i := range arr {
		if i != nil {
			v := i.(*models.DeviceAuth)
			res[index] = *v
		}
	}
	return res
}