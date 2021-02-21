package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommandFunc(t *testing.T) {
	eventChan := make(chan string)
	defer close(eventChan)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}
	}()

	commands := make(chan Command, 1000)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command Command
			switch event {
			case "start":
				command = StartCommandFunc()
			case "archive":
				command = ArchiveCommandFunc()
			}

			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}

}
