package artisan

import (
	"fmt"
	"testing"

	commands2 "github.com/agclqq/artisan/example/commands"
)

func TestNew(t *testing.T) {
	art := New()
	art.Add(commands2.TestCommand{})
	art.AddFunc("test2", func(ctx *Context) {
		fmt.Println("this is func test")
		fmt.Println(ctx.Param)
	}, "arg1", "arg2")
	art.Run()
}
