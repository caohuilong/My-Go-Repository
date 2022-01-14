package main

import (
	"fmt"
)

type Dog struct {
	age  int
	name string
}

func main() {
	roger := Dog{5, "Roger"}
	mydog := roger
	fmt.Printf("roger addr %p\n", &roger)
	fmt.Printf("mydog addr %p\n", &mydog)
	fmt.Println("Roger and mydog are equal structs?", roger == mydog)
	mydog.name = "piggie"
	fmt.Println("Roger and mydog are equal structs?", roger == mydog)
	fmt.Println(roger.name)

	networks := []string{}
	add(networks)
	for _, item := range networks {
		fmt.Println(item)
	}
}

func add(networks []string) {
	networks = append(networks, "aaa")
}
