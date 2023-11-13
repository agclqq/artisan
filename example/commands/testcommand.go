package commands

import (
	"fmt"

	"github.com/agclqq/prowjob"
)

type TestCommand struct {
}

func (t TestCommand) GetCommand() string {
	return "command:test"
}

func (t TestCommand) Usage() string {
	return "test command"
}

func (t TestCommand) Handle(context *prowjob.Context) {
	fmt.Println("this is test command")
}
