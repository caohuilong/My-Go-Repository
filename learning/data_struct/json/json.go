package main

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
)

type BwInfos struct {
	Index  int32
	Ip     string
	BwName string
	BwType int32
	BizType string
}

func main() {
	data := "[{\"Ip\":\"150.223.255.175\",\"BwName\":\"bw-c50lfp35jed1ene7tuv0\",\"BwType\":1,\"BizType\":\"EECDN\"}]"
	var bwInfos []*BwInfos
	err := json.Unmarshal([]byte(data), &bwInfos)
	if err != nil {
		log.Errorf("unmarshal err: %s", err.Error())
	}

}
