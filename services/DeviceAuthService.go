package services

import (
	"errors"
	"fmt"

	"github.com/prohevg/video-security/interfaces"
	"github.com/prohevg/video-security/logwriter"
	"github.com/prohevg/video-security/models"
)

//DeviceAuthService devices in application
type DeviceAuthService struct {
	Logger *logwriter.Logger
	Repository interfaces.IDeviceAuthRepository
}

//AddOrUpdate device
func (service *DeviceAuthService) AddOrUpdate(login string, password string) (models.DeviceAuth, error) {
	if login == "" {
		return models.DeviceAuth{}, fmt.Errorf("login is required")
	}

	if password == "" {
		return models.DeviceAuth{}, fmt.Errorf("password is required")
	}

	devAuth := models.DeviceAuth {
		Login: login,
		Password: password,
	}



	return service.Repository.AddOrUpdate(devAuth)
}

//Get device
func (service *DeviceAuthService) Get(login string) (models.DeviceAuth, error) {
	if login == "" {
		return models.DeviceAuth{}, fmt.Errorf("login is required")
	}

	return service.Repository.Get(login)
}

//Remove device
func (service *DeviceAuthService) Remove(login string) error {
	if login == "" {
		return errors.New("login is required")
	}

	return service.Repository.Remove(login)
}

//GetAll all in db
func (service *DeviceAuthService) GetAll() ([]models.DeviceAuth, error) {
	return service.Repository.GetAll()
}