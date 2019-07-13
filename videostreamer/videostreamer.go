package videostreamer

import (
	"log"
	"net/http"
	"sync"

	"github.com/prohevg/video-security/videostreamer/rtspstream"
)

//IVideoStreamer streamer
type IVideoStreamer interface {
	Run(rw http.ResponseWriter, r *http.Request, rtsp string, verbose bool) error 
}


//VideoStreamer streamer
type VideoStreamer struct {
	Mutex *sync.RWMutex
	//Dispatcher
	Dispatcher *StreamerDispatcher
}

//Run stream
func (v *VideoStreamer) Run(rw http.ResponseWriter, r *http.Request, rtsp string, verbose bool) error {
	v.Mutex.RLock()

	streamer := v.Dispatcher.Get(rtsp)

	log.Printf("exist for rtsp: %s, streamer is nil: %t", rtsp, streamer == nil)

	if streamer == nil {
		streamer = &rtspstream.RtspStreamer{
			Mutex: &sync.RWMutex{},
			URL:   rtsp,
		}

		streamer.Init()

		err := streamer.Run(rtsp, verbose)
		if err != nil {
			v.Mutex.RUnlock()
			return err
		}
	}

	err := streamer.AddOrUpdateOutput(rw, r)

	v.Mutex.RUnlock()

	return err
}
