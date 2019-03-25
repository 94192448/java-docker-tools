package docker

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

func BuildDockerImages(appName string, appVersion string) {
	command := `./upload/jar-to-docker.sh `
	command += appName
	command += " " + appVersion

	fmt.Print(command)

	cmd := exec.Command("/bin/bash", "-c", command)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()

}
