package pipe

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/natefinch/npipe"
)

//WinPipe windows pipe
type WinPipe struct {
	//listener
	Listener *npipe.PipeListener

	//connection
	Conn net.Conn

	//name of pipe (address)
	Address string

	isReady bool
}

//Create pipe
func (pipe *WinPipe) Create() error {
	pipeAddr := pipe.GetAddress()

	ln, err := npipe.Listen(pipeAddr)
	if err != nil {
		return err
	}
	pipe.Listener = ln

	return nil
}

//Accept connection
func (pipe *WinPipe) Accept() error {
	conn, err := pipe.Listener.Accept()
	if err != nil {
		return err
	}

	pipe.Conn = conn
	pipe.isReady =true

	return nil
}

//Read reads data from the connection.
func (pipe *WinPipe) Read(buf []byte) (n int, err error) {
	return pipe.Conn.Read(buf)
}

//Close pipe
func (pipe *WinPipe) Close() {
	pipe.Conn.Close()
	pipe.Listener.Close()
}

//GetAddress get address pipe
func (pipe *WinPipe) GetAddress() string {
	if pipe.Address == "" {
		pipe.Address = fmt.Sprintf(`\\.\pipe\rtspvideo\%d`, pipe.getRandomInt())
	}

	return pipe.Address
}

//getRandomInt
func (pipe *WinPipe) getRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 300000
	return rand.Intn(max-min) + min
}
