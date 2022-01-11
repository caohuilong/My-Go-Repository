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
}

func f(anno map[string]string) {
	anno["aaa"] = "1"
}
