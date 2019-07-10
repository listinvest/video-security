package rtspstream

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"videoSecurity/videostreamer/pipe"
)

// #include "videostreamer.h"
// #include <stdlib.h>
// #cgo LDFLAGS: -lavformat -lavdevice -lavcodec -lavutil
// #cgo CFLAGS: -std=c11
// #cgo pkg-config: libavcodec
import "C"

//RtspStreamer streamer for rtsp
type RtspStreamer struct {
	Mutex *sync.RWMutex

	//rtsp url
	URL string

	// Media input context.
	Input Input

	//Outputs
	Outputs []Output
}

//GetKey key of streamer
func (streamer *RtspStreamer) GetKey() string {
	return streamer.URL
}

//Init initialization
func (streamer *RtspStreamer) Init() {
	C.vs_setup()
}

//Run stream
func (streamer *RtspStreamer) Run(rtsp string, verbose bool) error {
	packetChan := make(chan *C.AVPacket, 100)

	streamer.Input = Input{
		Mutex:      streamer.Mutex,
		Format:     "rtsp",
		URL:        rtsp,
		Verbose:    verbose,
		PacketChan: packetChan,
	}

	err := streamer.Input.Open()
	if err != nil {
		return fmt.Errorf("Unable open video context: %s", err)
	}

	go streamer.Input.Run()
	go streamer.readPackets()

	return nil
}

//AddOrUpdateOutput new output
func (streamer *RtspStreamer) AddOrUpdateOutput(rw http.ResponseWriter, r *http.Request) error {
	log.Printf("%s: Client run", r.RemoteAddr)

	streamer.removeOutput(r.RemoteAddr, false)

	packetChan := make(chan *C.AVPacket, 32)
	winpipe := pipe.WinPipe{}
	output := Output{
		RemoteAddr: r.RemoteAddr,
		Format:     "mp4",
		URL:        winpipe.GetAddress(),
		Verbose:    streamer.Input.Verbose,
		Mutex:      streamer.Mutex,
		Input:      streamer.Input,
		Pipe:       &winpipe,
		PacketChan: packetChan,
	}

	err := winpipe.Create()
	if err != nil {
		return fmt.Errorf("Unable create pipe: %s", err)
	}

	go streamer.runOutput(output)

	streamer.Outputs = append(streamer.Outputs, output)
	streamer.Input.CanSend = true

	output.Write(rw)
	log.Printf("%s: output cleaned up", output.RemoteAddr)
	output.Dispose()

	streamer.removeOutput(r.RemoteAddr, true)

	return nil
}

//readPackets reads packets from source video context
func (streamer *RtspStreamer) readPackets() {
	for origPkt := range streamer.Input.PacketChan {
		for _, output := range streamer.Outputs {
			output.PacketChan <- origPkt
		}
	}
}

//readPackets create output video context
func (streamer *RtspStreamer) runOutput(o Output) {
	err := o.Open()
	if err != nil {
		o.Dispose()
		streamer.removeOutput(o.RemoteAddr, true)
	}

	o.ReceivePackage()
}

//removeOutput close output stream
func (streamer *RtspStreamer) removeOutput(remoteAddr string, needCloseInput bool) {
	index := streamer.indexOf(remoteAddr)
	if index > -1 {
		o := streamer.Outputs[index]
		o.Dispose()
		streamer.removeByIndex(index)
	}

	if needCloseInput {
		streamer.removeInput(remoteAddr)
	}
}

//removeInput close input stream
func (streamer *RtspStreamer) removeInput(remoteAddr string) {
	if len(streamer.Outputs) == 0 {
		log.Printf("stop stream: %s", streamer.Input.URL)
		streamer.Input.CanSend = false
		streamer.Input.Dispose()
	}
}

//indexOf
func (streamer *RtspStreamer) indexOf(key string) int {
	if streamer.Outputs == nil {
		return -1
	}

	for i, o := range streamer.Outputs {
		if o.RemoteAddr == key {
			return i
		}
	}
	return -1
}

//removeByIndex remove  from outputs
func (streamer *RtspStreamer) removeByIndex(i int) []Output {
	streamer.Outputs[len(streamer.Outputs)-1], streamer.Outputs[i] = streamer.Outputs[i], streamer.Outputs[len(streamer.Outputs)-1]
	return streamer.Outputs[:len(streamer.Outputs)-1]
}
