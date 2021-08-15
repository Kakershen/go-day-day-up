# go-day-day-up

## go命令



### go build

在main包下执行，用于将go文件编译成可执行文件。

1. go build：默认在当前目录下生成可执行文件。
2. go build -o 路径/文件名.exe：将编译生成的文件输出到指定目录，不加路径时默认当前路径。
3.  go build a.go：只编译当前目录中的a.go文件。



> go build会忽略目录下以“_”或“.”开头的go文件。



### go install

在main包下执行，编译生成可执行文件，并将可执行文件放在$GOPATH/bin目录中。

在非main下执行，用于生成库，将编译后的文件放在$GOPATH/pkg目录中。



### go clean

清除当前包中编译生成的文件。

1. go clean：执行清除命令，删除当前文件夹中编译产生的文件。
2. go clean -n：打印将要执行的清除命令但是不执行。
3. go clean -x：打印将要执行的清除命令且执行清除操作。



### go get

下载包及其依赖，默认还会执行 go install 。

- 默认将下载的包放在$GOPATH/pkg/mod目录下。
- 下载完成之后，自动完成编译安装。



### go run

直接编译执行main包中的main函数，且不会在当前目录中产生可执行文件。



如果执行 go run 命令出现以下错误：

```
go run: no go files listed
```

可以在命令后面加 . 执行

```shell
go run .
```











