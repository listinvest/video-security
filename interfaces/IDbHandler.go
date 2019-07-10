package interfaces

//IDbHandler handler for db
type IDbHandler interface {
	//Get from db
	Get(keyName string) (value []byte, err error)
	//put into db
	AddOrUpdate(keyName string, value []byte) error
	//remove into db
	Remove(keyName string) error
	//Scan find records into db by prefix
	Scan(prefix string) (rows []Row, err error)
}

//Row data in db
type Row struct {
	KeyName string
	Value   []byte
}