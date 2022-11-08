// go 标准库 os
package main

import (
	"fmt"
	"os"
)

func file() {
	// 预定义变量, 保存命令行参数
	fmt.Println(os.Args)

	// 获取host name
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())

	// 获取全部环境变量
	env := os.Environ()
	for k, v := range env {
		fmt.Println(k, v)
	}

	// 终止程序
	// os.Exit(1)

	// 获取一条环境变量
	fmt.Println(os.Getenv("PATH"))

	// 获取当前目录
	dir, err := os.Getwd()
	fmt.Println(dir, err)

	// 创建目录
	err = os.Mkdir(dir+"/new_file", 0755)
	fmt.Println(err)

	// 创建目录
	err = os.MkdirAll(dir+"/new", 0755)
	fmt.Println(err)

	// 删除目录
	err = os.Remove(dir + "/new_file")
	err = os.Remove(dir + "/new")
	fmt.Println(err)

	// 创建临时目录
	tmp_dir := os.TempDir()
	fmt.Println(tmp_dir)
}

func stdoutFunc() {
	stdout := os.Stdout.Name()
	fmt.Println(stdout)


}

func main() {
	stdoutFunc()

	return
}
