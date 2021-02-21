package command

import "fmt"

// Command func
type Command func() error

// StartCommandFunc start command func
func StartCommandFunc() Command {
	return func() error {
		fmt.Println("game start")
		return nil
	}
}

// ArchiveCommandFunc archive command func
func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("game archive")
		return nil
	}
}
