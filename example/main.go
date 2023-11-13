package main

import (
	"fmt"

	"github.com/agclqq/prowjob"
	"github.com/agclqq/prowjob/example/commands"
)

func main() {
	art := prowjob.New()
	art.Add(commands.TestCommand{})
	art.AddFunc("test2", func(ctx *prowjob.Context) {
		fmt.Println("this is func test")
		fmt.Println(ctx.Param)
		for k, v := range ctx.TidyParma {
			fmt.Println(k, v)
		}
	}, "arg1", "arg2")
	art.Run("test2", "arg1", "-arg2", "2", "arg7", "7", "--arg3", "3", "-arg4=4", "--arg5=5", "arg6=6")
	art.Run("command:test")
}
