package service

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
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

func BuildDockerImages(appName string, appVersion string) {

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
			log.Println("end of the reader error:", err2.Error())
			break
		}
		if line != "" {
			runMsg = time.Now().Format("2006-01-02 15:04:05") + " " + line
			log.Println(runMsg)
			time.Sleep(1 * time.Second)
		}
	}

	time.Sleep(500 * time.Millisecond)
	runMsg = time.Now().Format("2006-01-02 15:04:05") + " " + "Success created docker images...\n"

	cmd.Wait()
}
