package main

import (
	"fmt"
	"os/exec"
)

func main() {
	err := echo("0000:1a:00.0", "vfio-pci")
	if err != nil{
		panic(err)
	}

	err = bash("0000:1a:00.1", "vfio-pci")
	if err != nil {
		panic(err)
	}

	return
}

func echo(BDF string, driver string) error {
	cmd := exec.Command("echo", driver, ">", "./test.txt")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func bash(BDF string, driver string) error {
	command := fmt.Sprintf("echo %s > %s", driver, "test2.txt")
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}