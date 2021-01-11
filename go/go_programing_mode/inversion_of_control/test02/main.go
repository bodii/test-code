package main

import "errors"

// IntSet struct
type IntSet struct {
	data map[int]bool
}

// NewIntSet ...
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

// Add add item
func (set *IntSet) Add(x int) {
	set.data[x] = true
}

// Delete delete item
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

// Contains find item contains
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

// UndoableIntSet struct
type UndoableIntSet struct {
	IntSet
	functions []func()
}

// NewUndoableIntSet new undoable IntSet
func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

// Add ...
func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

// Delete ...
func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

// Undo function
func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil
	}
	set.functions = set.functions[:index]
	return nil
}
