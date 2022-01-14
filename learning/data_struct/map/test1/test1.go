package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	//定义map
	m := make(map[string]*student)
	n := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for i := 0; i < len(stus); i++ {
		n[stus[i].Name] = &stus[i]
	}

	//打印map
	fmt.Println("m:")
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
	fmt.Println("n:")
	for k, v := range n {
		fmt.Println(k, "=>", v.Name)
	}
}
