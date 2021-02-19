package strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

// StorageStrategy storage strategy
type StorageStrategy interface {
	Save(name string, data []byte) error
}

// fileStorage file storage
type fileStorage struct{}

// Save save file
func (f *fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

// encryptFileStorage encrypt file storage
type encryptFileStorage struct{}

// encrypt data algorithm
func encrypt(data []byte) ([]byte, error) {
	return data, nil
}

// Save encrypt data and save file
func (ef *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

// NewStorageStrategy create an new StorageStrategy
func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}
