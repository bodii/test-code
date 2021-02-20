package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInt_Iterator(t *testing.T) {
	data := ArrayInt{1, 2, 6, 9, 4}
	iterator := data.Iterator()

	i := 0
	for iterator.HasNext() {
		assert.Equal(t, data[i], iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
