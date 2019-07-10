package interfaces

//ICreateModelDelegate delegate to create a model
type ICreateModelDelegate interface {
    Create() interface{}
}

//IBaseRepository repository with basic methods
type IBaseRepository interface {
	//prefix on key name
	GetKeyPrefix() (string, error)
	//prefix on key name
	GetKey(postfix string) (string, error)
	//AddOrUpdate record to db
	AddOrUpdate(keyName string, e interface{}) error
	//Get record from db
	Get(keyName string, delegate ICreateModelDelegate) (interface{}, error)
	//Remove record into db
	Remove(keyName string) error
	//GetAll all records
	GetAll(delegate ICreateModelDelegate) ([]interface{}, error)
	//Encode model to []byte
	Encode(e interface{}) ([]byte, error) 
	//Decode []byte to model
	Decode(data []byte, e interface{}) error
}