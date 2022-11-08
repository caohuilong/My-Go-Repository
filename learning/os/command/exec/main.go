package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func SendSignalToContainers(sig syscall.Signal) error {
	containerId := "d05fe537b4aa81a2f3f87d69a9fc77966b86011cfb4682a6bf1dd0b31e690680"

	//args := fmt.Sprintf("-n k8s.io task kill -s %v %s", sig, containerId)
	args2 := []string{"-n", "k8s.io", "task", "kill", "-s", "SIGCONT", containerId}
	//cmd := exec.Command("ctr", "-n k8s.io task kill -s 3 d05fe537b4aa81a2f3f87d69a9fc77966b86011cfb4682a6bf1dd0b31e690680")
	//cmd := exec.Command("ctr", []string{"-n", "k8s.io", "container", "ls"}...)
	cmd := exec.Command("ctr", args2...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("send signal %s to container %s failed, err: %v", sig.String(), containerId, err.Error())
		return err
	}
	fmt.Printf("%v", out)

	return nil
}

func main() {
	//err := SendSignalToContainers(syscall.SIGQUIT)
	//if err != nil {
	//	panic(err)
	//}

	_ = SendSignalToContainers(syscall.SIGCONT)
}
