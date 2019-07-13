package controllers

import (
	"strings"
	"net/http"

	"github.com/prohevg/video-security/videostreamer"
	"github.com/prohevg/video-security/logwriter"

	"github.com/gin-gonic/gin"	
)

//VideoController search device in network
type VideoController struct {
	Logger *logwriter.Logger
	IVideoStreamer videostreamer.IVideoStreamer
}

//Run video stream
func (contrl *VideoController) Run(c *gin.Context) {
	c.Header("Content-Type", "video/mp4")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	url := getUrl(c.Request)
	err := contrl.IVideoStreamer.Run(c.Writer, c.Request, url, false)
	
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

//getUrl url from request
func getUrl(r *http.Request) string {
	defaultURL := "rtsp://admin:1q2w3e4r5t6y@192.168.11.131:554/cam/realmonitor?channel=1&subtype=1"

	urls, ok := r.URL.Query()["url"]

	if !ok {
		//log.Println("Url Param 'key' is missing")
		return defaultURL
	}

	if len(urls[0]) > 0 {
		url := strings.Replace(urls[0], "+", "?", -1)
		return strings.Replace(url, "$", "&", -1)
	}

	return defaultURL
}
