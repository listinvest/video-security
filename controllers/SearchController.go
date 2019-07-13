package controllers

import (
	"net/http"
	
	"github.com/prohevg/video-security/interfaces"
	"github.com/prohevg/video-security/logwriter"

	"github.com/gin-gonic/gin"
)

//SearchController search device in network
type SearchController struct {
	Logger *logwriter.Logger
	interfaces.ISearchService
}

//Manual search by parameters in network
func (contrl *SearchController) Manual(c *gin.Context) {

	ips := c.Query("ips")
	ports := c.Query("ports")

	all, err := contrl.ISearchService.Manual(ips, ports)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, all)
}

//Auto search all in network
func (contrl *SearchController) Auto(c *gin.Context) {
	all, err := contrl.ISearchService.Auto()

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, all)
}