package prowjob

import (
	"strings"
	"testing"
)

type Msg struct {
	command string
	msg     string
	params  string
}

var testMsg Msg

type testCommand struct {
}

func (t testCommand) GetCommand() string {
	return "command:test"
}

func (t testCommand) Usage() string {
	return `command:test`
}

func (t testCommand) Handle(context *Context) {
	testMsg = Msg{command: t.GetCommand(), msg: "test command", params: strings.Join(context.Param, " ")}
}

func TestNew(t *testing.T) {
	art := New()
	art.Add(testCommand{})
	art.AddFunc("command:testFunc", func(ctx *Context) {
		testMsg = Msg{command: "command:testFunc", msg: "test function", params: strings.Join(ctx.Param, " ")}
	})
	tests := []struct {
		name string
		fun  func()
		want Msg
	}{
		{name: "test1", fun: func() {
			art.Run("command:test", "arg1", "-arg2", "v2", "--arg3", "v3", "-arg4")
		}, want: Msg{command: "command:test", msg: "test command", params: "arg1 -arg2 v2 --arg3 v3 -arg4"}},
		{name: "test2", fun: func() {
			art.Run("command:testFunc", "arg1", "v1", "-arg2", "v2", "--arg3", "v3", "-arg4")
		}, want: Msg{command: "command:testFunc", msg: "test function", params: "arg1 v1 -arg2 v2 --arg3 v3 -arg4"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fun()
			if testMsg != tt.want {
				t.Errorf("testMsg = %v, want %v", testMsg, tt.want)
			}
		})
	}
}
