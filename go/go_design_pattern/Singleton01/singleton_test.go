package singleton01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
