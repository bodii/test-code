package singleton02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
