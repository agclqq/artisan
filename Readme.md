[中文说明](Readme_zh.md)
# what is prowJob?

prowJob is a custom command line tool for rapid development of JOB.

Common way： `prowjob $command [$arg1 ...]`

# Getting started
1. getting prowjob
   ```shell
   go get github.com/agclqq/prowjob
   ```
   1. Implement prowjob.Commander interface
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
      Compile prowjob.go to prowjob
      ```go
      package main
      import (
          "github.com/agclqq/prowjob"
      )
      
      func main() {
           job := prowjob.New()
           job.Add(TestCommand{})
           job.Run()
       }
      ```
      run job
      ```shell
      ./prowjob command:test arg1 -arg2 2 arg7 7 --arg3 3 -arg4=4 --arg5=5 arg6=6
      ```
   2. Implement custom commands

      Compile prowjob.go to prowjob
       ```go
       package main
       import (
           "fmt"
           "github.com/agclqq/prowjob"
       )
   
       func main() {
           job := prowjob.New()
           job.AddFunc("command:testFunc", func(ctx *prowjob.Context) {
               fmt.Println(ctx.params)
               fmt.Println("test command")
           })
           job.Run()
       }
       ```
      run job
      ```shell
      ./prowjob command:testFunc arg1 -arg2 2 arg7 7 --arg3 3 -arg4=4 --arg5=5 arg6=6
      ```
3. prowjob.Run method parameter usage

   The above two usage methods are compiled by the main method, using the command line to invoke.

   Here is how to call by code, which can be used for debug or inter-system calls
    ```go
    func main() {
        job.Run("your command", "arg1", "arg2","...")
    }
    ```

4. parameter description
   Command parameters support prefixes -, --, and no prefixes.
   
   The key values are separated by equal signs or Spaces. If the parameter has no value, it is considered to be an empty string value.
   
   The params value of prowjob.Context is map[string]string,key is the parameter name, and value is the parameter value

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
