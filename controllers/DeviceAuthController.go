package controllers

import (
	"net/http"
	
	"videoSecurity/interfaces"
	"videoSecurity/logwriter"

	"github.com/gin-gonic/gin"
)

//DeviceAuthController login/password for auth on device
type DeviceAuthController struct {
	Logger *logwriter.Logger
	interfaces.IDeviceAuthService
}

//All login/password 
func (contrl *DeviceAuthController) All(c *gin.Context) {
	all, err := contrl.IDeviceAuthService.GetAll()

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, all)
}

//AddOrUpdate add login/password 
func (contrl *DeviceAuthController) AddOrUpdate(c *gin.Context) {
	login  := c.Param("login")
	password  := c.Param("password")

	all, err := contrl.IDeviceAuthService.AddOrUpdate(login, password)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, all)
}

//Remove login/password 
func (contrl *DeviceAuthController) Remove(c *gin.Context) {
	login  := c.Param("login")

	err := contrl.IDeviceAuthService.Remove(login)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}