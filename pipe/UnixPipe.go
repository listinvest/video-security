package pipe

import (
	"fmt"
	"os"
)

//UnixPipe unix pipe
type UnixPipe struct {
	// writer
	wPipe *os.File

	// reader
	rPipe *os.File

	//name of pipe (address)
	Address string
}

//Create pipe
func (pipe *UnixPipe) Create() error {
	rPipe, wPipe, err := os.Pipe()
	if err != nil {
		return err
	}

	pipe.rPipe = rPipe
	pipe.wPipe = wPipe

	return nil
}	

//Accept connection
func (pipe *UnixPipe) Accept() error { 
	return nil
}

//Read reads data from the connection.
func (pipe *UnixPipe) Read(buf []byte) (n int, err error) {
	return pipe.rPipe.Read(buf)
}

//Close pipe
func (pipe *UnixPipe) Close() {
	pipe.rPipe.Close()
	pipe.wPipe.Close()
}

//GetAddress get address pipe
func (pipe *UnixPipe) GetAddress() string {
	if (pipe.Address == "") {
		pipe.Address = fmt.Sprintf("pipe:%d", pipe.wPipe.Fd())
	}

	return pipe.Address
}