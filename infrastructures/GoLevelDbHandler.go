package infrastructures

import (
	"errors"

	"videoSecurity/interfaces"
	"videoSecurity/logwriter"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//GoLevelDbHandler Handler for GoLevelDb
type GoLevelDbHandler struct {
	Db     *leveldb.DB
	Logger *logwriter.Logger
}

//Get find value by key
func (handler *GoLevelDbHandler) Get(keyName string) (value []byte, err error) {
	if keyName == "" {
		return nil, errors.New("key is required")
	}

	key := []byte(keyName)
	return handler.Db.Get(key, nil)
}

//AddOrUpdate add or update value by key
func (handler *GoLevelDbHandler) AddOrUpdate(keyName string, value []byte) (err error) {
	if keyName == "" {
		return errors.New("key is required")
	}

	key := []byte(keyName)
	err = handler.Db.Put(key, value, nil)
	if err != nil {
		return err
	}

	return
}

//Remove remove into db
func (handler *GoLevelDbHandler) Remove(keyName string) (err error) {
	if keyName == "" {
		return errors.New("key is required")
	}

	key := []byte(keyName)
	err = handler.Db.Delete(key, nil)
	if err != nil {
		return err
	}

	return
}

//Scan find records into db by prefix
func (handler *GoLevelDbHandler) Scan(prefix string) (rows []interfaces.Row, err error) {
	if prefix == "" {
		return nil, errors.New("prefix is required")
	}

	keyPrefix := util.BytesPrefix([]byte(prefix))
	iter := handler.Db.NewIterator(keyPrefix, nil)
	for iter.Next() {

		key := iter.Key()
		val := iter.Value()

		row := interfaces.Row{
			KeyName: string(key),
			Value:   make([]byte, len(val)),
		}

		copy(row.Value, val)

		rows = append(rows, row)
	}

	iter.Release()
	err = iter.Error()
	return rows, err
}