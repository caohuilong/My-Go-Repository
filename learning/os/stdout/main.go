package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("test.......")
	//logFile, _ := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	cmd := exec.Command("./test.sh")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	defer cmd.Wait()
	//err = cmd.Wait()
	//if err != nil {
	//	panic(err)
	//}
	time.Sleep(10)
	fmt.Printf("end.......")
	return
}
