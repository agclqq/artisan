# prowJob简介

prowJob是一个自定义命令行工具，用于JOB类的快速开发。

# 使用方式

1. 实现prowjob.Commander接口
   ```go
    type TestCommand struct {
    }
    
    func (t TestCommand) GetCommand() string {
        return "command:test"
    }
    
    func (t TestCommand) Usage() string {
        return "test command"
    }
    
    func (t TestCommand) Handle(ctx *prowjob.Context) {
        fmt.Println(ctx.params)
        fmt.Println("test command")
    }
   ```
   prowjob.go，并编译为./prowjob
   ```go
    func main() {
        job := prowjob.New()
        job.Add(commands.TestCommand{})
        job.Run()
    }
   ```
   编译后运行指定的job
   ```shell
   ./prowjob command:test arg1 -arg2 2 arg7 7 --arg3 3 -arg4=4 --arg5=5 arg6=6
   ```
2. 通过自定义方法来实现自定义命令

   prowjob.go，并编译为./prowjob
    ```go
    func main() {
        job := prowjob.New()
        job.AddFunc("command:testFunc", func(ctx *prowjob.Context) {
			fmt.Println(ctx.params)
            fmt.Println("test command")
        })
        job.Run()
    }
    ```
   编译后运行指定的job
   ```shell
   ./prowjob command:testFunc arg1 -arg2 2 arg7 7 --arg3 3 -arg4=4 --arg5=5 arg6=6
   ```
3. prowjob.Run方法的参数使用

   以上两种使用方法均是将main方法编译后，使用命令行调用。

   下面是通过代码调用的方式，可以用于debug或系统间调用。
    ```go
    func main() {
        job.Run("your command", "arg1", "arg2","...")
    }
    ```

4. 参数说明
   参数支持前缀`-`、`--`和无前缀，以等号或空格分隔键值，如果参数无值，则认为是字符串空值。

   `prowjob.Context`的`params`的值为map[string]string类型，key为参数名，value为参数值。

   | 参数       | 键    | 值   |
   |----------|------|-----|
   | arg1     | arg1 | ""  |
   | arg1 1   | arg1 | "1" |
   | arg1=1   | arg1 | "1" |
   | -arg2    | arg2 | ""  |
   | -arg2 2  | arg2 | "2" |
   | -arg2=2  | arg2 | "2" |
   | --arg3   | arg3 | ""  |
   | --arg3 3 | arg3 | "3" |
   | --arg3=3 | arg3 | "3" |

   ```shell
   ./prowjob command:test arg1 -arg2 2 arg7 7 --arg3 3 -arg4=4 --arg5=5 arg6=6
   ```
   ctx.params的值为：
   ```go
    map[arg1: arg2:2 arg3:3 arg4:4 arg5:5 arg6:6 arg7:7]
   ```
