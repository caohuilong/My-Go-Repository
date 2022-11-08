package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	str := make([]string, 2)
	str[0] = fmt.Sprintf(`echo -n '%s' > %s`, "test", "./network.yaml")
	str[1] = fmt.Sprintf(`echo -n '%s' > %s`, "test22", "./hash.yaml")

	script := strings.Join(str, "\n")
	base64Script := base64.StdEncoding.EncodeToString([]byte(script))
	ans := fmt.Sprintf("echo %s | base64 -d > %s; bash %s", base64Script, "./test.sh", "./test.sh")
	fmt.Println(ans)

	cmd := exec.Command("bash", "-c", `echo ZWNobyAtbiAndGVzdCcgPiAuL25ldHdvcmsueWFtbAplY2hvIC1uICd0ZXN0MjInID4gLi9oYXNoLnlhbWw= | base64 -d > ./test.sh`)
	o, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err.Error())
	}
	fmt.Printf("%s\n", o)
}
