package main

// 命令行相关操作

import (
	"flag"
	"fmt"
)

// run: go run . -surname="倪" -personName="豫才" -id="-1"
// 使用flag从命令行中读取参数
func readParamFromCmd() {
	// 定义一个类型为string，名称为surname的命令行参数
	surname := flag.String("surname", "王", "您的姓")
	// 定义一个类型为string，名称为personName的命令行参数
	// 直接传入变量的地址获取参数
	var personName string
	flag.StringVar(&personName, "personName", "小二", "名字")
	// 定义一个类型为int，名称为age的命令行参数
	age := flag.Int("age", 0, "年龄")
	// 解析命令行参数
	flag.Parse()

	fmt.Printf("I am %v %v, and age is %v\n", *surname, personName, *age)
}
