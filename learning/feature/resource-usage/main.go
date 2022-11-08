package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func getStorageInfo() {
	var storageinfo []storageInfo
	var loaclStorages []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		return
	}

	for _, storage := range storageinfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size / 1024 / 1024 / 1024,
			Free:       storage.FreeSpace / 1024 / 1024 / 1024,
		}
		if info.Total >= 1 {
			fmt.Printf("%s总大小%dG，可用%dG\n", info.Name, info.Total, info.Free)
			loaclStorages = append(loaclStorages, info)
		}
	}
	//fmt.Printf("localStorages:= %v\n", loaclStorages)
}

type ComputerMonitor struct {
	CPU float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}

// GetCPUPercent 获取CPU使用率
func GetCPUPercent() float64 {
	percent, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		log.Fatalln(err.Error())
		return -1
	}
	return percent[0]
}

// GetMemPercent 获取内存使用率
func GetMemPercent() float64 {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln(err.Error())
		return -1
	}
	fmt.Println("内存总量：", memInfo.Total)
	fmt.Println("内存已用量：", memInfo.Used)
	fmt.Println("内存可用量：", memInfo.Available)

	return memInfo.UsedPercent
}

func GetCpuMem() ComputerMonitor {
	var res ComputerMonitor
	res.CPU = GetCPUPercent()
	res.Mem = GetMemPercent()
	//fmt.Printf("%v", res)
	fmt.Printf("cpu使用率：%.2f%%\n", res.CPU)
	fmt.Printf("内存使用率：%.2f%%\n", res.Mem)
	return res
}

func GetGpuPercent() float64 {
	listCmd := exec.Command("cmd", "/C", "nvidia-smi -L")
	o1, err := listCmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return 0
	}
	str := string(o1)
	fmt.Print(str)

	strs := strings.Split(strings.TrimSuffix(string(o1), "\r\n"),"\r\n")
	gpuNum := len(strs)
	fmt.Println(len(strs))

	//total := make([]float64, 0)
	//usage := make([]float64, 0)
	for i := 0; i < gpuNum; i++ {
		nvidiaCmd := exec.Command("cmd", "/C", "nvidia-smi --query-gpu=utilization.gpu --format=csv,noheader,nounits")
		output2, err := nvidiaCmd.Output()
		if err != nil {
			fmt.Printf(err.Error())
			return 0
		}
		str := string(output2)
		strs := strings.Split(strings.TrimSuffix(str, "\r\n"), "\r\n")

		for i := 0; i < len(strs); i++ {
			if strings.Contains(strs[i], "Gpu") {
				fmt.Println(strs[i])
			}
			//if strings.Contains(strs[i], "")
		}
	}
	//cmd := exec.Command("cmd", "/C", "nvidia-smi -i 0 -q -d UTILIZATION,MEMORY")
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Print(string(output))
	return 0
}

func main() {
	//getStorageInfo()
	GetCpuMem()
	//GetGpuPercent()
}
