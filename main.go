package main

import "fmt"

func main() {
	patch := []byte(`{"metadata":{"labels":{"ray.io/woker-pod-name": "test"}`)
	fmt.Println(string(patch))

	str := fmt.Sprintf("{\"metadata\":{\"labels\":{\"ray.io/woker-pod-name\": \"%s\"}", "asdfa")
	fmt.Println(str)
}
