package observer

import "fmt"

// IObserver interface
type IObserver interface {
	Update(msg string)
}

// ISubject interface
type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(observer IObserver)
}

// Subject subject struct
type Subject struct {
	observers []IObserver
}

// Register register function
func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

// Remove remove function
func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

// Notify notify function
func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

// Observer1 struct
type Observer1 struct{}

// Update Observer1 update function
func (Observer1) Update(msg string) {
	fmt.Printf("Obsever1: %s", msg)
}

// Observer2 struct
type Observer2 struct{}

// Update Obsever2 update function
func (Observer2) Update(msg string) {
	fmt.Printf("Obsever2: %s", msg)
}
