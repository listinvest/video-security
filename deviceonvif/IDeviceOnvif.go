package deviceonvif

import (
	"github.com/yakovlevdmv/goonvif/Media"
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
)

//IDeviceOnvif onvif
type IDeviceOnvif interface {
	NewDevice(address string, login string, password string) *DeviceOnvif
	GetProfiles() []Media.GetProfilesResponse
	GetStreamUri(token onvif.ReferenceToken) []Media.GetStreamUriResponse 
	GetSnapshotUri(token onvif.ReferenceToken) []Media.GetSnapshotUriResponse
}