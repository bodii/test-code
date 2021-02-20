package iterator

// Iterator iterator interface
type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

// ArrayInt array int
type ArrayInt []int

// ArrayIntIterator ArrayInt iterator
type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

// HasNext ArrayIntIterator has next function
func (aIter *ArrayIntIterator) HasNext() bool {
	return aIter.index < len(aIter.arrayInt)-1
}

// Next ArrayIntIterator next index function
func (aIter *ArrayIntIterator) Next() {
	aIter.index++
}

// CurrentItem ArrayIntIterator get current item function
func (aIter *ArrayIntIterator) CurrentItem() interface{} {
	return aIter.arrayInt[aIter.index]
}

// Iterator function
func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}
