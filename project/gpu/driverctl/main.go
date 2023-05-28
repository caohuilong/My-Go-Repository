package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Args)

	cmd := exec.Command("driverctl", "--nosave", "set-override", os.Args[1], os.Args[2])
	err := cmd.Run()
	if err != nil {
		fmt.Errorf("err: %s", err)
		return
	}
	return
}
