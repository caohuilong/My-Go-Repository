package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	binary, err := exec.LookPath("top")
	if err != nil {
		panic(err)
	}

	args := []string{"300"}
	sigs := newOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGQUIT, syscall.SIGCONT)

	var proc *os.Process
	proc, err = os.StartProcess(binary, args, &os.ProcAttr{})
	if err != nil {
		panic(err)
	}
	log.Printf("starting gpu device plugin, PID: %d", proc.Pid)

L:
	for {
		select {
		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				log.Println("Wrap Received SIGHUP, restarting.")
				proc.Signal(syscall.SIGHUP)
			case syscall.SIGQUIT:
				log.Println("Wrap Received SIGQUIT, stopping.")
				proc.Signal(syscall.SIGTERM)
				proc.Wait()
			case syscall.SIGCONT:
				log.Println("Wrap Received SIGCONT, continue.")
				proc, err = os.StartProcess(binary, args, &os.ProcAttr{})
				if err != nil {
					panic(err)
				}
				log.Printf("continue gpu device plugin, PID: %d", proc.Pid)
			default:
				log.Printf("Wrap Received signal \"%v\", shutting down.", s)
				proc.Signal(syscall.SIGTERM)
				break L
			}
		}
	}

	log.Printf("stop plugin wrap")
}

func newOSWatcher(sigs ...os.Signal) chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sigs...)

	return sigChan
}
