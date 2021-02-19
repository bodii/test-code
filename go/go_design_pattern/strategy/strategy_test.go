package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getData() ([]byte, bool) {
	return []byte("test data"), false
}

func Test_demo(t *testing.T) {
	data, sensitive := getData()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}

	storage, err := NewStorageStrategy(strategyType)
	assert.NoError(t, err)
	assert.NoError(t, storage.Save("./test.txt", data))
}
