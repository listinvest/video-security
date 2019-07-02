package pipe

//IPipe interface windows\unix pipe
type IPipe interface {
	//create pipe
	Create() error

	//close pipe
	Close() 

	//Accept connection
	Accept() error

	//Read reads data from the connection.
	Read(buf []byte) (n int, err error)

	//address pipe
	GetAddress() string
}