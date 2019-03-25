package docker

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func Exec_shell(writer http.ResponseWriter, request *http.Request) {
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
