package docker

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func ExecShell(writer http.ResponseWriter, request *http.Request) {
	var s = "uname"
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func BuildDockerImages(writer http.ResponseWriter, request *http.Request) {
	command := `./jar-to-docker.sh .`

	cmd := exec.Command("/bin/bash", "-c", command)

	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content)) //输出ls命令查看到的内容

	/*output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))*/

}
