package videostreamer

import (
	"net/http"
)

//IStreamer streamer video
type IStreamer interface {
	//GetKey key of streamer
	GetKey() string

	//Init initialization
	Init() 

	//Run input video context
	Run(url string, verbose bool) error

	//Run output video context
	AddOrUpdateOutput(rw http.ResponseWriter, r *http.Request) error
}