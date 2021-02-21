package command

import "fmt"

// ICommand command interface
type ICommand interface {
	Execute() error
}

// StartCommand struct
type StartCommand struct{}

// NewStartCommand create an StartCommand
func NewStartCommand( /* params */ ) *StartCommand {
	return &StartCommand{}
}

// Execute execute StartCommand function
func (s *StartCommand) Execute() error {
	fmt.Println("game start")
	return nil
}

// ArchiveCommand struct
type ArchiveCommand struct{}

// NewArchiveCommand create an ArchiveCommand
func NewArchiveCommand( /* params */ ) *ArchiveCommand {
	return &ArchiveCommand{}
}

// Execute ArchiveCommand execute function
func (a *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}
