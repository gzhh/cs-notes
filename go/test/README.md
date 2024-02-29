## Go Test

docs
- https://pkg.go.dev/testing
- https://go.dev/doc/tutorial/add-a-test
- https://gobyexample.com/testing


### go test 命令
- https://pkg.go.dev/cmd/go#hdr-Testing_flags

执行当前目录下所有的 *_test.go 文件，不显示执行过程
- `go test`

执行当前目录下指定的 *_test.go 文件
- `go test utils_test.go utils.go`
- https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file

执行当前目录下所有的 *_test.go 文件
- `go test -v`

执行当前目录下及其所有子目录下的 *_test.go 文件
- `go test -v ./...`

执行指定目录下所有的 *_test.go 文件
- `go test -v ./{specific_dir}/`

执行指定目录下及其所有子目录下的 *_test.go 文件
- `go test -v ./{specific_dir}/...`

### go test -run 指定测试函数
- https://go.dev/blog/subtests

精准匹配
- `go test -v -run={func_name}`

正则匹配
- `go test -v -run "^{some_words}$"`

### go test -cover 测试覆盖率
- https://go.dev/blog/cover
- https://go.dev/blog/integration-test-coverage

当前目录及其子目录的测试覆盖率
- `go test -cover ./...`
