package models

//DeviceAuth login/password for auth on device 
type DeviceAuth struct {
	Login string `json:"login"`
	Password string `json:"password"`
}