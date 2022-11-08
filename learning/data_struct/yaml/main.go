package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type GpuDriverConfig struct {
	NodeName string `yaml:"nodeName"`
	Nvidia   int    `yaml:"nvidia"`
	Vfio     int    `yaml:"vfio"`
}

func ReadFromFile(filePath string) ( nodeConfigs []*GpuDriverConfig, err error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read gpu controller config file, %v", err.Error())
	}
	err = yaml.Unmarshal(f, &nodeConfigs)
	if err != nil {
		return nil, err
	}

	return nodeConfigs, nil
}

func main() {
	nc, err := ReadFromFile("learning/data_struct/yaml/config.yaml")
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Printf("%++v", *nc[0])
}
