package commands

import (
	"fmt"

	"github.com/agclqq/artisan"
)

type TestCommand struct {
}

func (t TestCommand) GetCommand() string {
	return "command:test"
}

func (t TestCommand) Usage() string {
	return "test command"
}

func (t TestCommand) Handle(context *artisan.Context) {
	fmt.Println("test command")
}
