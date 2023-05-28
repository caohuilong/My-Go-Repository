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

type GamecoreConfig struct {
	RendezvousBaseUrl string `yaml:"rendezvous_base_url"`
	TrainId string `yaml:"train_id"`
	ResourceId string `yaml:"resource_id"`
	HostName string `yaml:"hostname"`
	StartupCmd string `yaml:"startup_cmd"`
	Gamecore_startup string `yaml:"gamecore_startup"`
}

func ReadFromFile1(filePath string) ( nodeConfigs []*GpuDriverConfig, err error) {
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

func ReadFromFile2(filePath string) (gamecoreConfig *GamecoreConfig, err error) {
	gamecoreConfig = &GamecoreConfig{}
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read gpu controller config file, %v", err.Error())
	}
	err = yaml.Unmarshal(f, gamecoreConfig)
	if err != nil {
		return nil, err
	}
	return gamecoreConfig, nil
}

func main() {
	nc, err := ReadFromFile1("learning/data_struct/yaml/config.yaml")
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Printf("%++v", *nc[0])

	c2, err := ReadFromFile2("learning/data_struct/yaml/config2.yaml")
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Printf("%+v", c2)
}
