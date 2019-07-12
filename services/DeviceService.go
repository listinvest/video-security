package services

import (
	"regexp"
	"errors"
	"fmt"

	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
	"videoSecurity/models"
)

//DeviceService devices in application
type DeviceService struct {
	Logger *logwriter.Logger
	Repository interfaces.IDeviceRepository
}

//AddOrUpdate device
func (service *DeviceService) AddOrUpdate(ip string, port int) (models.Device, error) {
	if ip == "" {
		return models.Device{}, fmt.Errorf("ip is required")
	}

	if !service.validIP4(ip) {
		return models.Device{}, fmt.Errorf("ip isn't valid")
	}
	

	if port == 0 {
		return models.Device{}, fmt.Errorf("port is required")
	}

	dev := models.Device {
		IP: ip,
		Port: port,
	}

	return service.Repository.AddOrUpdate(dev)
}

//Get device
func (service *DeviceService) Get(ip string) (models.Device, error) {
	if ip == "" {
		return models.Device{}, fmt.Errorf("ip is required")
	}

	return service.Repository.Get(ip)
}

//Remove device
func (service *DeviceService) Remove(ip string) error {
	if ip == "" {
		return errors.New("ip is required")
	}

	return service.Repository.Remove(ip)
}

//GetAll all in db
func (service *DeviceService) GetAll() ([]models.Device, error) {
	return service.Repository.GetAll()
}

//validate ip
func  (service *DeviceService) validIP4(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
			return true
	}
	return false
}