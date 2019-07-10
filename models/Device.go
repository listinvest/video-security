package models

//Device camera 
type Device struct {
	IP string `json:"ip"`
	Port int `json:"port"`
	Xaddr string `json:"xaddr"`
	Login string `json:"login"`
	Password string `json:"password"`
	Channels []Channel `json:"channels"`
}