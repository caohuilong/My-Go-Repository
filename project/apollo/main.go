package main

import (
	"fmt"
	"github.com/philchia/agollo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string `yaml:"name"`
}

func main() {
	apollo := agollo.NewClient(&agollo.Conf{
		AppID:              "dl-rendezvous",
		Cluster:            "RL",
		NameSpaceNames:     []string{"application.yml"},
		MetaAddr:           "http://apps.danlu.netease.com:37918",
		InsecureSkipVerify: true,
	})

	err := apollo.Start()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	logrus.Infof("Init apollo success")
	namespaceContent := apollo.GetContent(agollo.WithNamespace("application.yml"))

	apolloConfig := &Config{}
	if err := yaml.Unmarshal([]byte(namespaceContent), apolloConfig); err != nil {
		logrus.Fatal(err)
		return
	}
	fmt.Printf("%+v", apolloConfig)
}
