package main

import (
	"fmt"
	"strings"
)

func main() {
	anno := make(map[string][]string)
	list := anno["1"]
	str := strings.Join(list, ",")
	fmt.Println(str)

	f(anno)
	for k, _ := range anno {
		fmt.Println(k)
	}
}

func f(anno map[string][]string) {
	anno["aaa"] = []string{"1"}
}
