package repositories

import (
	"bytes"
	"encoding/gob"
	"strings"
	"errors"
	"fmt"
	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
)

//BaseRepository repository for FolderReceive model
type BaseRepository struct {
	interfaces.IDbHandler
	KeyPrefix string
	Logger *logwriter.Logger
}

//CreateModelDelegate delegate to create a model
type CreateModelDelegate struct {
}

//GetKeyPrefix prefix on key name
func (rep *BaseRepository) GetKeyPrefix() (string, error) {
	if rep.KeyPrefix == "" {
		return "", errors.New("prefix not found for repository")
	}

	return rep.KeyPrefix, nil;	
}

//GetKey prefix on key name
func (rep *BaseRepository) GetKey(postfix string) (string, error) {
	if rep.KeyPrefix == "" {
		return "", errors.New("prefix not found for repository")
	}

	if postfix == "" {
		return "", errors.New("postfix not found for repository")
	}

	if (strings.HasPrefix(postfix, rep.KeyPrefix)) {
		return postfix, nil
	}

	return rep.KeyPrefix + postfix, nil;	
}

//AddOrUpdate record to db
func (rep *BaseRepository) AddOrUpdate(keyName string, e interface{}) error {
	keyName, err := rep.GetKey(keyName)
	if err != nil {
		return err
	}

	value, err := rep.Encode(e)
	if err != nil {
		return err
	}

	return rep.IDbHandler.AddOrUpdate(keyName, value)
}

//Get record from db
func (rep *BaseRepository) Get(keyName string, delegate interfaces.ICreateModelDelegate) (interface{}, error) {
	keyName, err := rep.GetKey(keyName)
	if err != nil {
		return nil, err
	}

	e := delegate.Create()
	data, err := rep.IDbHandler.Get(keyName)
	if err != nil {
		return e, err
	}

	if data == nil || len(data) == 0 {
		return e, fmt.Errorf("not found record with key %s", keyName)
	}
	
	err = rep.Decode(data, e)
	if err != nil {
		return e, err
	}

	return e, err
}

//Remove record into db
func (rep *BaseRepository) Remove(keyName string) error {
	keyName, err := rep.GetKey(keyName)
	if err != nil {
		return err
	}

	return rep.IDbHandler.Remove(keyName)
}

//GetAll all records
func (rep *BaseRepository) GetAll(delegate interfaces.ICreateModelDelegate) ([]interface{}, error) {
	prefix, err := rep.GetKeyPrefix()
	if err != nil {
		return nil, err
	}

	rows, err := rep.IDbHandler.Scan(prefix)
	if err != nil {
		return nil, err
	}

	res := make([]interface{}, 0)
	for _, row := range rows {
		e := delegate.Create()
		err := rep.Decode(row.Value, e)
		if err == nil {
			res = append(res, e)
		} else {
			rep.Logger.Error("didn't decode: %v", err)
		}
	}

	return res, nil
}

//Encode model to []byte
func (rep *BaseRepository) Encode(e interface{}) ([]byte, error) {
	if e == nil {
		return nil, errors.New("model is nil")
	}

	encBuf := new(bytes.Buffer)
	err := gob.NewEncoder(encBuf).Encode(e)
	if err != nil {
		return nil, err
	}
	return encBuf.Bytes(), nil
}

//Decode []byte to model
func  (rep *BaseRepository) Decode(data []byte, e interface{}) error {
	if data == nil {
		return errors.New("data is nil")
	}

	if e == nil {
		return errors.New("model is nil")
	}

	decBuf := bytes.NewBuffer(data)
	return gob.NewDecoder(decBuf).Decode(e)
}