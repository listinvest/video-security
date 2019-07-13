package controllers

import (
	"strconv"
	"net/http"
	
	"github.com/prohevg/video-security/interfaces"
	"github.com/prohevg/video-security/logwriter"

	"github.com/gin-gonic/gin"
)

//DeviceController device in network
type DeviceController struct {
	Logger *logwriter.Logger
	interfaces.IDeviceService
}

//All devices
func (contrl *DeviceController) All(c *gin.Context) {
	all, err := contrl.IDeviceService.GetAll()

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, all)
}

//AddOrUpdate add device
func (contrl *DeviceController) AddOrUpdate(c *gin.Context) {

	ip := c.Param("ip")
	port := c.Param("port")

	p, _ := strconv.Atoi(port)

	all, err := contrl.IDeviceService.AddOrUpdate(ip, p)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, all)
}

//Remove device
func (contrl *DeviceController) Remove(c *gin.Context) {
	ip := c.Param("ip")

	err := contrl.IDeviceService.Remove(ip)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}