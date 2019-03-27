package docker

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/cihub/seelog"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func ExecShell(writer http.ResponseWriter, request *http.Request) {
	var s = "uname"
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("%s", out.String())
}

func BuildDockerImages(appName string, appVersion string) {
	command := "./upload/jar-to-docker.sh "
	command += strings.ToLower(appName)
	command += " " + strings.ToLower(appVersion)

	fmt.Println("Starting exec-> " + command)

	//cmd := exec.Command("/bin/bash", "-c", command)
	cmd := exec.Command(command)

	stdout, err := cmd.StdoutPipe()
	fmt.Println("exec end")

	if err != nil {
		fmt.Println(err)
		log.Error(err)
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			fmt.Println(err2)
			break
		}
		fmt.Println(line)
		log.Info(line)
	}
	fmt.Println("exec println end")

	cmd.Wait()

}
