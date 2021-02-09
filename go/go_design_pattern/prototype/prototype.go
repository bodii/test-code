package prototype

import (
	"encoding/json"
	"time"
)

// Keyword search key word
type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// Clone this is use json and unjson deep clone
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword

	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// Keywords key word map
type Keywords map[string]*Keyword

// Clone new a keywords
func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}

	for _, word := range updatedWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
