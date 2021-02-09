package prototype

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKeywords_Clone(t *testing.T) {
	updatedAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			word:      "testA",
			visit:     1,
			UpdatedAt: &updatedAt,
		},
		"testB": &Keyword{
			word:      "testB",
			visit:     2,
			UpdatedAt: &updatedAt,
		},
		"testC": &Keyword{
			word:      "testC",
			visit:     3,
			UpdatedAt: &updatedAt,
		},
	}

	now := time.Now()
	updatedWords := []*Keyword{
		{
			word:      "testB",
			visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)
	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
