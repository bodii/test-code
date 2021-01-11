package main

import "errors"

// Undo type
type Undo []func()

// Add ...
func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

// Undo ...
func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil
	}
	*undo = functions[:index]
	return nil
}

// IntSet struct
type IntSet struct {
	data map[int]bool
	undo Undo
}

// NewIntSet create new IntSet
func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

// Undo ...
func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

// Contains ...
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

// Add ...
func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

// Delete ...
func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
