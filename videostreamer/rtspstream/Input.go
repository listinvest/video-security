package rtspstream

import (
	"fmt"
	"log"
	"sync"
	"unsafe"
)

// #include "videostreamer.h"
// #include <stdlib.h>
// #cgo LDFLAGS: -lavformat -lavdevice -lavcodec -lavutil
// #cgo CFLAGS: -std=c11
// #cgo pkg-config: libavcodec
import "C"

//Input video context
type Input struct {
	Mutex      *sync.RWMutex
	Format     string
	URL        string
	Verbose    bool
	vsInput    *C.struct_VSInput
	PacketChan chan *C.AVPacket
	CanSend    bool
}

//Open video context
func (input *Input) Open() error {
	inputFormatC := C.CString(input.Format)
	inputURLC := C.CString(input.URL)
	VerboseC := C.bool(input.Verbose)

	input.vsInput = C.vs_open_input(inputFormatC, inputURLC, VerboseC)

	C.free(unsafe.Pointer(inputFormatC))
	C.free(unsafe.Pointer(inputURLC))

	if input.vsInput == nil {
		return fmt.Errorf("didn't create input context stream")
	}

	return nil
}

//Run video context
func (input *Input) Run() {
	for {
		// Read a packet.
		var pkt C.AVPacket
		readRes := C.int(0)
		readRes = C.vs_read_packet(input.vsInput, &pkt, C.bool(input.Verbose))

		if readRes == -1 {
			log.Printf("encoder: Failure reading packet")
			break
		}

		if readRes == 0 || !input.CanSend {
			continue
		}

		pktCopy := C.av_packet_clone(&pkt)
		if pktCopy == nil {
			log.Printf("Unable to clone packet")
			continue
		}

		input.PacketChan <- pktCopy

		C.av_packet_unref(&pkt)
	}
}

//Dispose video context
func (input *Input) Dispose() {
	input.Mutex.Lock()
	defer input.Mutex.Unlock()

	if input.vsInput != nil {
		C.vs_destroy_input(input.vsInput)
		input.vsInput = nil
	}
}
