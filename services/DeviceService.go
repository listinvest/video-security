package services

import (
	"strings"
	"errors"
	"fmt"
	"regexp"

	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
	"videoSecurity/models"
	"videoSecurity/deviceonvif"
)

//DeviceService devices in application
type DeviceService struct {
	Logger     *logwriter.Logger
	Repository interfaces.IDeviceRepository
	IDeviceAuthRepository interfaces.IDeviceAuthRepository
	IDeviceOnvif deviceonvif.IDeviceOnvif
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

	xAddr := fmt.Sprintf("%s:%v", ip, port)

	dev := models.Device{
		IP:       ip,
		Port:     port,
	}

	allAuth, err := service.IDeviceAuthRepository.GetAll()
	if err != nil {
		return models.Device{}, err
	}

	if len(allAuth) == 0 {
		return models.Device{}, errors.New("not found login/passwoed for devices")
	}

	for _, auth := range allAuth {
		channels, err := service.getChannels(xAddr, auth.Login, auth.Password)
		if err == nil {
			dev.Login = auth.Login
			dev.Password = auth.Password
			dev.Channels = channels
			break
		}
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
func (service *DeviceService) validIP4(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
		return true
	}
	return false
}

//getChannels channel for device
func (service *DeviceService) getChannels(xaddr string, login string, password string) ([]models.Channel, error) {
	channels := make([]models.Channel, 0)

	deviceOnf := service.IDeviceOnvif.NewDevice(xaddr, login, password)
	profiles := deviceOnf.GetProfiles()

	if len(profiles) == 0 {
		return channels, fmt.Errorf("%s has't profiles", xaddr)
	}

	for _, p := range profiles {
		streams := deviceOnf.GetStreamUri(p.Profiles.Token)
		for _, s := range streams {
			ch := models.Channel {
				Name: string(p.Profiles.Name),
				Rtsp: strings.Replace(string(s.MediaUri.Uri), "rtsp://", fmt.Sprintf("rtsp://%s:%s@", login, password), 1),
			}
			channels = append(channels, ch)
		}
	}	

	return channels, nil
}
