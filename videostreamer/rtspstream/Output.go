package rtspstream

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"unsafe"

	"../../pipe"
)

// #include "videostreamer.h"
// #include <stdlib.h>
// #cgo LDFLAGS: -lavformat -lavdevice -lavcodec -lavutil
// #cgo CFLAGS: -std=c11
// #cgo pkg-config: libavcodec
import "C"

//Output video context
type Output struct {
	RemoteAddr string
	Format     string
	URL        string
	Verbose    bool
	Mutex      *sync.RWMutex
	Input      Input
	vsOutput   *C.struct_VSOutput
	Pipe       pipe.IPipe
	PacketChan chan *C.AVPacket
}

//Open output video context
func (o *Output) Open() error {
	outputFormatC := C.CString(o.Format)
	outputURLC := C.CString(o.URL)

	o.Mutex.RLock()
	o.vsOutput = C.vs_open_output(outputFormatC, outputURLC, o.Input.vsInput, C.bool(o.Verbose))
	o.Mutex.RUnlock()

	C.free(unsafe.Pointer(outputFormatC))
	C.free(unsafe.Pointer(outputURLC))

	if o.vsOutput == nil {
		return fmt.Errorf("Unable to open output")
	}

	return nil
}

//ReceivePackage write packets in output
func (o *Output) ReceivePackage() {
	for pkt := range o.PacketChan {
		writeRes := C.int(0)
		o.Mutex.RLock()
		writeRes = C.vs_write_packet(o.Input.vsInput, o.vsOutput, pkt, C.bool(o.Verbose))
		o.Mutex.RUnlock()
		if writeRes == -1 {
			log.Printf("Failure writing packet")
			C.av_packet_free(&pkt)
			continue
		}
		C.av_packet_free(&pkt)
	}
}

//Run write packets in output
func (o *Output) Write(rw http.ResponseWriter) error {
	o.Pipe.Accept()

	for {
		buf := make([]byte, 1024*5)

		readSize, err := o.Pipe.Read(buf)
		if err != nil {
			return fmt.Errorf("Read error: %s", err)
		}

		if readSize == 0 {
			return fmt.Errorf("EOF")
		}

		_, err = rw.Write(buf[:readSize])
		if err != nil {
			return fmt.Errorf("Write error: %s", err)
		}

		if flusher, ok := rw.(http.Flusher); ok {
			flusher.Flush()
		}
	}
}

//Dispose video context
func (o *Output) Dispose() {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()

	if o.Pipe != nil {
		o.Pipe.Close()
	}

	if o.vsOutput != nil {
		C.vs_destroy_output(o.vsOutput)
		o.vsOutput = nil
	}
}
