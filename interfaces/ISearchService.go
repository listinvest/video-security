package interfaces

import (
	"videoSecurity/models"
)

//ISearchService search device in network
type ISearchService interface {
	//search by parameters in network
	Manual(ips string, ports string) ([]models.Device, error)
	//search all in network
	Auto() ([]models.Device, error)
}