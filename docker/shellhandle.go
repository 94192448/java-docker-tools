package docker

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
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
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

var runMsg = ""
var runEnd = false

func BuildDockerImages(appName string, appVersion string) {
	runEnd = false

	command := "./upload/jar-to-docker.sh " + strings.ToLower(appName) + " " + strings.ToLower(appVersion)

	cmd := exec.Command("/bin/bash", "-c", command)
	//cmd := exec.Command("test.bat")

	log.Println("Starting exec-> ", command, cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal("start error:", err)
	}

	reader := bufio.NewReader(stdout)

	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			log.Println("reader error:", err2.Error())
			break
		}
		runMsg = line
	}

	cmd.Wait()

	runEnd = true
}
